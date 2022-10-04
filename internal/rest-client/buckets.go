package client

type Account struct {
	Name         string `json:"name"`
	Id           string `json:"id"`
	ResourceType string `json:"resource_type"`
}

type Bucket struct {
	Name        string  `json:"name"`
	Id          string  `json:"id"`
	Account     Account `json:"account"`
	Created     int     `json:"created"`
	destroyed   bool    `json:"destroyed"`
	ObjectCount int     `json:"object_count"`
	Space       Space   `json:"space"`
}

type BucketsList struct {
	CntToken     string   `json:"continuation_token"`
	TotalItemCnt int      `json:"total_item_count"`
	Items        []Bucket `json:"items"`
	Total        Bucket   `json:"total"`
}

func (fb *FBClient) GetBuckets() *BucketsList {
	result := new(BucketsList)
	res, _ := fb.RestClient.R().
		SetResult(&result).
		Get("/buckets")
	if res.StatusCode() == 401 {
                fb.RefreshSession()
		fb.RestClient.R().
			SetResult(&result).
			Get("/buckets")
        }	
	return result
}
