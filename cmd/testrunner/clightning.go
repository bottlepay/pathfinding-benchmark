package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bottlepay/lightning-benchmark/graphreader"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/niftynei/glightning/glightning"
)

type clightningConnection struct {
	client *glightning.Lightning

	pubKey string
	alias  string
}

func getClightningConnection(alias string) (*clightningConnection, error) {
	client, pubkey, err := getClightningConnect(alias)
	if err != nil {
		return nil, err
	}

	return &clightningConnection{
		client: client,
		pubKey: pubkey,
		alias:  alias,
	}, nil

}

func getClightningConnect(alias string) (*glightning.Lightning, string, error) {
	logger := log.With("node", alias)

	rpcHost := fmt.Sprintf("node-%v:9835", alias)

	client := glightning.NewLightning()
	client.StartUp(rpcHost)

	logger.Infow("Attempting to connect to c-lightning (please be patient)")
	for {
		info, err := client.GetInfo()
		if err == nil {
			if !info.IsBitcoindSync() || !info.IsLightningdSync() {
				time.Sleep(time.Second)

				continue
			}

			logger.Infow("Connected to c-lightning", "key", info.Id)

			return client, info.Id, nil
		}

		time.Sleep(time.Second)
	}
}

func (l *clightningConnection) PubKey() string {
	return l.pubKey
}

func (l *clightningConnection) Close() {
}

func (l *clightningConnection) GetInfo() (*info, error) {
	infoResp, err := l.client.GetInfo()
	if err != nil {
		return nil, err
	}

	return &info{
		key:    infoResp.Id,
		synced: infoResp.IsBitcoindSync(),
	}, nil
}

func (l *clightningConnection) Connect(key, address string) error {
	parts := strings.Split(address, ":")
	host := parts[0]
	port, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return err
	}

	_, err = l.client.ConnectPeer(key, host, uint(port))
	return err
}

func (l *clightningConnection) NewAddress() (string, error) {
	return l.client.NewAddr()
}

func (l *clightningConnection) OpenChannel(peerKey string, amtSat int64, pushAmtSat int64) (*lnrpc.ChannelPoint, error) {
	sat := glightning.NewSat64(uint64(amtSat))
	pushMsat := glightning.NewMsat(uint64(pushAmtSat) * 1e3)
	resp, err := l.client.FundChannelExt(peerKey, sat, nil, false, nil, pushMsat)
	if err != nil {
		return nil, err
	}

	txID, err := chainhash.NewHashFromStr(resp.FundingTxId)
	if err != nil {
		return nil, err
	}

	return &lnrpc.ChannelPoint{
		FundingTxid: &lnrpc.ChannelPoint_FundingTxidBytes{FundingTxidBytes: txID[:]},
		// Index unavailable for cln?
	}, err
}

func (l *clightningConnection) ActiveChannels() (int, error) {
	infoResp, err := l.client.GetInfo()
	if err != nil {
		return 0, err
	}

	return infoResp.ActiveChannelCount, nil

}

func (l *clightningConnection) IsSynced(totalEdges, localChannels int) (bool, error) {
	channels, err := l.client.ListChannelsBySource("")
	if err != nil {
		return false, err
	}

	var activeCount int
	for _, ch := range channels {
		if ch.IsActive {
			activeCount++
		}
	}

	log.Debugw("Syncing", "edges", activeCount, "totalEdges", totalEdges)

	return activeCount == totalEdges, nil
}

func (l *clightningConnection) AddInvoice(amtMsat int64) (string, error) {
	label := randomString(20)
	invoice, err := l.client.Invoice(uint64(amtMsat), label, "test")
	if err != nil {
		return "", err
	}
	return invoice.Bolt11, nil
}

func (l *clightningConnection) SendPayment(invoice string, aliasMap map[string]string) error {
	status, err := l.client.PayBolt(invoice)
	if err != nil {
		return err
	}

	if status.Status != "complete" {
		return errors.New("payment failed")
	}

	return nil
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func (l *clightningConnection) HasFunds() error {
	for {
		resp, err := l.client.ListFunds()
		if err != nil {
			return err
		}
		if len(resp.Outputs) > 0 {
			return nil
		}

		time.Sleep(time.Second)
	}
}

func (l *clightningConnection) SetPolicy(chanPoint *lnrpc.ChannelPoint,
	policy *graphreader.PolicyData) error {

	return nil

	// fundingTxID, err := chainhash.NewHash(chanPoint.GetFundingTxidBytes())
	// if err != nil {
	// 	return err
	// }

	// _, err = l.client.SetChannelFee(
	// 	fundingTxID.String(),
	// 	strconv.FormatInt(int64(policy.BaseFee), 10),
	// 	uint32(policy.FeeRate),
	// )

	// return err
}

func (l *clightningConnection) Restart() (nodeInterface, error) {
	_, err := l.client.Stop()
	if err != nil {
		return nil, err
	}

	time.Sleep(3 * time.Second)

	conn, err := getClightningConnection(l.alias)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
