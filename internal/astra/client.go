package astra

import (
	"crypto/tls"
	"errors"
	"time"

	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/auth"
	"github.com/stargate/stargate-grpc-go-client/stargate/pkg/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	uri           string
	token         string
	config        *tls.Config
	conn          *grpc.ClientConn
	sgClient      *client.StargateClient
	keyspace      string
	countersTable string
}

func NewClient(uri, token, keyspace string) (*Client, error) {

	if uri == "" {
		return nil, errors.New("uri not specified")
	}

	if token == "" {
		return nil, errors.New("token not specified")
	}

	if keyspace == "" {
		return nil, errors.New("keyspace not specified")
	}

	c := Client{
		uri:   uri,
		token: token,
		config: &tls.Config{
			InsecureSkipVerify: false,
		},
		keyspace:      keyspace,
		countersTable: "counters",
	}

	conn, err := grpc.Dial(c.uri, grpc.WithTransportCredentials(credentials.NewTLS(c.config)),
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(
			auth.NewStaticTokenProvider(c.token),
		),
	)
	if err != nil {
		return nil, err
	}

	c.conn = conn

	stargateClient, err := client.NewStargateClientWithConn(conn, client.WithTimeout(5*time.Second))
	if err != nil {
		return nil, err
	}

	c.sgClient = stargateClient

	err = c.CreateCountersTable()
	if err != nil {
		return nil, err
	}

	return &c, nil
}
