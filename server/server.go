package server

import (
	"fmt"
	"net"

	"github.com/atmom/flowcenter/config"
	flowcenter "github.com/atmom/flowcenter/grpc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// start server.
func Start(config *config.FlowCenterConfig) {
	if config.Port < 1024 || config.Port > 65536 {
		logrus.Fatalf("Flow center server start error, invalid server port %d.", config.Port)
	}

	// listen tcp.
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logrus.Fatalf("Flow center server start error %v.", err)
	}

	// create server.
	s := grpc.NewServer()
	flowcenter.RegisterFlowCenterServer(s, &service{Config: config})

	logrus.Infof("Flow center server start success, listening at %v", listen.Addr())

	if err := s.Serve(listen); err != nil {
		logrus.Fatalf("Flow center server start error %v.", err)
	}
}
