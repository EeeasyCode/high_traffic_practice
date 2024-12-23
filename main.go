package main

import (
	"flag"
	"high-traffic-practice/cmd"
	"high-traffic-practice/config"
)

var configFlag = flag.String("config", "./config.toml", "path to config file")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*configFlag)

	cmd.NewApp(cfg)
}
