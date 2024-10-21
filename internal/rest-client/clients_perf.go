package client

type ClientPerformance struct {
	Name             string  `json:"name"`
	BytesPerOp       float64 `json:"bytes_per_op"`
	BytesPerRead     float64 `json:"bytes_per_read"`
	BytesPerWrite    float64 `json:"bytes_per_write"`
	OthersPerSec     float64 `json:"others_per_sec"`
	ReadBytesPerSec  float64 `json:"read_bytes_per_sec"`
	ReadsPerSec      float64 `json:"reads_per_sec"`
	Time             int     `json:"time"`
	UsecPerOtherOp   float64 `json:"usec_per_other_op"`
	UsecPerReadOp    float64 `json:"usec_per_read_op"`
	UsecPerWriteOp   float64 `json:"usec_per_write_op"`
	WriteBytesPerSec float64 `json:"write_bytes_per_sec"`
	WritesPerSec     float64 `json:"writes_per_sec"`
}

type ClientsPerformanceList struct {
	CntToken      string              `json:"continuation_token"`
	TotalItemsCnt int                 `json:"total_item_count"`
	Items         []ClientPerformance `json:"items"`
	Total         []ClientPerformance `json:"total"`
}

func (fb *FBClient) GetClientsPerformance() *ClientsPerformanceList {
	uri := "/arrays/clients/performance"
	result := new(ClientsPerformanceList)
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
