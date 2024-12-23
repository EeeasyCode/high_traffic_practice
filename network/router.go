package network

import (
	"github.com/gin-gonic/gin"
	"high-traffic-practice/config"
	"high-traffic-practice/gRPC/client"
	"high-traffic-practice/service"
)

type Network struct {
	cfg *config.Config

	service *service.Service

	gRPCClient *client.GRPCClient

	engin *gin.Engine
}

func NewNetwork(cfg *config.Config, service *service.Service, gRPCClient *client.GRPCClient) (*Network, error) {
	r := &Network{cfg: cfg, service: service, engin: gin.New(), gRPCClient: gRPCClient}

	return r, nil
}

func (n *Network) StartServer() {
	n.engin.Run(":8080")
}
