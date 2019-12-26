package fmscore

import (
	"log"
	"testing"
	"time"
)

func TestUtil(t *testing.T) {
	log.Println(">>>>>>>>>>>>>>>>>> Date Util test")

	now := time.Now()
	log.Println("now : ", now)
	log.Println("unix sec : ", now.Unix())

	time.Sleep(2 * time.Second)
	util := DateUtil{}
	dif := util.GetDifNow(now.Unix())

	log.Println("dif : ", dif)
}
