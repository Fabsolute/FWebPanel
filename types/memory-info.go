package types

type MemoryInfo struct {
	Name      string `json:"name,omitempty"`
	Total     string `json:"total,omitempty"`
	Used      string `json:"used,omitempty"`
	Free      string `json:"free,omitempty"`
	Shared    string `json:"shared,omitempty"`
	Buffers   string `json:"buffers,omitempty"`
	Cache     string `json:"cache,omitempty"`
	Available string `json:"available,omitempty"`
}
