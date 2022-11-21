package client

type ArraySpace struct {
	Name     string  `json:"name"`
	Id       string  `json:"id"`
	Capacity float64 `json:"capacity"`
	Parity   float64 `json:"parity"`
	Space    Space   `json:"space"`
	Time     int     `json:"time"`
}

type ArraysSpaceList struct {
	CntToken     string       `json:"continuation_token"`
	TotalItemCnt int          `json:"total_item_count"`
	Items        []ArraySpace `json:"items"`
}

func (fb *FBClient) GetArraysSpace(t string) *ArraysSpaceList {
	result := new(ArraysSpaceList)
	switch t {
	case "array", "file-system", "object-store":
		res, _ := fb.RestClient.R().
			SetResult(&result).
			SetQueryParam("type", t).
			Get("/arrays/space")
		if res.StatusCode() == 401 {
			fb.RefreshSession()
			fb.RestClient.R().
				SetResult(&result).
				SetQueryParam("type", t).
				Get("/arrays/space")
		}
	}
	return result
}
