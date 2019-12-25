package fmscore

import (
	"errors"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

/*
노드의 정보를 관리한다.
*/
type NodeManager struct {
	dbstore *SqliteOrm
	con     *gorm.DB
}

//생성자
func NewNodeManager(fname string) *NodeManager {
	//store := &sqlitestore{DbFile: "./fms.db"}
	store := &SqliteOrm{DbFile: "./fms.db"}
	con, err := store.Open()
	if err != nil {
		log.Println("fms.db 내부 데이터베이스를 열수 없습니다. ")
	}
	store.Migrate(con)
	return &NodeManager{dbstore: store, con: con}
}

//IP가 이미 저장하고 있을 때 에러 메시지
var ExistsIpError = errors.New("Exists ip")

//node 정보를 데이터베이스에 저장한다.
func (s *NodeManager) RegisterNode(ip string, hostName string) (string, error) {
	log.Println("ip :", ip)
	log.Println("host_name : ", hostName)
	var nodeinfo NodeInfo
	var co int
	s.con.Where("IP = ?").Find(&nodeinfo).Count(&co)

	if co > 0 {
		msg := fmt.Sprintf("%s는 이미 있는 IP입니다.", ip)
		return msg, ExistsIpError
	}
	nodeinfo.IP = ip
	nodeinfo.HostName = hostName
	nodeinfo.Status = 1
	nodeinfo.HeatBeat = 300
	s.con.Create(&nodeinfo)
	return "새로운 데이터 생성", nil
}

// Agent에서 온 ping 신호를 저장한다.
//func (s *NodeManager) PingHeartBeat(ip string, status int) (string, error) {

//}
