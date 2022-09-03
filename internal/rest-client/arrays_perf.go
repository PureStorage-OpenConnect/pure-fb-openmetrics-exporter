package client

type ArraysPerformanceList struct {
	CntToken     string        `json:"continuation_token"`
	TotalItemCnt int           `json:"total_item_count"`
	Items        []Performance `json:"items"`
}

func (fb *FBClient) GetArraysPerformance(protocol string) *ArraysPerformanceList {
	result := new(ArraysPerformanceList)
	switch protocol {
	case "all", "HTTP", "NFS", "SMB", "S3":
		_, err := fb.RestClient.R().
			SetResult(&result).
			SetQueryParam("protocol", protocol).
			Get("/arrays/performance")

		if err != nil {
			fb.Error = err
		}
	}
	return result
}
