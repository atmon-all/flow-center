package server

import (
	"context"

	"github.com/atmom/flowcenter/config"
	flowcenter "github.com/atmom/flowcenter/grpc"
	"github.com/sirupsen/logrus"
)

type service struct {
	flowcenter.UnimplementedFlowCenterServer
	Config *config.FlowCenterConfig
}

// Flow node server poll flows from flow center.
// Report the current flow nodes in the current flow node server,
// and get the next nodes of them.
func (service *service) Poll(ctx context.Context, request *flowcenter.FlowPollRequest) (*flowcenter.FlowPollResponse, error) {
	logrus.Infof("Received: %v", request)
	return &flowcenter.FlowPollResponse{Code: 0, Message: "success"}, nil
}
