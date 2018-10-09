// +build darwin

package memory

import (
	"bufio"
	"bytes"
	"fwebpanel/types"
	"fwebpanel/utils/cmd"
	"fwebpanel/utils/cmd/pipe"
	"fwebpanel/utils/math"
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
		memoryInfo.Total = math.Format(totalMemory)
		totalPaged := getPaged("wired down", "active", "inactive")
		memoryInfo.Used = math.Format(totalPaged * 4096)
		memoryInfo.Free = math.Format(totalMemory - totalPaged*4096)

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

		total := math.Deformat(usedSplit[0])
		used := math.Deformat(freeSplit[0])
		free := math.Deformat(encryptedSplit[0])

		memoryInfo := types.MemoryInfo{}
		memoryInfo.Name = "Swap"
		memoryInfo.Total = math.Format(total)
		memoryInfo.Used = math.Format(used)
		memoryInfo.Free = math.Format(free)

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
		pipes :=  pipe.New(&out)
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
