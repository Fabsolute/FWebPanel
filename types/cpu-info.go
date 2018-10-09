package types

type CPUInfo struct {
	CoreCount   int   `json:"core_count"`
	ThreadCount int   `json:"thread_count"`
	Brand       string   `json:"brand"`
	Usage       CPUUsage `json:"usage"`
}
