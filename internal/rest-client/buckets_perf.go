package client

type BucketsPerformanceList struct {
	CntToken     string        `json:"continuation_token"`
	TotalItemCnt int           `json:"total_item_count"`
	Items        []Performance `json:"items"`
	Total        []Performance `json:"total"`
}

func (fb *FBClient) GetBucketsPerformance(b *BucketsList) *BucketsPerformanceList {
	result := new(BucketsPerformanceList)
	if b == nil {
		return result
	}
	temp := new(BucketsPerformanceList)
	for i := 0; i < len(b.Items); i += 5 {
		n := ""
		for j := 0; (j < 5) && (i+j < len(b.Items)); j++ {
			n = n + b.Items[i+j].Name + ","
		}
		n = n[:len(n)-1]
		res, _ := fb.RestClient.R().
			SetResult(&temp).
			SetQueryParam("names", n).
			Get("/buckets/performance")
		if res.StatusCode() == 401 {
			fb.RefreshSession()
			fb.RestClient.R().
				SetResult(&temp).
				SetQueryParam("names", n).
				Get("/buckets/performance")
		}
		result.Items = append(result.Items, temp.Items...)
	}
	return result
}
