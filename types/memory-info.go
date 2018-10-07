package types

type MemoryInfo struct {
	Name      string `json:"name"`
	Total     string `json:"total"`
	Used      string `json:"used"`
	Free      string `json:"free"`
	Shared    string `json:"shared"`
	Buffers   string `json:"buffers"`
	Cache     string `json:"cache"`
	Available string `json:"available"`
}
