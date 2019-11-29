package main

import (
	"context"
	"log"
	"net"

	pb "github.com/moonsungchul/fmsgo/fmsgrpc"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedHeartBeatServer
}

func (s *server) PingHeartBeat(ctx context.Context, in *pb.HeartbeatMsg) (*pb.HeartbeatReply, error) {
	log.Printf("Received : %v", in.GetIp())
	return &pb.HeartbeatReply{Ret: "im live"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHeartBeatServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("faile to serer : %v", err)
	}

}
