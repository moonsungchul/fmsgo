package fmscore

import (
	"log"
	"testing"
)

func TestNetwork(t *testing.T) {
	util := NewNetworkUtil()
	ip := util.getHostIp()
	log.Println("ip : ", ip)
}
