package fbopenmetrics

import (
    "purestorage.com/flashblade/client"
    "github.com/prometheus/client_golang/prometheus"
)

type HttpPerfCollector struct {
    LatencyDesc    *prometheus.Desc
    ThroughputDesc *prometheus.Desc
    Client         *restclient.FBClient
}

func (c *HttpPerfCollector) Describe(ch chan<- *prometheus.Desc) {
    prometheus.DescribeByCollect(c, ch)
}

func (c *HttpPerfCollector) Collect(ch chan<- prometheus.Metric) {
    arraysperf := c.Client.GetArraysHttpPerformance()
    if len(arraysperf.Items) == 0 {
        return
    }

    hp := arraysperf.Items[0]
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            hp.UsecPerOtherOp,
            "usec_per_other_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            hp.UsecPerReadDirOp,
            "usec_per_read_dir_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            hp.UsecPerReadFileOp,
            "usec_per_read_file_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            hp.UsecPerWriteDirOp,
            "usec_per_write_dir_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            hp.UsecPerWriteFileOp,
            "usec_per_write_file_op",
    )

    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            hp.OthersPerSec,
            "others_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            hp.ReadDirsPerSec,
            "read_dirs_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            hp.ReadFilesPerSec,
            "read_files_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            hp.WriteDirsPerSec,
            "write_dirs_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            hp.WriteFilesPerSec,
            "write_files_per_sec",
    )
}

func NewHttpPerfCollector(fb *restclient.FBClient) *HttpPerfCollector {
    return &HttpPerfCollector{
        LatencyDesc: prometheus.NewDesc(
            "purefb_array_http_specific_performance_latency_usec",
            "FlashBlade array HTTP specific latency",
            []string{"dimension"},
            prometheus.Labels{},
        ),
        ThroughputDesc: prometheus.NewDesc(
            "purefb_array_http_specific_performance_throughput_iops",
            "FlashBlade array HTTP specific throughput",
            []string{"dimension"},
            prometheus.Labels{},
        ),
        Client: fb,
    }
}
