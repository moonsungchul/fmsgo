package fmscore

import (
	"log"
	"time"
)

type DateUtil struct {
}

// 두 시간의 차이를 초로 리턴한다.
func (s *DateUtil) GetDif(sdate string) float64 {
	layout := "2006-01-02 15:04:05"
	a, err := time.Parse(layout, sdate)
	if err != nil {
		log.Println("error : ", err)
		return 0
	}
	//loc, _ := time.LoadLocation("Asia/Seoul")
	//kst := a.In(loc)

	log.Println(">>>>> a: ", a)
	log.Println(">>>>> now: ", time.Now())

	dif := time.Now().Sub(a).Hours()
	log.Println("dif : ", dif)
	return dif
}
