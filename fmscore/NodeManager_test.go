package fmscore

import (
	"log"
	"testing"
)

func TestManager(t *testing.T) {
	log.Println(">>>>>>>>>>>>>>>>>> test manager ")
	man := NewNodeManager("./fms.db")
	/*
		man.RegisterNode("192.168.0.15", "moonstar")
		man.PingHeartBeat("192.168.0.15", 1)
	*/

	man.CheckNodes()
}
