package fmscore

import (
	"log"
	"testing"
)

func TestUtil(t *testing.T) {
	log.Println(">>>>>>>>>>>>>>>>>> Date Util test")
	util := DateUtil{}
	dif := util.GetDif("2019-12-26 00:30:00")
	log.Println("dif second : ", dif)

}
