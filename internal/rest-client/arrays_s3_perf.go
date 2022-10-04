package client

type ArraysS3PerformanceList struct {
	CntToken     string          `json:"continuation_token"`
	TotalItemCnt int             `json:"total_item_count"`
	Items        []S3Performance `json:"items"`
}

func (fb *FBClient) GetArraysS3Performance() *ArraysS3PerformanceList {
	result := new(ArraysS3PerformanceList)
	res, _ := fb.RestClient.R().
		SetResult(&result).
		Get("/arrays/s3-specific-performance")
	if res.StatusCode() == 401 {
		fb.RefreshSession()
		fb.RestClient.R().
			SetResult(&result).
			Get("/arrays/s3-specific-performance")
        }
	return result
}
