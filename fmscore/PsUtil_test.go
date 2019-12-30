package fmscore

import (
	"log"
	"testing"
)

func TestPsUtil(t *testing.T) {
	util := NewPsUtil()
	smem := util.GetMemory()
	log.Println("smem : ", smem)

	phyCPU, logicCPU := util.GetCPU()
	log.Println("scpu : ", phyCPU, logicCPU)

	cpu := util.GetCPUInfo()
	log.Println("cpu info : ", cpu)

	dis := util.GetDiskInfo()
	log.Println("dis : ", dis)

	uu := util.GetDiskUsage()
	log.Println("uu : ", uu)

}
