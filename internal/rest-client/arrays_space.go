package client

type ArraySpace struct {
	Name     string  `json:"name"`
	Id       string  `json:"id"`
	Capacity int     `json:"capacity"`
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
		_, err := fb.RestClient.R().
			SetResult(&result).
			SetQueryParam("type", t).
			Get("/arrays/space")

		if err != nil {
			fb.Error = err
		}
	}
	return result
}
