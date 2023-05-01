package client

type ArraysPerformanceList struct {
	CntToken     string        `json:"continuation_token"`
	TotalItemCnt int           `json:"total_item_count"`
	Items        []Performance `json:"items"`
}

func (fb *FBClient) GetArraysPerformance(protocol string) *ArraysPerformanceList {
	uri := "/arrays/performance"
	result := new(ArraysPerformanceList)
	switch protocol {
	case "all", "HTTP", "NFS", "SMB", "S3":
		res, _ := fb.RestClient.R().
			SetResult(&result).
			SetQueryParam("protocol", protocol).
			Get(uri)
		if res.StatusCode() == 401 {
                	fb.RefreshSession()
			fb.RestClient.R().
			        SetResult(&result).
			        SetQueryParam("protocol", protocol).
			        Get(uri)
		}
	}
	return result
}
