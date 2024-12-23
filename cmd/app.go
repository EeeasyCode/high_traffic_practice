package cmd

import (
	"high-traffic-practice/config"
	"high-traffic-practice/network"
	"high-traffic-practice/repository"
	"high-traffic-practice/service"
)

type App struct {
	cfg *config.Config

	service    *service.Service
	repository *repository.Repository
	network    *network.Network
}

func NewApp(cfg *config.Config) {
	a := &App{cfg: cfg}

	var err error

	if a.repository, err = repository.NewRepository(cfg); err != nil {
		panic(err)
	} else if a.service, err = service.NewService(cfg, a.repository); err != nil {
		panic(err)
	} else if a.network, err = network.NewNetwork(cfg, a.service); err != nil {
		panic(err)
	} else {
		// TODO -> start server
	}

}
