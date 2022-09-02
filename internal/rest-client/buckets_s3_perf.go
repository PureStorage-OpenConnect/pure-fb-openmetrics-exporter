package client


type BucketsS3PerformanceList struct {
    CntToken       string            `json:"continuation_token"`
    TotalItemCnt   int               `json:"total_item_count"`
    Items          []S3Performance   `json:"items"`
    Total          []S3Performance   `json:"total"`
}

func (fb *FBClient) GetBucketsS3Performance(b *BucketsList) *BucketsS3PerformanceList {
    result := new(BucketsS3PerformanceList)
    if b == nil {
        return result
    }
    temp := new(BucketsS3PerformanceList)
    for i := 0; i < len(b.Items); i += 5 {
        n := ""
        for j := 0; (j < 5) && (i + j < len(b.Items)); j++ {
           n = n + b.Items[i+j].Name + "," 
        }
        n = n[:len(n)-1]
        _, err := fb.RestClient.R().
                        SetResult(&temp).
                        SetQueryParam("names", n).
                        Get("/buckets/s3-specific-performance")
    
        if err != nil {
            fb.Error = err
        }
        result.Items = append(result.Items, temp.Items...)
    }
    return result
}
