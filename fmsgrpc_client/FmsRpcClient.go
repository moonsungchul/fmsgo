package main

import (
	"context"
	"log"
	"time"

	pb "github.com/moonsungchul/fmsgo/fmsgrpc"
	grpc "google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("dit not connect : %v", err)
	}
	defer conn.Close()
	c := pb.NewHeartBeatClient(conn)

	ctx, canncel := context.WithTimeout(context.Background(), time.Second)
	defer canncel()
	r, err := c.PingHeartBeat(ctx, &pb.HeartbeatMsg{Ip: "192.168.0.13"})
	if err != nil {
		log.Fatalf("cound not greet :%v", err)
	}
	log.Printf("Greeting: %s", r.GetRet())

}
