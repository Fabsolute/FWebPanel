// +build darwin

package memory

import (
	"bufio"
	"bytes"
	"fmt"
	"fwebpanel/types"
	"fwebpanel/utils/cmd"
	"math"
	"strconv"
	"strings"
)

func GetAll() []types.MemoryInfo {
	memoryList := make([]types.MemoryInfo, 0)
	{
		response, ok := cmd.Exec("sysctl", "-n", "hw.memsize")
		if !ok {
			return nil
		}

		reader := bufio.NewReader(strings.NewReader(response))
		line, _, err := reader.ReadLine()
		if err != nil {
			return nil
		}

		totalMemory, _ := strconv.ParseUint(string(line), 10, 64)
		memoryInfo := types.MemoryInfo{}
		memoryInfo.Name = "Mem"
		memoryInfo.Total = format(totalMemory)
		totalPaged := getPaged("wired down", "active", "inactive")
		memoryInfo.Used = format(totalPaged * 4096)
		memoryInfo.Free = format(totalMemory - totalPaged*4096)

		memoryList = append(memoryList, memoryInfo)
	}
	{
		response, ok := cmd.Exec("sysctl", "vm.swapusage")
		if !ok {
			return nil
		}

		reader := bufio.NewReader(strings.NewReader(response))
		line, _, err := reader.ReadLine()
		if err != nil {
			return nil
		}

		totalRight := strings.Split(string(line), "total =")[1]
		usedSplit := strings.Split(totalRight, "used =")
		usedRight := usedSplit[1]
		freeSplit := strings.Split(usedRight, "free =")
		freeRight := freeSplit[1]
		encryptedSplit := strings.Split(freeRight, "(encrypted")

		total := strings.TrimSpace(usedSplit[0])
		used := strings.TrimSpace(freeSplit[0])
		free := strings.TrimSpace(encryptedSplit[0])

		memoryInfo := types.MemoryInfo{}
		memoryInfo.Name = "Swap"
		memoryInfo.Total = total
		memoryInfo.Used = used
		memoryInfo.Free = free

		memoryList = append(memoryList, memoryInfo)
	}
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

func getPaged(match ...string) uint64 {
	if len(match) == 1 {
		count := "3"
		if len(strings.Fields(match[0])) > 1 {
			count = "4"
		}

		var out bytes.Buffer
		pipes := cmd.NewPipe(&out)
		pipes.Pipe("vm_stat").
			Pipe("grep", "Pages "+match[0]).
			Pipe("awk", "{print $"+count+"}").
			Run()

		response := strings.TrimSpace(out.String())
		response = strings.TrimSuffix(response, ".")
		memory, _ := strconv.ParseUint(response, 10, 64)
		return memory
	}

	var total uint64
	for _, page := range match {
		total += getPaged(page)
	}

	return total
}

func format(s uint64) string {
	sizes := []string{"B", "K", "M", "G", "T", "P", "E"}
	return makeHumanReadable(s, 1000, sizes)
}

func makeHumanReadable(s uint64, base float64, sizes []string) string {
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
