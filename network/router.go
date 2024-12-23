package network

import (
	"high-traffic-practice/config"
	"high-traffic-practice/service"
)

type Network struct {
	cfg *config.Config

	service *service.Service
}

func NewNetwork(cfg *config.Config, service *service.Service) (*Network, error) {
	r := &Network{cfg: cfg, service: service}

	return r, nil
}
