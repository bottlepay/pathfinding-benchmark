package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lightningnetwork/lnd/build"
	"github.com/lightningnetwork/lnd/macaroons"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	macaroon "gopkg.in/macaroon.v2"
)

func getClientConn(alias string) (*grpc.ClientConn, error) {
	var (
		tlsCertPath = fmt.Sprintf("/lnd/%v/tls.cert", alias)
		macPath     = fmt.Sprintf("/lnd/%v/admin.macaroon", alias)
		rpcHost     = fmt.Sprintf("lnd_%v:10009", alias)
	)

	creds, err := credentials.NewClientTLSFromFile(tlsCertPath, "")
	if err != nil {
		return nil, err
	}

	macBytes, err := ioutil.ReadFile(macPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read macaroon path (check "+
			"the network setting!): %v", err)
	}

	mac := &macaroon.Macaroon{}
	if err = mac.UnmarshalBinary(macBytes); err != nil {
		return nil, fmt.Errorf("unable to decode macaroon: %v", err)
	}

	cred, err := macaroons.NewMacaroonCredential(mac)
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(cred),
	}

	conn, err := grpc.Dial(rpcHost, opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to RPC server: %v", err)
	}

	return conn, nil
}

func main() {
	app := cli.NewApp()
	app.Name = "load test"
	app.Version = build.Version() + " commit=" + build.Commit
	app.Commands = []cli.Command{runCommand}
	if err := app.Run(os.Args); err != nil {
		log.Errorw("Exiting", "err", err)
		os.Exit(1)
	}
}
