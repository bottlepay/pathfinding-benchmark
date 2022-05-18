package main

import (
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/bottlepay/lightning-benchmark/graphreader"
	"gopkg.in/yaml.v2"
)

type config struct {
	Version  string
	Volumes  map[string]volume
	Services map[string]service
}

type volume struct {
}

type service struct {
	Restart        string `yaml:",omitempty"`
	Image          string `yaml:",omitempty"`
	Ports, Volumes []string
	Command        string
	DependsOn      []string "depends_on"
	Build          string   `yaml:",omitempty"`
	Environment    []string
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

// This function generates a docker-compose.yml file from a graph definition.
func run() error {
	graph, err := graphreader.Read("graph.yml")
	if err != nil {
		return err
	}

	cfg := config{
		Services: map[string]service{
			"bitcoind": {
				Image:   "kylemanna/bitcoind",
				Volumes: []string{"./bitcoin.conf:/bitcoin/.bitcoin/bitcoin.conf"},
			},
		},
		Volumes: map[string]volume{
			"lnd": {},
		},
		Version: "3.4",
	}

	nodesSet := make(map[string]graphreader.PolicyData)
	for alias, peers := range graph.Nodes {
		var ok bool
		nodesSet[alias], ok = graph.Policies[peers.Policy]
		if !ok {
			return errors.New("unknown policy")
		}
	}

	var nodes []string
	for n := range nodesSet {
		nodes = append(nodes, n)
	}
	sort.SliceStable(nodes, func(i, j int) bool { return nodes[i] < nodes[j] })

	target := os.Getenv("TARGET")
	if target != "cln" && target != "lnd" && target != "lnd-managej" {
		return fmt.Errorf("unknown target %v", target)
	}

	const concurrency = 5

	prevLnd := make([]string, concurrency)

	cfg.Services["_lnd_build"] = service{
		Image:   "pathfinding-benchmark-lnd",
		Command: "echo build completed",
		Build:   "lnd",
	}

	for i, alias := range nodes {
		startupChainIdx := i % concurrency

		name := "node-" + alias

		policy := nodesSet[alias]

		var serv service

		switch {
		case alias == "start" && target == "cln":
			serv = service{
				// TODO: Replace with 0.11 image.
				Image: "elementsproject/lightningd:v0.11.0.1",
				// Image:     "cln:latest",

				DependsOn: []string{"bitcoind"},
				Volumes: []string{
					"./c-lightning.conf:/root/.lightning/regtest/config",
				},
				Environment: []string{
					"EXPOSE_TCP=true",
					"LIGHTNINGD_NETWORK=regtest",
				},
				Command: "--network=regtest",
			}

		default:
			serv = service{
				Image:     "pathfinding-benchmark-lnd",
				DependsOn: []string{"bitcoind", "_lnd_build"},
				Volumes: []string{
					"./lnd.conf:/root/.lnd/lnd.conf",
					"lnd:/cfg",
				},
				Command: fmt.Sprintf("--tlsextradomain=%v --alias=%v --tlscertpath=/cfg/%v/tls.cert --adminmacaroonpath=/cfg/%v/admin.macaroon --bitcoin.basefee=%v --bitcoin.feerate=%v",
					name, alias, alias, alias, policy.BaseFee, policy.FeeRate),
			}

			// All nodes except the start node are started in a chain to keep
			// peak mem usage lower.
			if alias != "start" {
				if prevLnd[startupChainIdx] != "" {
					serv.Environment = []string{fmt.Sprintf("WAIT_FOR_LND=%v", prevLnd[startupChainIdx])}
				}

				prevLnd[startupChainIdx] = alias
			}
		}

		serv.Restart = "unless-stopped"

		cfg.Services[name] = serv

	}

	if target == "lnd-managej" {
		cfg.Services["lnd-managej"] = service{
			Build:     "lnd-managej",
			DependsOn: []string{"node-start"},
			Volumes: []string{
				"lnd:/cfg",
			},
			Ports:   []string{"9081:8081"},
			Restart: "unless-stopped",
		}
	}

	cfg.Services["testrunner"] = service{
		Build: ".",
		Volumes: []string{
			"lnd:/lnd",
			"./graph.yml:/graph.yml",
		},
		Environment: []string{
			fmt.Sprintf("TARGET=%v", target),
		},
	}

	yamlBytes, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	fmt.Println(string(yamlBytes))

	return nil
}
