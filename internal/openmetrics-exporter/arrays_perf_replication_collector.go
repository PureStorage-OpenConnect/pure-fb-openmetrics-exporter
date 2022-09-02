package collectors

import (
    "purestorage/fb-openmetrics-exporter/internal/rest-client"
    "github.com/prometheus/client_golang/prometheus"
)

type PerfReplicationCollector struct {
    ThroughputDesc *prometheus.Desc
    Client         *client.FBClient
}

func (c *PerfReplicationCollector) Describe(ch chan<- *prometheus.Desc) {
    prometheus.DescribeByCollect(c, ch)
}

func (c *PerfReplicationCollector) Collect(ch chan<- prometheus.Metric) {
    arraysreplperf := c.Client.GetArraysPerformanceReplication()
    if len(arraysreplperf.Items) == 0 {
        return
    }

    arp := arraysreplperf.Items[0]
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            arp.Continuos.TransmittedBytesPerSec,
            "transmitted_bytes_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            arp.Continuos.ReceivedBytesPerSec,
            "received_bytes_per_sec",
    )
}

func NewPerfReplicationCollector(fb *client.FBClient) *PerfReplicationCollector {
    return &PerfReplicationCollector{
        ThroughputDesc: prometheus.NewDesc(
            "purefb_array_performance_replication",
            "FlashBlade array replication throughput",
            []string{"dimension"},
            prometheus.Labels{},
        ),
        Client: fb,
    }
}
