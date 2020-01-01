package fmscore

import (
	"crypto/tls"
	"log"
	"net"
	"time"

	ignite "github.com/amsokol/ignite-go-client/binary/v1"
)

type Ignite struct {
}

func (s *Ignite) Open() *ignite.Client {
	coninfo := ignite.ConnInfo{
		Network:  "tcp",
		Host:     "localhost",
		Port:     10800,
		Major:    1,
		Minor:    1,
		Patch:    0,
		Username: "ignite",
		Password: "ignite",
		Dialer: net.Dialer{
			Timeout: 10 * time.Second,
		},
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	c, err := ignite.Connect(coninfo)
	if err != nil {
		log.Fatalln("failed connect to server : %v", err)
	}
	return &c
}

func (s *Ignite) Execute(cl ignite.Client, cache string,
	sql string, arg []interface{}, pageSize int) (ignite.QuerySQLFieldsResult, error) {
	res, err := cl.QuerySQLFields(cache, false, ignite.QuerySQLFieldsData{
		PageSize:  pageSize,
		Query:     sql,
		QueryArgs: arg,
	})
	if err != nil {
		log.Fatalf("failed insert data: %v", err)
	}
	return res, err
}

func (s *Ignite) Select(cl ignite.Client, cache string, table string,
	sql string, pageSize int) (ignite.QuerySQLResult, error) {
	res, err := cl.QuerySQL(cache, false, ignite.QuerySQLData{
		Table:    table,
		Query:    sql,
		PageSize: pageSize,
	})
	return res, err
}
