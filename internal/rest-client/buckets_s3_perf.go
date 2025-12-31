package client

import (
	"strings"
)

type BucketsS3PerformanceList struct {
	CntToken     string          `json:"continuation_token"`
	TotalItemCnt int             `json:"total_item_count"`
	Items        []S3Performance `json:"items"`
	Total        []S3Performance `json:"total"`
}

func (fb *FBClient) GetBucketsS3Performance(b *BucketsList) *BucketsS3PerformanceList {
	uri := "/buckets/s3-specific-performance"
	result := new(BucketsS3PerformanceList)
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
		temp := new(BucketsS3PerformanceList)
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
