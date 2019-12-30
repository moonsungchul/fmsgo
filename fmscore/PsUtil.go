package fmscore

import (
	"log"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"

	pb "github.com/moonsungchul/fmsgo/fmsgrpc"
	"github.com/shirou/gopsutil/mem"
)

type PsUtil struct {
}

func NewPsUtil() *PsUtil {
	return &PsUtil{}
}

func (s *PsUtil) GetMemory() *pb.MemoryInfo {
	v, _ := mem.VirtualMemory()
	log.Println(v)
	return &pb.MemoryInfo{Total: v.Total, Free: v.Free, UsePercent: v.UsedPercent}
}

func (s *PsUtil) GetCPU() (int, int) {
	logicCPU, _ := cpu.Counts(true)
	phyCPU, _ := cpu.Counts(false)
	return phyCPU, logicCPU
}

func (s *PsUtil) GetCPUPercent() []float64 {
	cpuPercent, _ := cpu.Percent(0, true)
	return cpuPercent
}

func (s *PsUtil) GetCPUInfo() *pb.CPUInfo {
	logicCPU, physicsCPU := s.GetCPU()
	cpuPercent := s.GetCPUPercent()
	cpuinfo := &pb.CPUInfo{LogicCPU: int32(logicCPU), PhysicsCPU: int32(physicsCPU), PercentCPU: cpuPercent}
	copy(cpuinfo.PercentCPU, cpuPercent)
	return cpuinfo
}

func (s *PsUtil) GetDiskInfo() []disk.PartitionStat {
	par, _ := disk.Partitions(true)
	log.Println("par : ", par)
	return par
}

func (s *PsUtil) GetDiskUsage(path string) (uint64, uint64, uint64, float64) {
	uu, _ := disk.Usage(path)
	log.Println(uu)
	return uu.Total, uu.Free, uu.Used, uu.UsedPercent
}
