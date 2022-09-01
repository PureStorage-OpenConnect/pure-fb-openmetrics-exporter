package restclient


type ArraysS3PerformanceList struct {
    CntToken       string           `json:"continuation_token"`
    TotalItemCnt   int              `json:"total_item_count"`
    Items          []S3Performance  `json:"items"`
}

func (fb *FBClient) GetArraysS3Performance() *ArraysS3PerformanceList {
    result := new(ArraysS3PerformanceList)
    _, err := fb.RestClient.R().
                    SetResult(&result).
                    Get("/arrays/s3-specific-performance")
    
    if err != nil {
        fb.Error = err
    }
    return result
}
