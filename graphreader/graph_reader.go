package graphreader

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type PolicyData struct {
	CltvDelta  int `yaml:"cltvDelta"`
	BaseFee    int `yaml:"baseFee"`
	FeeRate    int `yaml:"feeRate"`
	HtlcMaxSat int `yaml:"htlcMaxSat"`
}

type YamlChannel struct {
	Capacity      int
	Policy        string
	RemoteBalance int    `yaml:"remoteBalance"`
	RemotePolicy  string `yaml:"remotePolicy"`
}

type YamlGraph struct {
	Policies map[string]PolicyData
	Channels map[string]YamlNode
	Tests    []map[string]int
}

type YamlNode map[string][]YamlChannel

func Read(file string) (*YamlGraph, error) {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var g YamlGraph
	err = yaml.UnmarshalStrict(yamlFile, &g)
	if err != nil {
		return nil, err
	}

	policyNames := make(map[string]struct{})
	for name := range g.Policies {
		policyNames[name] = struct{}{}
	}

	nodeNames := make(map[string]struct{})
	for alias := range g.Channels {
		nodeNames[alias] = struct{}{}
	}

	for _, node := range g.Channels {
		for peerName, peer := range node {
			if _, ok := nodeNames[peerName]; !ok {
				return nil, fmt.Errorf("undefined node %v", peerName)
			}
			for _, channel := range peer {
				if _, ok := policyNames[channel.Policy]; !ok {
					return nil, fmt.Errorf("undefined policy %v", channel.Policy)
				}
				if _, ok := policyNames[channel.RemotePolicy]; !ok {
					return nil, fmt.Errorf("undefined policy %v", channel.RemotePolicy)
				}
			}
		}
	}

	for _, test := range g.Tests {
		if len(test) != 1 {
			return nil, fmt.Errorf("invalid test definition")
		}

		for dest, amt := range test {
			if amt == 0 {
				return nil, fmt.Errorf("test amount zero")
			}

			if _, ok := nodeNames[dest]; !ok {
				return nil, fmt.Errorf("undefined node %v in test", dest)
			}
		}
	}

	return &g, nil
}
