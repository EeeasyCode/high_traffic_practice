package main

import (
	"flag"
	"high-traffic-practice/cmd"
	"high-traffic-practice/config"
	"high-traffic-practice/gRPC/server"
	"time"
)

var configFlag = flag.String("config", "./config.toml", "path to config file")

func main() {
	flag.Parse()
	cfg := config.NewConfig(*configFlag)

	if err := server.NewGRPCServer(cfg); err != nil {
		panic(err)
	} else {
		time.Sleep(1e9)
		cmd.NewApp(cfg)
	}
}
