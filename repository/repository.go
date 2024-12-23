package repository

import (
	"high-traffic-practice/config"
	"high-traffic-practice/gRPC/client"
	auth "high-traffic-practice/gRPC/proto"
)

type Repository struct {
	cfg *config.Config

	gRPCClient *client.GRPCClient
}

func NewRepository(cfg *config.Config, gRPCClient *client.GRPCClient) (*Repository, error) {
	r := &Repository{cfg: cfg, gRPCClient: gRPCClient}

	return r, nil
}

func (r *Repository) CreateAuth(username string) (*auth.AuthData, error) {
	return r.gRPCClient.CreateAuth(username)
}
