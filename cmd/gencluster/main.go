package main

import (
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
	Restart        string
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

	var depends = []string{
		"bitcoind",
	}

	nodesSet := make(map[string]struct{})
	for node, peers := range graph.Channels {
		nodesSet[node] = struct{}{}
		for peer := range peers {
			nodesSet[peer] = struct{}{}
		}
	}

	var nodes []string
	for n := range nodesSet {
		nodes = append(nodes, n)
	}
	sort.SliceStable(nodes, func(i, j int) bool { return nodes[i] < nodes[j] })

	target := os.Getenv("TARGET")
	if target != "cln" && target != "lnd" {
		return fmt.Errorf("unknown target %v", target)
	}

	for _, alias := range nodes {
		name := "lnd_" + alias
		depends = append(depends, name)

		var serv service

		if alias == "start" && target == "cln" {
			serv = service{
				// TODO: Replace with 0.11 image.
				Image: "elementsproject/lightningd:v0.9.3",
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
		} else {
			serv = service{
				Image:     "lightninglabs/lnd:v0.14.3-beta",
				DependsOn: []string{"bitcoind"},
				Volumes: []string{
					"./lnd.conf:/root/.lnd/lnd.conf",
					"lnd:/cfg",
				},
				Command: fmt.Sprintf("--tlsextradomain=%v --alias=%v --tlscertpath=/cfg/%v/tls.cert --adminmacaroonpath=/cfg/%v/admin.macaroon",
					name, alias, alias, alias),
			}
		}

		serv.Restart = "unless-stopped"

		cfg.Services[name] = serv
	}

	cfg.Services["testrunner"] = service{
		Build:     ".",
		DependsOn: depends,
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
