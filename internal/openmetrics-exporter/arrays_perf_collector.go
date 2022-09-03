package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fb-openmetrics-exporter/internal/rest-client"
)

type PerfCollector struct {
	LatencyDesc     *prometheus.Desc
	ThroughputDesc  *prometheus.Desc
	BandwidthDesc   *prometheus.Desc
	AverageSizeDesc *prometheus.Desc
	Client          *client.FBClient
}

func (c *PerfCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *PerfCollector) Collect(ch chan<- prometheus.Metric) {
	protocols := []string{"all", "HTTP", "SMB", "NFS", "S3"}
	for _, proto := range protocols {
		arraysperf := c.Client.GetArraysPerformance(proto)
		if len(arraysperf.Items) == 0 {
			continue
		}

		ap := arraysperf.Items[0]
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			ap.UsecPerOtherOp,
			proto, "usec_per_other_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			ap.UsecPerReadOp,
			proto, "usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			ap.UsecPerWriteOp,
			proto, "usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			ap.OthersPerSec,
			proto, "others_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			ap.ReadsPerSec,
			proto, "reads_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			ap.WritesPerSec,
			proto, "writes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			ap.ReadBytesPerSec,
			proto, "read_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			ap.WriteBytesPerSec,
			proto, "write_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			ap.BytesPerOp,
			proto, "bytes_per_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			ap.BytesPerRead,
			proto, "bytes_per_read",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			ap.BytesPerWrite,
			proto, "bytes_per_write",
		)
	}
}

func NewPerfCollector(fb *client.FBClient) *PerfCollector {
	return &PerfCollector{
		LatencyDesc: prometheus.NewDesc(
			"purefb_array_performance_latency_usec",
			"FlashBlade array latency",
			[]string{"protocol", "dimension"},
			prometheus.Labels{},
		),
		ThroughputDesc: prometheus.NewDesc(
			"purefb_array_performance_throughput_iops",
			"FlashBlade array throughput",
			[]string{"protocol", "dimension"},
			prometheus.Labels{},
		),
		BandwidthDesc: prometheus.NewDesc(
			"purefb_array_performance_bandwidth_bytes",
			"FlashBlade array throughput",
			[]string{"protocol", "dimension"},
			prometheus.Labels{},
		),
		AverageSizeDesc: prometheus.NewDesc(
			"purefb_array_performance_average_bytes",
			"FlashBlade array average operations size",
			[]string{"protocol", "dimension"},
			prometheus.Labels{},
		),
		Client: fb,
	}
}
