package fmscore

import (
	"log"

	pb "github.com/moonsungchul/fmsgo/fmsgrpc"
)

type CmdRunEngine struct {
}

/*
 */
func (s *CmdRunEngine) CallCmd(cmd *pb.Cmd) *pb.CmdReply {
	log.Println("ip : ", cmd.GetIp())
	return &pb.CmdReply{Ip: cmd.GetIp(),
		Cmd: cmd.GetCmd(), Result: "result ok"}
}

/*
ip 등록 요청을 하면 정보를 데이터베이스에 저장한다.
*/
func (s *CmdRunEngine) RegisterNodeInfo(ip string, host string) {

}
