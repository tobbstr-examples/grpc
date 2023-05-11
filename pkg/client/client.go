package client

import (
	"context"

	pb "github.com/tobbstr-examples/grpc/internal/pb/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	pb.AuditerClient
}

func NewGrpcClient(ctx context.Context, url, clientName string) (*Client, func() error, error) {
	conn, err := grpc.DialContext(
		ctx,
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithUserAgent(clientName),
	)
	if err != nil {
		return nil, nil, err
	}

	pbCli := pb.NewAuditerClient(conn)
	return &Client{AuditerClient: pbCli}, conn.Close, nil
}
