package client

type Array struct {
	Name        string   `json:"name"`
	Id          string   `json:"id"`
	AsOf        int      `json:"_as_of"`
	Banner      string   `json:"banner"`
	IdleTimeout int      `json:"idle_timeout"`
	NtpServers  []string `json:"ntp_servers"`
	Os          string   `json:"os"`
	Revision    string   `json:"revision"`
	TimeZone    string   `json:"time_zone"`
	Version     string   `json:"version"`
}

type ArraysList struct {
	CntToken     string  `json:"continuation_token"`
	TotalItemCnt int     `json:"total_item_count"`
	Items        []Array `json:"items"`
}

func (fb *FBClient) GetArrays() *ArraysList {
	uri := "/arrays"
	result := new(ArraysList)
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
