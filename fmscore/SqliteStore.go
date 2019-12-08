package fmscore

import (
	"container/list"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	RUNNING = 1
	STOP    = 0
)

type sqlitestore struct {
	DbFile string // Sqlite3 db file name
}

type nodeinfo struct {
	IP       string `json:"ip"`
	HostName string `json:"host_name"`
	Status   int    `json:"status"`
}

type cmdresult struct {
	IP     string `json:"ip"`
	Cmd    string `json:"cmd"`
	Result string `json:"Result"`
}

/*
Sqlite3 dbfile을 연다.
*/
func (r *sqlitestore) Open() (*sql.DB, error) {
	return sql.Open("sqlite3", r.DbFile)
}

/*
새로운 테이블을 생성한다.
*/
func (r *sqlitestore) CreateTable() {
	db, err := r.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `
		create table if not exists NODE_INFO (
			ip char(16) primary key, 
			host_name varchar(100), 
			status int 
		);
		create table if not exists CMD_RESULT(
			ip char(16) primary key, 
			cmd varchar(1000),
			result varchar(10000)
		);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *sqlitestore) insertCmdResult(ip string, cmd string, result string) {
	db, err := r.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	sql := `insert into CMD_RESULT (ip, cmd, result) values (?, ?, ? )`
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(ip, cmd, result)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

/*
노드 정보를 저장한다.
*/
func (r *sqlitestore) insertNodeInfo(ip string, host string, status int) {

	log.Println(">>>> ip : ", ip)
	log.Println(">>>> host : ", host)
	log.Println(">>>> status : ", status)

	db, err := r.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	sql := `insert into NODE_INFO (ip, host_name, status) values (?, ?, ? )`
	stmt, err := tx.Prepare(sql)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(ip, host, status)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func (r *sqlitestore) getCmdResult(ip string) *list.List {
	db, err := r.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select ip,cmd, result from CMD_RESULT where ip = $1", ip)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var sip string
	var cmd string
	var result string
	rets := list.New()

	for rows.Next() {
		err := rows.Scan(&sip, &cmd, &result)
		log.Println("ipc:", ip)
		log.Println("cmd", cmd)
		log.Println("result", result)
		if err != nil {
			log.Fatal(err)
		}
		cmdresult := &cmdresult{IP: sip, Cmd: cmd, Result: result}
		rets.PushBack(cmdresult)
	}
	return rets
}

func (r *sqlitestore) getNodeInfo(ip string) (*nodeinfo, error) {
	db, err := r.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows := db.QueryRow("select ip, host_name, status from NODE_INFO where ip = $1", ip)
	var sip string
	var shost string
	var status int
	err = rows.Scan(&sip, &shost, &status)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	ninfo := &nodeinfo{IP: sip, HostName: shost, Status: status}
	return ninfo, nil
}
