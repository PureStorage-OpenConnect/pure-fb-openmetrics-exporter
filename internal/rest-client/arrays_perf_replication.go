package restclient

type RemoteArray struct {
    Id       string  `json:"id"`
    Name     string  `json:"name"`
}

type PerformanceReplication struct {
    TransmittedBytesPerSec   float64  `json:"transmitted_bytes_per_sec"`
    ReceivedBytesPerSec      float64  `json:"received_bytes_per_sec"`
}

type ArrayPerformanceReplication struct {
    Id           string                   `json:"id"`
    Periodic     PerformanceReplication   `json:"periodic"`
    Remote       RemoteArray              `json:"remote"`
    Aggreate     PerformanceReplication   `json:"aggregate"`
    Continuos    PerformanceReplication   `json:"continuos"`
    Time         int                      `json:"time"`
}

type ArraysPerformanceReplicationList struct {
    CntToken       string            `json:"continuation_token"`
    TotalItemCnt   int               `json:"total_item_count"`
    Items          []ArrayPerformanceReplication     `json:"items"`
}

func (fb *FBClient) GetArraysPerformanceReplication() *ArraysPerformanceReplicationList {
    result := new(ArraysPerformanceReplicationList)
    _, err := fb.RestClient.R().
                    SetResult(&result).
                    Get("/arrays/performance/replication")
    
    if err != nil {
        fb.Error = err
    }
    return result
}
