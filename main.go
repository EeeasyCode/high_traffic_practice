package main

import (
	"flag"
	"high-traffic-practice/config"
)

var configFlag = flag.String("config", "./config.toml", "path to config file")

func main() {
	flag.Parse()

	config.NewConfig(*configFlag)
}
