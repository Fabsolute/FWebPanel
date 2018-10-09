package cpu

import (
	"bytes"
	"fwebpanel/types"
	"fwebpanel/utils/cmd"
	pipe2 "fwebpanel/utils/cmd/pipe"
	"regexp"
	"strconv"
	"strings"
)

func GetUsage() types.CPUUsage {
	var out bytes.Buffer
	pipe2.New(&out).
		Pipe("top", "-F", "-l1", "-s3").
		Pipe("grep", "CPU usage").
		Run()

	r := regexp.MustCompile("CPU usage: ([0-9.]+%) user, ([0-9.]+%) sys, ([0-9.]+%) idle")
	split := r.FindStringSubmatch(out.String())
	user := split[1]
	sys := split[2]
	idle := split[3]
	return types.CPUUsage{Idle: idle, Sys: sys, User: user}
}

func GetInfo() types.CPUInfo {
	usage := GetUsage()

	coreCountString, _ := cmd.Exec("sysctl", "-n", "machdep.cpu.core_count")
	threadCountString, _ := cmd.Exec("sysctl", "-n", "machdep.cpu.thread_count")
	brand, _ := cmd.Exec("sysctl", "-n", "machdep.cpu.brand_string")

	coreCount,_ := strconv.Atoi(strings.TrimSpace(coreCountString))
	threadCount,_ := strconv.Atoi(strings.TrimSpace(threadCountString))

	brand = strings.TrimSpace(brand)

	return types.CPUInfo{CoreCount: coreCount, ThreadCount: threadCount, Brand: brand, Usage: usage}
}
