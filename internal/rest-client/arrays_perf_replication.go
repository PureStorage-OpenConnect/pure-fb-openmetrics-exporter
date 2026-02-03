package client

type ResourceRef struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ResourceType string `json:"resource_type"`
}

type PerformanceReplication struct {
	TransmittedBytesPerSec float64 `json:"transmitted_bytes_per_sec"`
	ReceivedBytesPerSec    float64 `json:"received_bytes_per_sec"`
}

type ObjectBacklog struct {
	BytesCount     float64 `json:"bytes_count"`
	DeleteOpsCount float64 `json:"delete_ops_count"`
	OtherOpsCount  float64 `json:"other_ops_count"`
	PutOpsCount    float64 `json:"put_ops_count"`
}

type ContinuousPerformanceReplication struct {
	TransmittedBytesPerSec float64        `json:"transmitted_bytes_per_sec"`
	ReceivedBytesPerSec    float64        `json:"received_bytes_per_sec"`
	ObjectBacklog          *ObjectBacklog `json:"object_backlog,omitempty"`
}

type ArrayPerformanceReplication struct {
	Id         string                            `json:"id"`
	Name       string                            `json:"name,omitempty"`
	Periodic   PerformanceReplication            `json:"periodic"`
	Context    *ResourceRef                      `json:"context,omitempty"`
	Remote     *ResourceRef                      `json:"remote,omitempty"` // legacy schemas may use 'remote'
	Aggreate   PerformanceReplication            `json:"aggregate"`
	Continuous *ContinuousPerformanceReplication `json:"continuous,omitempty"`
	Continuos  *ContinuousPerformanceReplication `json:"continuos,omitempty"` // legacy misspelling
	Time       int64                             `json:"time"`
}

func (a *ArrayPerformanceReplication) GetContinuous() *ContinuousPerformanceReplication {
	if a == nil {
		return nil
	}
	if a.Continuous != nil {
		return a.Continuous
	}
	return a.Continuos
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
