package main

import (
	"flag"

	"github.com/atmom/flowcenter/common"
	"github.com/atmom/flowcenter/config"
	"github.com/atmom/flowcenter/server"
)

var (
	configPath = flag.String("config", "", "The flow center server configuration path")
)

func main() {
	// parse command flags
	flag.Parse()

	// load configuration from json file.
	c := config.Load(*configPath)

	// init log.
	common.InitLog(c)

	// start flow center server.
	server.Start(c)
}
