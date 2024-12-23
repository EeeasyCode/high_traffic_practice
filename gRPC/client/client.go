package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"high-traffic-practice/config"
	"high-traffic-practice/gRPC/paseto"
	auth "high-traffic-practice/gRPC/proto"
)

type GRPCClient struct {
	client      *grpc.ClientConn
	authClient  auth.AuthServiceClient
	pasetoMaker *paseto.PasetoMaker
}

func newGRPCClient(cfg *config.Config) (*GRPCClient, error) {
	c := new(GRPCClient)

	if client, err := grpc.NewClient(cfg.GRPC.URL, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, err
	} else {
		c.client = client
		c.authClient = auth.NewAuthServiceClient(c.client)
		c.pasetoMaker = paseto.NewPasetoMaker(cfg)
	}
	return c, nil
}

func (g *GRPCClient) CreateAuth(address string) (*auth.AuthData, error) {
	return nil, nil
}

func (g *GRPCClient) verifyAuth(token string) (*auth.ValidateTokenResponse, error) {
	return nil, nil
}
