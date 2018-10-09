//+build linux

package memory

import (
	"bufio"
	"fwebpanel/types"
	"fwebpanel/utils/cmd"
	"strings"
)

func GetAll() []types.MemoryInfo {
	response, ok := cmd.Exec("free", "-wh")
	if !ok {
		return nil
	}

	memoryList := make([]types.MemoryInfo, 0, 2)

	reader := bufio.NewReader(strings.NewReader(response))
	var (
		err error = nil
	)

	isFirstLine := true
	for err == nil {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		if isFirstLine {
			isFirstLine = false
			continue
		}

		parts := strings.Fields(string(line))

		memoryStatus := types.MemoryInfo{}
		memoryStatus.Name = strings.TrimSuffix(parts[0], ":")
		memoryStatus.Total = parts[1]
		memoryStatus.Used = parts[2]
		memoryStatus.Free = parts[3]
		if len(parts) > 4 {
			memoryStatus.Shared = parts[4]
			memoryStatus.Buffers = parts[5]
			memoryStatus.Cache = parts[6]
			memoryStatus.Available = parts[7]
		}
		memoryList = append(memoryList, memoryStatus)
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
