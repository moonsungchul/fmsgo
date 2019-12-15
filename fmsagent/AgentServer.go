package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/moonsungchul/fmsgo/fmscore"
	pb "github.com/moonsungchul/fmsgo/fmsgrpc"
	grpc "google.golang.org/grpc"
)

const (
	port = ":60050"
)

type Agent struct {
	pb.UnimplementedFmsRpcServiceServer
	Manager    *fmscore.NodeManager
	ServerIP   string
	ServerPort string
}

func NewAgent() *Agent {
	man := fmscore.NewNodeManager("./agent.db")
	return &Agent{Manager: man}
}

func main() {
	serverIP := flag.String("server", "localost", "Server ip ")
	serverPort := flag.String("server_port", "50051", "Server port")
	flag.Parse()

	log.Println("server ip : ", *serverIP)
	log.Println("server prot  : ", *serverPort)

	client := NewAgentClient(*serverIP, *serverPort)

	con, err := client.OpenConn()
	if err != nil {
		log.Println("error ")
	}
	defer con.Close()
	log.Println("test ................>>>>>. : ")

	go func() {
		client.RegisterNodeInfo("192.168.0.13", "moonstar")
		for {
			fmt.Println("time loop ")
			time.Sleep(30 * time.Second)
		}
	}()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}

	server := NewAgent()
	server.ServerIP = *serverIP
	server.ServerPort = port

	s := grpc.NewServer()
	pb.RegisterFmsRpcServiceServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to server : %v", err)
	}
	fmt.Println("server start ...")

}
