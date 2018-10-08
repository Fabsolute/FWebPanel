// +build darwin

package memory

import (
	"fmt"
	"fwebpanel/types"
	"math"
	"runtime"
)

func GetAll() []types.MemoryInfo {
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)

	memoryList := make([]types.MemoryInfo, 0, 2)

	memoryStatus := types.MemoryInfo{}
	memoryStatus.Name = "Mem"
	memoryStatus.Total = format(m.Sys)
	memoryStatus.Used = format(m.TotalAlloc)
	memoryStatus.Free = format(m.Frees)

	memoryList = append(memoryList, memoryStatus)

	return memoryList
}

func GetByName(name string) types.MemoryInfo {
	for _, memoryInfo := range GetAll() {
		if memoryInfo.Name == name {
			return memoryInfo
		}
	}
	return types.MemoryInfo{}
}

func format(s uint64) string {
	sizes := []string{"B", "K", "M", "G", "T", "P", "E"}
	return humanateBytes(s, 1000, sizes)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := math.Floor(float64(s)/math.Pow(base, e)*10+0.5) / 10
	f := "%.0f %s"
	if val < 10 {
		f = "%.1f %s"
	}

	return fmt.Sprintf(f, val, suffix)
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}
