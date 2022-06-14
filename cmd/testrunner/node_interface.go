package main

import (
	"github.com/bottlepay/lightning-benchmark/graphreader"
	"github.com/lightningnetwork/lnd/lnrpc"
)

type nodeInterface interface {
	Connect(key, host string) error
	NewAddress() (string, error)
	OpenChannel(peerKey string, amtSat int64, pushAmtSat int64, private bool) error
	ActiveChannels() (int, error)
	AddInvoice(amtMsat int64) (string, error)
	SendPayment(invoice string, aliasMap map[string]string) error
	Close()
	HasFunds() error
	PubKey() string
	SetPolicy(chanPoint *lnrpc.ChannelPoint, policy *graphreader.PolicyData) error
	IsSynced(int, int) (bool, error)
	Restart() (nodeInterface, error)
	GetChannelBalance() (int, error)
}

type info struct {
	key    string
	synced bool
}
