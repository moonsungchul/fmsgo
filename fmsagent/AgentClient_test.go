package main

import (
	"testing"
)

func TestAgent(t *testing.T) {
	client := NewAgentClient("192.168.0.13", "55005")
	ip2 := client.GetIpAddress()
	print("ip2 : ", ip2, "\n")
}
