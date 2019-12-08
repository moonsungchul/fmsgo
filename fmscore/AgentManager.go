package fmscore

type AgentManager struct {
	DBStore  *sqlitestore
	ServerIp string
}

func NewAgentManager(fname string, serverIp string) *AgentManager {
	store := &sqlitestore{DbFile: "./agent.db"}
	return &AgentManager{DBStore: store, ServerIp: serverIp}
}

/*
자신의 노드 정보를 서버에 저장한다.
*/
//func (s *AgentManager) RegNodeInfo() (string, error) {

//}
