package fbopenmetrics

import (
    "purestorage.com/flashblade/client"
    "github.com/prometheus/client_golang/prometheus"
)

type ClientsPerfCollector struct {
    LatencyDesc      *prometheus.Desc
    ThroughputDesc   *prometheus.Desc
    BandwidthDesc    *prometheus.Desc
    AverageSizeDesc  *prometheus.Desc
    Client           *restclient.FBClient
}

func (c *ClientsPerfCollector) Describe(ch chan<- *prometheus.Desc) {
    prometheus.DescribeByCollect(c, ch)
}

func (c *ClientsPerfCollector) Collect(ch chan<- prometheus.Metric) {
    clientsperf := c.Client.GetClientsPerformance()
    if len(clientsperf.Items) == 0 {
        return
    }

    for _, cp := range clientsperf.Items {
        ch <- prometheus.MustNewConstMetric(
                c.LatencyDesc,
                prometheus.GaugeValue,
                cp.UsecPerOtherOp,
                cp.Name, "usec_per_other_op",
        )
        ch <- prometheus.MustNewConstMetric(
                c.LatencyDesc,
                prometheus.GaugeValue,
                cp.UsecPerReadOp,
                cp.Name, "usec_per_read_op",
        )
        ch <- prometheus.MustNewConstMetric(
                c.LatencyDesc,
                prometheus.GaugeValue,
                cp.UsecPerWriteOp,
                cp.Name, "usec_per_write_op",
        )
        ch <- prometheus.MustNewConstMetric(
                c.ThroughputDesc,
                prometheus.GaugeValue,
                cp.OthersPerSec,
                cp.Name, "others_per_sec",
        )
        ch <- prometheus.MustNewConstMetric(
                c.ThroughputDesc,
                prometheus.GaugeValue,
                cp.ReadsPerSec,
                cp.Name, "reads_per_sec",
        )
        ch <- prometheus.MustNewConstMetric(
                c.ThroughputDesc,
                prometheus.GaugeValue,
                cp.WritesPerSec,
                cp.Name, "writes_per_sec",
        )
        ch <- prometheus.MustNewConstMetric(
                c.BandwidthDesc,
                prometheus.GaugeValue,
                cp.ReadBytesPerSec,
                cp.Name, "read_bytes_per_sec",
        )
        ch <- prometheus.MustNewConstMetric(
                c.BandwidthDesc,
                prometheus.GaugeValue,
                cp.WriteBytesPerSec,
                cp.Name, "write_bytes_per_sec",
        )
        ch <- prometheus.MustNewConstMetric(
                c.AverageSizeDesc,
                prometheus.GaugeValue,
                cp.BytesPerOp,
                cp.Name, "bytes_per_op",
        )
        ch <- prometheus.MustNewConstMetric(
                c.AverageSizeDesc,
                prometheus.GaugeValue,
                cp.BytesPerRead,
                cp.Name, "bytes_per_read",
        )
        ch <- prometheus.MustNewConstMetric(
                c.AverageSizeDesc,
                prometheus.GaugeValue,
                cp.BytesPerWrite,
                cp.Name, "bytes_per_write",
        )
    }
}

func NewClientsPerfCollector(fb *restclient.FBClient) *ClientsPerfCollector {
    return &ClientsPerfCollector{
        LatencyDesc: prometheus.NewDesc(
            "purefb_clients_performance_latency_usec",
            "FlashBlade clients latency",
            []string{"name", "dimension"},
            prometheus.Labels{},
        ),
        ThroughputDesc: prometheus.NewDesc(
            "purefb_clients_performance_throughput_iops",
            "FlashBlade clients throughput",
            []string{"name", "dimension"},
            prometheus.Labels{},
        ),
        BandwidthDesc: prometheus.NewDesc(
            "purefb_clients_performance_bandwidth_bytes",
            "FlashBlade clients bandwidth",
            []string{"name", "dimension"},
            prometheus.Labels{},
        ),
        AverageSizeDesc: prometheus.NewDesc(
            "purefb_clients_performance_average_bytes",
            "FlashBlade clients average operations size",
            []string{"name", "dimension"},
            prometheus.Labels{},
        ),
        Client: fb,
    }
}
