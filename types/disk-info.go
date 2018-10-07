package types

type DiskInfo struct {
	FileSystem      string `json:"file_system"`
	All             string `json:"all"`
	Used            string `json:"used"`
	Free            string `json:"free"`
	UsagePercentage uint   `json:"usage_percentage"`
	Mounted         string `json:"mounted"`
}
