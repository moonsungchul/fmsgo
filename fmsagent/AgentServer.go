package main

import (
	"github.com/moonsungchul/fmsgo/fmscore"
	pb "github.com/moonsungchul/fmsgo/fmsgrpc"
)

const (
	port = ":60050"
)

type Agent struct {
	pb.UnimplementedFmsRpcServiceServer
	Manager *fmscore.NodeManager
}

func NewAgent() *Agent {
	man := fmscore.NewNodeManager("./agent.db")
	return &Agent{Manager: man}
}
