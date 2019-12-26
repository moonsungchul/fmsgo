package fmscore

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

const (
	HeartBeatTime  = 30  // HeartBeat  체크 주기
	heartBeatLimit = 300 // heart beat 체크 후 300초 동안 지나면 상태 아웃으로 판단
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
	store := &SqliteOrm{DbFile: fname}
	con, err := store.Open()
	if err != nil {
		log.Println("fms.db 내부 데이터베이스를 열수 없습니다. ")
	}
	store.Migrate(con)
	return &NodeManager{dbstore: store, con: con}
}

//IP가 이미 저장하고 있을 때 에러 메시지
var ExistsIpError = errors.New("Exists ip")
var NotNodeError = errors.New("Node Found not")

//node 정보를 데이터베이스에 저장한다.
func (s *NodeManager) RegisterNode(ip string, hostName string) (string, error) {
	log.Println("ip :", ip)
	log.Println("host_name : ", hostName)
	var nodeinfo NodeInfo
	var co int
	s.con.Where("IP = ?", ip).Find(&nodeinfo).Count(&co)

	if co > 0 {
		msg := fmt.Sprintf("%s는 이미 있는 IP입니다.", ip)
		return msg, ExistsIpError
	}
	nodeinfo.IP = ip
	nodeinfo.HostName = hostName
	nodeinfo.Status = 1
	nodeinfo.HeatBeat = 300
	nodeinfo.LastTime = time.Now().Unix()
	s.con.Create(&nodeinfo)
	return "새로운 데이터 생성", nil
}

// Agent에서 온 ping 신호를 처리한다.
func (s *NodeManager) PingHeartBeat(ip string, status int) (string, error) {
	// ip에 해당하는 노드 정보를 읽어온다.
	var nodeinfo NodeInfo
	var co int
	s.con.Where("IP = ?", ip).Find(&nodeinfo).Count(&co)
	if co == 0 {
		msg := fmt.Sprintf("%s 주소를 갖는 노드를 찾을 수 없습니다.", ip)
		return msg, NotNodeError
	}
	s.con.Where("IP = ?", ip).Find(&nodeinfo)
	// 최종 접속 시간을 갱신한다.
	nodeinfo.LastTime = time.Now().Unix()
	nodeinfo.Status = status
	s.con.Save(&nodeinfo)
	return "Update ok", nil
}
