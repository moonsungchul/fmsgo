package fmscore

import (
	"errors"
	"fmt"
	"log"
)

/*
노드의 정보를 관리한다.
*/
type NodeManager struct {
	dbstore *sqlitestore
}

/*
생성자
*/
func NewNodeManager(fname string) *NodeManager {
	store := &sqlitestore{DbFile: "./fms.db"}
	return &NodeManager{dbstore: store}
}

/*
IP가 이미 저장하고 있을 때 에러 메시지
*/
var ExistsIpError = errors.New("Exists ip")

/*
node 정보를 데이터베이스에 저장한다.
*/
func (s *NodeManager) RegisterNode(ip string, hostName string) (string, error) {
	log.Println("ip :", ip)
	log.Println("host_name : ", hostName)
	ninfo, err := s.dbstore.getNodeInfo(ip)
	if err != nil {
		log.Println(err)
		return "error", err
	}
	if ninfo != nil {
		msg := fmt.Sprintf("%s는 이미 있는 IP입니다.", ip)
		return msg, ExistsIpError
	}
	s.dbstore.insertNodeInfo(ip, hostName, 1)
	return "새로운 데이터 생성", nil
}
