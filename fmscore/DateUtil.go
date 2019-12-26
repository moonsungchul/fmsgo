package fmscore

import (
	"time"
)

type DateUtil struct {
}

// 이전 시간 부터 오늘 시간까지의 차리를 리턴 한다.
func (s *DateUtil) GetDifNow(uinx_time int64) int64 {
	ut := time.Now().Unix()
	return ut - uinx_time
}

// end_time - start_time 의 차를 구한하다. 두 시간은 unix time이다.
func (s *DateUtil) GetDif(start_time int64, end_time int64) int64 {
	return end_time - start_time
}
