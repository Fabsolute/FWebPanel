package memory

import (
	"bufio"
	"fwebpanel/types"
	"fwebpanel/utils"
	"strings"
)

func GetAll() []types.MemoryInfo {
	response, ok := utils.Exec("free", "-wth")
	if !ok {
		return nil
	}

	memoryList := make([]types.MemoryInfo, 0, 0)

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

		diskStatus := types.MemoryInfo{}
		diskStatus.Name = strings.TrimSuffix(parts[0], ":")
		diskStatus.Total = parts[1]
		diskStatus.Used = parts[2]
		diskStatus.Free = parts[3]
		if len(parts) > 4 {
			diskStatus.Shared = parts[4]
			diskStatus.Buffers = parts[5]
			diskStatus.Cache = parts[6]
			diskStatus.Available = parts[7]
		}
		memoryList = append(memoryList, diskStatus)
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
