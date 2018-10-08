// +build darwin

package disk

import (
	"bufio"
	"fwebpanel/types"
	"fwebpanel/utils"
	"strconv"
	"strings"
)

func GetAll() []types.DiskInfo {
	response, ok := utils.Exec("df", "-h")
	if !ok {
		return nil
	}

	diskList := make([]types.DiskInfo, 0, 0)

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

		fileSystem := parts[0]
		all := parts[1]
		used := parts[2]
		free := parts[3]
		usagePercentage, _ := strconv.Atoi(strings.TrimSuffix(parts[4], "%"))
		mounted := parts[5]

		diskStatus := types.DiskInfo{}
		diskStatus.FileSystem = fileSystem
		diskStatus.All = all
		diskStatus.Used = used
		diskStatus.Free = free
		diskStatus.UsagePercentage = uint(usagePercentage)
		diskStatus.Mounted = mounted

		diskList = append(diskList, diskStatus)
	}
	return diskList
}
