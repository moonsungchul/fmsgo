package fmscore

import (
	"log"

	"github.com/jinzhu/gorm"
)

type NodeInfo struct {
	gorm.Model
	IP       string
	HostName string
	Status   int   // 1 : 서버 정상 실행 , 0 : 서버 죽어 있음.
	HeatBeat int   // 초기 60 * 5 = 300 초 5분 동안 응답이 없으면 죽었다고 판단
	LastTime int64 // 마지막으로 보내온 시간 unix time
}

type CmdResult struct {
	gorm.Model
	IP     string
	Cmd    string
	Result string
}

type MemoryInfo struct {
	gorm.Model
	Total      uint64
	Free       uint64
	UsePercent uint64
}

type CPUInfo struct {
	gorm.Model
	phyCPU   int
	logicCPU int
}

type DiskInfo struct {
	gorm.Model
	Total       uint64
	Free        uint64
	Used        uint64
	UsedPercent uint64
}

type SqliteOrm struct {
	DbFile string
}

func NewSqliteOrm(fname string) *SqliteOrm {
	log.Println(">>>>>>>>>>>>>>>>>>>>>>> db file : ", fname)
	return &SqliteOrm{DbFile: fname}
}

func (r *SqliteOrm) Open() (*gorm.DB, error) {
	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>> test open", r.DbFile)
	db, err := gorm.Open("sqlite3", r.DbFile)
	if err != nil {
		panic("failed to connect database")
	}
	return db, err
}

func (r *SqliteOrm) Migrate(db *gorm.DB) {
	db.AutoMigrate(&NodeInfo{})
	db.AutoMigrate(&CmdResult{})
}
