package client

type HardwareComponent struct {
	Name            string  `json:"name"`
	Id              string  `json:"id"`
	Details         string  `json:"details"`
	IdentifyEnabled bool    `json:"identify_enabled"`
	Index           int     `json:"index"`
	Model           string  `json:"model"`
	Serial          string  `json:"serial"`
	Slot            int     `json:"slot"`
	Speed           float64 `json:"speed"`
	Status          string  `json:"status"`
	Temperature     float64 `json:"temperature"`
	Type            string  `json:"type"`
}

type HardwareList struct {
	CntToken     string              `json:"continuation_token"`
	TotalItemCnt int                 `json:"total_item_count"`
	Items        []HardwareComponent `json:"items"`
}

func (fb *FBClient) GetHardware() *HardwareList {
	result := new(HardwareList)
	_, err := fb.RestClient.R().
		SetResult(&result).
		Get("/hardware")

	if err != nil {
		fb.Error = err
	}
	return result
}
