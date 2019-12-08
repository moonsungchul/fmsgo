package fmscore

import (
	"fmt"
	"log"
	"net"
	"os"
)

type NetworkUtil struct {
}

func NewNetworkUtil() *NetworkUtil {
	return &NetworkUtil{}
}

func (s *NetworkUtil) getHostIp() string {
	host, _ := os.Hostname()
	log.Println("host : ", host)
	addrs, _ := net.LookupIP(host)
	log.Println("Addrs : ", addrs)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Println("IPv4 : ", ipv4)
		}
	}
	return ""
}
