package client

type RemoteArray struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PerformanceReplication struct {
	TransmittedBytesPerSec float64 `json:"transmitted_bytes_per_sec"`
	ReceivedBytesPerSec    float64 `json:"received_bytes_per_sec"`
}

type ArrayPerformanceReplication struct {
	Id        string                 `json:"id"`
	Periodic  PerformanceReplication `json:"periodic"`
	Remote    RemoteArray            `json:"remote"`
	Aggreate  PerformanceReplication `json:"aggregate"`
	Continuos PerformanceReplication `json:"continuos"`
	Time      int                    `json:"time"`
}

type ArraysPerformanceReplicationList struct {
	CntToken     string                        `json:"continuation_token"`
	TotalItemCnt int                           `json:"total_item_count"`
	Items        []ArrayPerformanceReplication `json:"items"`
}

func (fb *FBClient) GetArraysPerformanceReplication() *ArraysPerformanceReplicationList {
	uri := "/arrays/performance/replication"
	result := new(ArraysPerformanceReplicationList)
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
