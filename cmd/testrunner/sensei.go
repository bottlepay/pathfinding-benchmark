package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bottlepay/lightning-benchmark/graphreader"
	"github.com/bottlepay/lightning-benchmark/sensei"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/routing/route"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type senseiConnection struct {
	conn   *grpc.ClientConn
	pubKey route.Vertex
	alias  string
	client sensei.NodeClient
}

func getSenseiConnection(alias string) (*senseiConnection, error) {
	logger := log.With("node", alias)

	rpcHost := fmt.Sprintf("node-%v:5401", alias)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	logger.Infow("Attempting to connect to sensei")

	var adminConn *grpc.ClientConn
	err := tryFunc(
		func() error {
			var err error
			adminConn, err = grpc.Dial(rpcHost, opts...)
			if err != nil {
				return err
			}
			return err
		}, 60)
	if err != nil {
		return nil, err
	}
	defer adminConn.Close()

	logger.Infow("Creating node")

	adminClient := sensei.NewAdminClient(adminConn)

	createResp, err := adminClient.CreateAdmin(context.Background(), &sensei.CreateAdminRequest{
		Username:   "admin",
		Alias:      alias,
		Passphrase: "secret",
		Start:      true,
	})
	if err != nil {
		return nil, err
	}

	pubkey, err := route.NewVertexFromStr(createResp.Pubkey)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Second)

	mac := createResp.Macaroon

	cred, err := NewMacaroonCredential(mac)
	if err != nil {
		return nil, err
	}

	opts = append(opts,
		grpc.WithPerRPCCredentials(cred),
	)

	logger.Infow("Connecting with macaroon")

	var conn *grpc.ClientConn
	err = tryFunc(
		func() error {
			var err error
			conn, err = grpc.Dial(rpcHost, opts...)
			if err != nil {
				return err
			}
			return err
		}, 60)
	if err != nil {
		return nil, err
	}

	return &senseiConnection{
		alias:  alias,
		conn:   conn,
		pubKey: pubkey,
		client: sensei.NewNodeClient(conn),
	}, nil
}

func (l *senseiConnection) PubKey() string {
	return l.pubKey.String()
}

func (l *senseiConnection) Close() {
	l.conn.Close()
}

func (l *senseiConnection) Connect(key, host string) error {
	_, err := l.client.ConnectPeer(context.Background(), &sensei.ConnectPeerRequest{
		NodeConnectionString: fmt.Sprintf("%v@%v", key, host),
	})

	return err
}

func (l *senseiConnection) NewAddress() (string, error) {
	resp, err := l.client.GetUnusedAddress(
		context.Background(),
		&sensei.GetUnusedAddressRequest{},
	)
	if err != nil {
		return "", err
	}
	return resp.Address, nil
}

func (l *senseiConnection) OpenChannel(peerKey string, amtSat int64,
	pushAmtSat int64, private bool) error {

	var pushAmt uint64 = uint64(pushAmtSat)

	_, err := l.client.OpenChannels(context.Background(), &sensei.OpenChannelsRequest{
		Requests: []*sensei.OpenChannelRequest{
			{
				AmountSats:         uint64(amtSat),
				Public:             !private,
				CounterpartyPubkey: peerKey,
				PushAmountMsats:    &pushAmt,
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (l *senseiConnection) SetPolicy(chanPoint *lnrpc.ChannelPoint,
	policy *graphreader.PolicyData) error {

	// Not necessary for sender node.
	panic("not implemented")
}

func (l *senseiConnection) ActiveChannels() (int, error) {
	resp, err := l.client.ListChannels(context.Background(), &sensei.ListChannelsRequest{})
	if err != nil {
		return 0, err
	}

	var active = 0
	for _, ch := range resp.Channels {
		if ch.IsUsable {
			active++
		}
	}

	return active, nil
}

func (l *senseiConnection) IsSynced(totalEdges, localChannels int) (bool, error) {
	resp, err := l.client.NetworkGraphInfo(context.Background(), &sensei.NetworkGraphInfoRequest{})
	if err != nil {
		return false, err
	}

	log.Debugw("Syncing",
		"edges", resp.NumKnownEdgePolicies,
		"totalEdges", totalEdges,
		"localChannels", localChannels)

	synced := resp.NumKnownEdgePolicies == uint64(totalEdges)

	return synced, nil
}

func (l *senseiConnection) AddInvoice(amtMsat int64) (string, error) {
	// Not necessary for sender node.
	panic("not implemented")
}

func (l *senseiConnection) SendPayment(invoice string, aliasMap map[string]string) error {
	_, err := l.client.PayInvoice(context.Background(), &sensei.PayInvoiceRequest{
		Invoice: invoice,
	})
	if err != nil {
		return err
	}

	return nil
}

func (l *senseiConnection) HasFunds() error {
	for {
		resp, err := l.client.GetBalance(context.Background(),
			&sensei.GetBalanceRequest{})
		if err != nil {
			return err
		}

		if resp.OnchainBalanceSats > 0 {
			return nil
		}

		time.Sleep(time.Second)
	}
}

func (l *senseiConnection) Restart() (nodeInterface, error) {
	// Not needed for sensei to get synced correctly - hopefully.
	panic("not implemented")
}

func (l *senseiConnection) GetChannelBalance() (int, error) {
	resp, err := l.client.GetBalance(context.Background(),
		&sensei.GetBalanceRequest{})
	if err != nil {
		return 0, err
	}

	return int(resp.ChannelBalanceMsats) / 1000, nil
}

// MacaroonCredential wraps a macaroon to implement the
// credentials.PerRPCCredentials interface.
type MacaroonCredential struct {
	mac string
}

// RequireTransportSecurity implements the PerRPCCredentials interface.
func (m MacaroonCredential) RequireTransportSecurity() bool {
	return false
}

// GetRequestMetadata implements the PerRPCCredentials interface. This method
// is required in order to pass the wrapped macaroon into the gRPC context.
// With this, the macaroon will be available within the request handling scope
// of the ultimate gRPC server implementation.
func (m MacaroonCredential) GetRequestMetadata(ctx context.Context,
	uri ...string) (map[string]string, error) {

	md := make(map[string]string)
	md["macaroon"] = m.mac
	return md, nil
}

// NewMacaroonCredential returns a copy of the passed macaroon wrapped in a
// MacaroonCredential struct which implements PerRPCCredentials.
func NewMacaroonCredential(mac string) (MacaroonCredential, error) {
	ms := MacaroonCredential{mac: mac}

	return ms, nil
}
