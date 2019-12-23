package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	pb "github.com/moonsungchul/fmsgo/fmsgrpc"
	grpc "google.golang.org/grpc"
)

// AgentClient : 연결 구조체
// ServerIP : string
// Port : string
// ClientConn : string
type AgentClient struct {
	ServerIP   string
	Port       string
	ClientConn *grpc.ClientConn
}

// NewAgentClient :  AgentClient를  생성한다.
// serverIp : string
// port : string
// return : AgentClient 포인터
func NewAgentClient(serverIP string, port string) *AgentClient {
	return &AgentClient{ServerIP: serverIP, Port: port}
}

// OpenConn : RPC Connection 열고 리턴한다.
func (s *AgentClient) OpenConn() (*grpc.ClientConn, error) {
	address := s.ServerIP + ":" + s.Port
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	log.Println("test conn :", conn)
	if err != nil {
		return nil, err
	}
	s.ClientConn = conn
	return conn, err
}

// RegisterNodeInfo : 자신의 ip, hostname을 서버에 등록한다.
func (s *AgentClient) RegisterNodeInfo(ip string, hostname string) (string, error) {
	c := pb.NewFmsRpcServiceClient(s.ClientConn)
	ctx, canncel := context.WithTimeout(context.Background(), time.Second)
	defer canncel()
	r, err := c.RegNodeInfo(ctx, &pb.NodeInfo{Ip: ip, HostName: hostname})
	if err != nil {
		return "node register error ", err
	}
	return r.GetMsg(), err
}

// PingHeartbeat : heartbeat을 전송한다.
func (s *AgentClient) PingHeartbeat() (string, error) {
	c := pb.NewFmsRpcServiceClient(s.ClientConn)
	ctx, canncel := context.WithTimeout(context.Background(), time.Second)
	defer canncel()
	mm := pb.HeartbeatMsg{Ip: "192.168.0.13"}
	r, err := c.PingHeartBeat(ctx, &mm)
	if err != nil {
		return "node register error ", err
	}
	return r.GetRet(), err
}

func (s *AgentClient) GetIpAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
		//if (strings.Contains(addr.String(), "127.0.0.1") == false)
		if strings.Contains(addr.String(), "127.0.0.1") == false && strings.Contains(addr.String(), "::") == false {
			ss := strings.Split(addr.String(), "/")
			return ss[0]
		}
	}
	return ""
}
