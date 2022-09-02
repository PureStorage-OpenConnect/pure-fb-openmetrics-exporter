package collectors

import (
    "purestorage/fb-openmetrics-exporter/internal/rest-client"
    "github.com/prometheus/client_golang/prometheus"
)

type HwConnectorsPerfCollector struct {
    ThroughputDesc   *prometheus.Desc
    BandwidthDesc    *prometheus.Desc
    ErrorsDesc       *prometheus.Desc
    Client           *client.FBClient
}

func (c *HwConnectorsPerfCollector) Describe(ch chan<- *prometheus.Desc) {
    prometheus.DescribeByCollect(c, ch)
}

func (c *HwConnectorsPerfCollector) Collect(ch chan<- prometheus.Metric) {
    hwconn := c.Client.GetHwConnectorsPerformance()

    for _, conn := range hwconn.Items {
        ch <- prometheus.MustNewConstMetric(
                c.ThroughputDesc,
                prometheus.GaugeValue,
                conn.ReceivedPacketsPerSec,
                conn.Name, "received_packets_per_sec",
        )
        ch <- prometheus.MustNewConstMetric(
                c.ThroughputDesc,
                prometheus.GaugeValue,
                conn.TransmittedPacketsPerSec,
                conn.Name, "transmitted_packets_per_sec",
        )
        ch <- prometheus.MustNewConstMetric(
                c.BandwidthDesc,
                prometheus.GaugeValue,
                conn.ReceivedBytesPerSec,
                conn.Name, "received_bytes_per_sec",
        )
        ch <- prometheus.MustNewConstMetric(
                c.BandwidthDesc,
                prometheus.GaugeValue,
                conn.TransmittedBytesPerSec,
                conn.Name, "transmitted_bytes_per_sec",
        )
        ch <- prometheus.MustNewConstMetric(
                c.ErrorsDesc,
                prometheus.GaugeValue,
                conn.OtherErrorsPerSec,
                conn.Name, "other_errors_per_sec",
        )
        ch <- prometheus.MustNewConstMetric(
                c.ErrorsDesc,
                prometheus.GaugeValue,
                conn.ReceivedCRCErrorsPerSec,
                conn.Name, "received_crc_errors_per_sec",
        )
        ch <- prometheus.MustNewConstMetric(
                c.ErrorsDesc,
                prometheus.GaugeValue,
                conn.ReceivedFrameErrorsPerSec,
                conn.Name, "received_frame_errors_per_sec",
        )
        ch <- prometheus.MustNewConstMetric(
                c.ErrorsDesc,
                prometheus.GaugeValue,
                conn.TransmittedCarrierErrorsPerSec,
                conn.Name, "transmitted_carrier_errors_per_sec",
        )
        ch <- prometheus.MustNewConstMetric(
                c.ErrorsDesc,
                prometheus.GaugeValue,
                conn.TransmittedDroppedErrorsPerSec,
                conn.Name, "transmitted_dropped_errors_per_sec",
        )
        ch <- prometheus.MustNewConstMetric(
                c.ErrorsDesc,
                prometheus.GaugeValue,
                conn.TotalErrorsPerSec,
                conn.Name, "total_errors_per_sec",
        )
    }
}

func NewHwConnectorsPerfCollector(fb *client.FBClient) *HwConnectorsPerfCollector {
    return &HwConnectorsPerfCollector{
        ThroughputDesc: prometheus.NewDesc(
            "purefb_hardware_connectors_performance_throughput_pkts",
            "FlashBlade hardware connectors performance throughput",
            []string{"name", "dimension"},
            prometheus.Labels{},
        ),
        BandwidthDesc: prometheus.NewDesc(
            "purefb_hardware_connectors_performance_bandwidth_bytes",
            "FlashBlade hardware connectors performance bandwidth",
            []string{"name", "dimension"},
            prometheus.Labels{},
        ),
        ErrorsDesc: prometheus.NewDesc(
            "purefb_shardware_connectors_performance_errors",
            "FlashBlade hardware connectors performance errors per sec",
            []string{"name", "dimension"},
            prometheus.Labels{},
        ),
        Client: fb,
    }
}
