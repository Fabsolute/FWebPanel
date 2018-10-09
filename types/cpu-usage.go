package types

type CPUUsage struct {
	Idle string `json:"idle"`
	User string `json:"user"`
	Sys  string `json:"sys"`
}
