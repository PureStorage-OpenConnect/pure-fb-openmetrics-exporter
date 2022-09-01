package restclient

type Blade struct {
    Name          string   `json:"name"`
    Id            string   `json:"id"`
    Details       string   `json:"details"`
    Progress      int      `json:"progress"`
    RawCapacity   int      `json:"raw_capacity"`
    Status        string   `json:"status"`
    Target        string   `json:"target"`
}

type BladesList struct {
    CntToken       string    `json:"continuation_token"`
    TotalItemCnt   int       `json:"total_item_count"`
    Items          []Blade   `json:"items"`
    Total          Blade     `json:"total"`
}

func (fb *FBClient) GetBlades() *BladesList {
    result := new(BladesList)
    _, err := fb.RestClient.R().
                    SetResult(&result).
                    Get("/blades")
    
    if err != nil {
        fb.Error = err
    }
    return result
}
