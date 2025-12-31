package client

import (
	"strings"
)

type BucketsPerformanceList struct {
	CntToken     string        `json:"continuation_token"`
	TotalItemCnt int           `json:"total_item_count"`
	Items        []Performance `json:"items"`
	Total        []Performance `json:"total"`
}

func (fb *FBClient) GetBucketsPerformance(b *BucketsList) *BucketsPerformanceList {
	uri := "/buckets/performance"
	result := new(BucketsPerformanceList)
	if b == nil {
		return result
	}
	const chunkSize = 10

	for i := 0; i < len(b.Items); i += chunkSize {
		names := make([]string, 0, chunkSize)
		for _, bucket := range b.Items[i:min(i+chunkSize, len(b.Items))] {
			names = append(names, bucket.Name)
		}
		n := strings.Join(names, ",")
		temp := new(BucketsPerformanceList)
		res, _ := fb.RestClient.R().
			SetResult(&temp).
			SetQueryParam("names", n).
			Get(uri)
		if res.StatusCode() == 401 {
			fb.RefreshSession()
			fb.RestClient.R().
				SetResult(&temp).
				SetQueryParam("names", n).
				Get(uri)
		}
		result.Items = append(result.Items, temp.Items...)
	}
	return result
}
