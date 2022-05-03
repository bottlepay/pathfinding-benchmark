package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/bottlepay/lightning-benchmark/graphreader"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:   "run",
	Action: run,
}

func getBitcoindConnection() (*rpcclient.Client, error) {
	connConfig := rpcclient.ConnConfig{
		Host:                 "bitcoind:8332",
		User:                 "test",
		Pass:                 "test",
		DisableConnectOnNew:  true,
		DisableAutoReconnect: false,
		DisableTLS:           true,
		HTTPPostMode:         true,
	}

	bitcoindConn, err := rpcclient.New(&connConfig, nil)
	if err != nil {
		log.Errorw("New rpc connection", "err", err)
		return nil, err
	}

	log.Infow("Attempting to connect to bitcoind")

	for {
		_, err := bitcoindConn.GetBlockChainInfo()
		if err == nil {
			log.Infow("Connected to bitcoind")
			return bitcoindConn, nil
		}

		time.Sleep(time.Second)
	}
}

func run(_ *cli.Context) error {
	target := os.Getenv("TARGET")
	if target != "cln" && target != "lnd" {
		return fmt.Errorf("unknown target %v", target)
	}

	graph, err := graphreader.Read("/graph.yml")
	if err != nil {
		return err
	}

	bitcoindConn, err := getBitcoindConnection()
	if err != nil {
		return err
	}

	log.Infow("Creating bitcoind wallet")
	_, err = bitcoindConn.CreateWallet("")
	if err != nil {
		return err
	}

	addr, err := bitcoindConn.GetNewAddress("")
	if err != nil {
		return err
	}

	log.Infow("Bitcoin address", "address", addr.String())

	log.Infow("Activate segwit")
	_, err = bitcoindConn.GenerateToAddress(400, addr, nil)
	if err != nil {
		return err
	}

	log.Infow("Fund senders")

	clients := make(map[string]nodeInterface)
	aliasMap := make(map[string]string)
	for node := range graph.Channels {
		var client nodeInterface
		var err error

		if node == "start" && target == "cln" {
			client, err = getClightningConnection(node)
		} else {
			client, err = getLndConnection(node)
		}

		if err != nil {
			return err
		}
		defer client.Close()

		clients[node] = client
		aliasMap[client.PubKey()] = node

		addrResp, err := client.NewAddress()
		if err != nil {
			return err
		}
		log.Infow("Generated funding address",
			"node", node, "address", addrResp)

		senderAddr, err := btcutil.DecodeAddress(addrResp, &chaincfg.RegressionNetParams)
		if err != nil {
			return err
		}

		for i := 0; i < 2; i++ {
			_, err = bitcoindConn.SendToAddress(senderAddr, 1e8)
			if err != nil {
				return err
			}
		}
	}

	_, err = bitcoindConn.GenerateToAddress(1, addr, nil)
	if err != nil {
		return err
	}

	log.Infow("Wait for coin and open channels")
	expectedChannelCount := make(map[string]int)

	type policyTask struct {
		node      string
		chanPoint *lnrpc.ChannelPoint
		policy    *graphreader.PolicyData
	}

	var policyTasks []*policyTask

	for alias, peers := range graph.Channels {
		nodeLog := log.With("node", alias)
		client := clients[alias]

		if err := client.HasFunds(); err != nil {
			return err
		}

		for peer, channels := range peers {
			peerClient := clients[peer]
			peerKey := peerClient.PubKey()
			rpcHost := fmt.Sprintf("lnd_%v:9735", peer)

			nodeLog.Infow("Connecting", "peer", peer,
				"peerPubKey", peerKey, "host", rpcHost)

			err := client.Connect(peerKey, rpcHost)
			if err != nil {
				return err
			}

			nodeLog.Infow("Open channel", "peer", peer)

			for _, channel := range channels {
				chanPoint, err := client.OpenChannel(
					peerKey, int64(channel.Capacity),
					int64(channel.RemoteBalance),
				)
				if err != nil {
					return err
				}

				expectedChannelCount[alias]++
				expectedChannelCount[peer]++

				policy, ok := graph.Policies[channel.Policy]
				if !ok {
					return fmt.Errorf("policy %v not found", channel.Policy)
				}

				remotePolicy, ok := graph.Policies[channel.RemotePolicy]
				if !ok {
					return fmt.Errorf("policy %v not found", channel.RemotePolicy)
				}

				policyTasks = append(policyTasks,
					&policyTask{
						node:      alias,
						chanPoint: chanPoint,
						policy:    &policy,
					},
					&policyTask{
						node:      peer,
						chanPoint: chanPoint,
						policy:    &remotePolicy,
					},
				)
			}
		}
	}

	// Mine 6 blocks.
	log.Infow("Confirm channels")
	_, err = bitcoindConn.GenerateToAddress(6, addr, nil)
	if err != nil {
		return err
	}

	// Wait for channels to become active.
	for alias, expectedCount := range expectedChannelCount {
		err := try(func() error {
			count, err := clients[alias].ActiveChannels()
			if err != nil {
				return err
			}

			log.Debugw("Waiting for active channels",
				"node", alias, "expected", expectedCount,
				"count", count)

			if count != expectedCount {
				return errors.New("still waiting for channels")
			}

			return nil
		})
		if err != nil {
			return err
		}
	}

	// Set policies.
	for _, task := range policyTasks {
		if task.chanPoint == nil {
			log.Debugw("Skipping policy setting", "node", task.node)

			continue
		}
		log.Debugw("Setting policy", "node", task.node, "channel", task.chanPoint, "policy", task.policy)
		err := try(func() error {
			return clients[task.node].SetPolicy(task.chanPoint, task.policy)
		})
		if err != nil {
			return err
		}
	}

	// Generate test invoices.
	var invoices []string

	invoice, err := clients["destination1"].AddInvoice(50000 * 1e3)
	if err != nil {
		return err
	}
	invoices = append(invoices, invoice)

	invoice, err = clients["destination2"].AddInvoice(50000 * 1e3)
	if err != nil {
		return err
	}
	invoices = append(invoices, invoice)

	// Wait for propagation
	log.Infow("Wait for propagation")
	time.Sleep(10 * time.Second)

	start := time.Now()
	for _, invoice := range invoices {
		startPayment := time.Now()
		log.Infow("Sending payment", "invoice", invoice)
		err = clients["start"].SendPayment(invoice, aliasMap)
		if err != nil {
			return err
		}
		elapsed := time.Since(startPayment)
		log.Infow("Time elapsed", "time", elapsed.String())
	}

	elapsed := time.Since(start)
	log.Infow("Total time elapsed", "time", elapsed.String())

	log.Infow("Done")

	return nil
}

func try(f func() error) error {
	attempts := 0
	for {
		err := f()
		if err == nil {
			return nil
		}

		attempts++
		if attempts == 300 {
			return fmt.Errorf("too many errors: %w", err)
		}

		time.Sleep(time.Second)
	}
}
