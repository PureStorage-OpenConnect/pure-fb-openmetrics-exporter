package client

type Blade struct {
	Name        string `json:"name"`
	Id          string `json:"id"`
	Details     string `json:"details"`
	Progress    int    `json:"progress"`
	RawCapacity int    `json:"raw_capacity"`
	Status      string `json:"status"`
	Target      string `json:"target"`
}

type BladesList struct {
	CntToken     string  `json:"continuation_token"`
	TotalItemCnt int     `json:"total_item_count"`
	Items        []Blade `json:"items"`
	Total        Blade   `json:"total"`
}

func (fb *FBClient) GetBlades() *BladesList {
	uri := "/blades"
	result := new(BladesList)
	res, _ := fb.RestClient.R().
		SetResult(&result).
		Get(uri)

	if res.StatusCode() == 401 {
		fb.RefreshSession()
		fb.RestClient.R().
			SetResult(&result).
			Get(uri)
	}
	return result
}
