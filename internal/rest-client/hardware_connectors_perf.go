package client


type HwConnectorPerformance struct {
    Name                            string `json:"name"`
    Id                              string `json:"id"`
    LAG                             LinkAggregationGroup   `json:"link_aggregation_group"`
    OtherErrorsPerSec               float64 `json:"other_errors_per_sec"`
    ReceivedBytesPerSec             float64 `json:"received_bytes_per_sec"`
    ReceivedCRCErrorsPerSec         float64 `json:"received_crc_errors_per_sec"`
    ReceivedFrameErrorsPerSec       float64 `json:"received_frame_errors_per_sec"`
    ReceivedPacketsPerSec           float64 `json:"received_packets_per_sec"`
    Time                            float64 `json:"time"`
    TotalErrorsPerSec               float64 `json:"total_errors_per_sec"`
    TransmittedBytesPerSec          float64 `json:"transmitted_bytes_per_sec"`
    TransmittedCarrierErrorsPerSec  float64 `json:"transmitted_carrier_errors_per_sec"`
    TransmittedDroppedErrorsPerSec  float64 `json:"transmitted_dropped_errors_per_sec"`
    TransmittedPacketsPerSec        float64 `json:"transmitted_packets_per_sec"`
}

type LinkAggregationGroup struct {
    Name             string   `json:"name"`
    Id               string   `json:"id"`
    ResourceType     string   `json:"resource_type"`
}

type HwConnectorsPerformanceList struct {
    CntToken       string                    `json:"continuation_token"`
    TotalItemCnt   int                       `json:"total_item_count"`
    Items          []HwConnectorPerformance   `json:"items"`
    Total          []HwConnectorPerformance   `json:"total"`
}

func (fb *FBClient) GetHwConnectorsPerformance() *HwConnectorsPerformanceList {
    result := new(HwConnectorsPerformanceList)
    _, err := fb.RestClient.R().
                    SetResult(&result).
                    Get("/hardware-connectors/performance")

    if err != nil {
        fb.Error = err
    }
    return result
}
