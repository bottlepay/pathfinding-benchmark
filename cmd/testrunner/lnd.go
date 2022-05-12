package main

import (
	"context"
	"time"

	"github.com/bottlepay/lightning-benchmark/graphreader"
	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/lnrpc/routerrpc"
	"github.com/lightningnetwork/lnd/routing/route"
	"google.golang.org/grpc"
)

type lndConnection struct {
	conn            *grpc.ClientConn
	routerClient    routerrpc.RouterClient
	lightningClient lnrpc.LightningClient
	pubKey          route.Vertex
	alias           string
}

func tryFunc(f func() error, maxAttempts int) error {
	var attempts int
	for {
		err := f()
		if err == nil {
			return err
		}

		attempts++
		if attempts == maxAttempts {
			return err
		}

		time.Sleep(time.Second)
	}
}

func getLndConnection(alias string) (*lndConnection, error) {
	logger := log.With("node", alias)

	var conn *grpc.ClientConn
	err := tryFunc(
		func() error {
			var err error
			conn, err = getClientConn(alias)
			return err
		}, 10)
	if err != nil {
		return nil, err
	}

	senderClient := lnrpc.NewLightningClient(conn)

	logger.Infow("Attempting to connect to lnd")
	for {
		resp, err := senderClient.GetInfo(context.Background(), &lnrpc.GetInfoRequest{})
		if err == nil {
			if !resp.SyncedToChain {
				time.Sleep(time.Second)

				continue
			}

			logger.Infow("Connected to lnd", "key", resp.IdentityPubkey)
			pubKey, err := route.NewVertexFromStr(resp.IdentityPubkey)
			if err != nil {
				return nil, err
			}

			return &lndConnection{
				pubKey:          pubKey,
				conn:            conn,
				routerClient:    routerrpc.NewRouterClient(conn),
				lightningClient: lnrpc.NewLightningClient(conn),
				alias:           alias,
			}, nil
		}

		logger.Errorw("Lnd connection error", "error", err)

		time.Sleep(time.Second)
	}
}

func (l *lndConnection) PubKey() string {
	return l.pubKey.String()
}

func (l *lndConnection) Close() {
	l.conn.Close()
}

func (l *lndConnection) GetInfo() (*info, error) {
	infoResp, err := l.lightningClient.GetInfo(context.Background(), &lnrpc.GetInfoRequest{})
	if err != nil {
		return nil, err
	}

	return &info{
		key:    infoResp.IdentityPubkey,
		synced: infoResp.SyncedToChain,
	}, nil
}

func (l *lndConnection) Connect(key, host string) error {
	_, err := l.lightningClient.ConnectPeer(context.Background(), &lnrpc.ConnectPeerRequest{
		Addr: &lnrpc.LightningAddress{
			Host:   host,
			Pubkey: key,
		},
	})
	return err
}

func (l *lndConnection) NewAddress() (string, error) {
	addrResp, err := l.lightningClient.NewAddress(context.Background(), &lnrpc.NewAddressRequest{
		Type: lnrpc.AddressType_WITNESS_PUBKEY_HASH,
	})
	if err != nil {
		return "", err
	}

	return addrResp.Address, nil
}

func (l *lndConnection) OpenChannel(peerKey string, amtSat int64, pushAmtSat int64) (
	*lnrpc.ChannelPoint, error) {

	resp, err := l.lightningClient.OpenChannelSync(context.Background(), &lnrpc.OpenChannelRequest{
		LocalFundingAmount: amtSat,
		NodePubkeyString:   peerKey,
		SpendUnconfirmed:   true,
		PushSat:            pushAmtSat,
	})
	if err != nil {
		return nil, err
	}

	chanPoint := lnrpc.ChannelPoint{
		FundingTxid: &lnrpc.ChannelPoint_FundingTxidBytes{
			FundingTxidBytes: resp.GetFundingTxidBytes(),
		},
		OutputIndex: resp.OutputIndex,
	}
	return &chanPoint, nil
}

func (l *lndConnection) SetPolicy(chanPoint *lnrpc.ChannelPoint,
	policy *graphreader.PolicyData) error {

	_, err := l.lightningClient.UpdateChannelPolicy(context.Background(),
		&lnrpc.PolicyUpdateRequest{
			Scope: &lnrpc.PolicyUpdateRequest_ChanPoint{
				ChanPoint: chanPoint,
			},
			BaseFeeMsat:   int64(policy.BaseFee),
			FeeRate:       float64(policy.FeeRate) / 1e6,
			TimeLockDelta: uint32(policy.CltvDelta),
			MaxHtlcMsat:   uint64(policy.HtlcMaxSat) * 1e3,
		},
	)
	return err
}

func (l *lndConnection) ActiveChannels() (int, error) {
	resp, err := l.lightningClient.ListChannels(context.Background(), &lnrpc.ListChannelsRequest{
		ActiveOnly: true,
	})
	if err != nil {
		return 0, err
	}
	return len(resp.Channels), nil
}

func (l *lndConnection) NetworkEdgeCount() (int, error) {
	resp, err := l.lightningClient.DescribeGraph(context.Background(), &lnrpc.ChannelGraphRequest{})
	if err != nil {
		return 0, err
	}

	edges := 0
	for _, edge := range resp.Edges {
		// 999 is the default fee. If that's the policy, it means that the
		// latest channel update hasn't been received yet.
		if edge.Node1Policy != nil && edge.Node1Policy.FeeBaseMsat != 999 {
			edges++
		}
		if edge.Node2Policy != nil && edge.Node2Policy.FeeBaseMsat != 999 {
			edges++
		}
	}

	return edges, nil
}

func (l *lndConnection) AddInvoice(amtMsat int64) (string, error) {
	addResp, err := l.lightningClient.AddInvoice(context.Background(), &lnrpc.Invoice{
		ValueMsat: amtMsat,
	})
	if err != nil {
		return "", err
	}
	return addResp.PaymentRequest, nil
}

func (l *lndConnection) SendPayment(invoice string, aliasMap map[string]string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := l.routerClient.SendPayment(ctx, &routerrpc.SendPaymentRequest{
		FeeLimitSat:    1e8,
		PaymentRequest: invoice,
		TimeoutSeconds: 60,
	})
	if err != nil {
		return err
	}

	prevHtlcs := 0
	for {
		update, err := stream.Recv()
		if err != nil {
			return err
		}

		htlcCount := len(update.Htlcs)
		for i := prevHtlcs; i < htlcCount; i++ {

			route := update.Htlcs[i].Route
			var hops []string
			for _, hop := range route.Hops {
				hops = append(hops, aliasMap[hop.PubKey])
			}
			log.Debugw("Payment update", "htlcIdx", i, "route", hops)
		}
		prevHtlcs = htlcCount

		if update.State == routerrpc.PaymentState_SUCCEEDED {
			return nil
		}
	}
}

func (l *lndConnection) HasFunds() error {
	for {
		resp, err := l.lightningClient.WalletBalance(context.Background(), &lnrpc.WalletBalanceRequest{})
		if err != nil {
			return err
		}
		if resp.ConfirmedBalance > 0 {
			return nil
		}

		time.Sleep(time.Second)
	}
}

func (l *lndConnection) Restart() (nodeInterface, error) {
	_, err := l.lightningClient.StopDaemon(context.Background(), &lnrpc.StopRequest{})
	if err != nil {
		return nil, err
	}
	conn, err := getLndConnection(l.alias)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
