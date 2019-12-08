package main

import (
	"context"
	"log"
	"net"

	"github.com/moonsungchul/fmsgo/fmscore"
	pb "github.com/moonsungchul/fmsgo/fmsgrpc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedFmsRpcServiceServer
	Manager *fmscore.NodeManager
}

/*
생성자
*/
func NewServer() *server {
	man := fmscore.NewNodeManager("./fms.db")
	return &server{Manager: man}
}

/*
Client에서 온 노드 정보를 등록 요청을 처리한다.
*/
func (s *server) RegNodeInfo(ctx context.Context, in *pb.NodeInfo) (*pb.RetMsg, error) {
	log.Println("input ip :", in.GetIp())
	res, err := s.Manager.RegisterNode(in.GetIp(), in.GetHostName())
	if err != nil {
		return nil, err
	}
	return &pb.RetMsg{MsgNo: 1, Msg: res}, nil
}

func (s *server) PingHeartBeat(ctx context.Context, in *pb.HeartbeatMsg) (*pb.HeartbeatReply, error) {
	log.Printf("Received : %v", in.GetIp())
	return &pb.HeartbeatReply{Ret: "im live"}, nil
}

func (s *server) CallCmd(ctx context.Context, in *pb.Cmd) (*pb.CmdReply, error) {
	log.Printf("received : %v", in.GetIp())
	return &pb.CmdReply{Ip: in.GetIp(), Cmd: in.GetCmd(), Result: "test"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	sserver := NewServer()
	s := grpc.NewServer()
	pb.RegisterFmsRpcServiceServer(s, sserver)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("faile to serer : %v", err)
	}

}
