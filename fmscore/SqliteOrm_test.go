package fmscore

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert := assert.New(t)
	store := NewSqliteOrm("test.db")
	db, err := store.Open()
	if err != nil {
		log.Println("open error : ", err)
	}
	store.Migrate(db)

	db.Create(&NodeInfo{IP: "192.168.0.12", HostName: "test1", Status: 1})
	db.Create(&CmdResult{IP: "192.168.0.12", Cmd: "Test", Result: "Result1"})

	var nodeInfo NodeInfo
	var cmdResult CmdResult
	db.First(&nodeInfo, "IP = ?", "192.168.0.12")
	db.First(&cmdResult, "IP = ?", "192.168.0.12")

	assert.Equal(nodeInfo.IP, "192.168.0.12", "두 IP가 같아야 한다.")
	assert.Equal(nodeInfo.IP, "192.168.0.12", "두 IP가 같아야 한다.")

	defer db.Close()
}
