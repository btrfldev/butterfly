package butterfly

type Query struct {
	Objects []Object `json:"objects"`
}

type Object struct {
	Lib   string `json:"lib"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ListResp struct {
	Lists []List `json:"keysLists"`
}

type List struct {
	Prefix string   `json:"prefix"`
	Count  int      `json:"count"`
	Keys   []string `json:"libs:keys"`
}

type Health struct {
	Status          string `json:"status"`
	UTC             string `json:"utc"`
	NodeType        string `json:"type"`
	Version         string `json:"version"`
	TotalMemory     int    `json:"totalMemory"`
	AvailableMemory int    `json:"availableMemory"`
}
