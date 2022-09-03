package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fb-openmetrics-exporter/internal/rest-client"
)

type BucketsPerfCollector struct {
	LatencyDesc     *prometheus.Desc
	ThroughputDesc  *prometheus.Desc
	BandwidthDesc   *prometheus.Desc
	AverageSizeDesc *prometheus.Desc
	Client          *client.FBClient
	Buckets         *client.BucketsList
}

func (c *BucketsPerfCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *BucketsPerfCollector) Collect(ch chan<- prometheus.Metric) {
	bucketsperf := c.Client.GetBucketsPerformance(c.Buckets)
	if len(bucketsperf.Items) == 0 {
		return
	}

	for _, bp := range bucketsperf.Items {
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			bp.UsecPerOtherOp,
			bp.Name, "usec_per_other_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			bp.UsecPerReadOp,
			bp.Name, "usec_per_read_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			bp.UsecPerWriteOp,
			bp.Name, "usec_per_write_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			bp.OthersPerSec,
			bp.Name, "others_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			bp.ReadsPerSec,
			bp.Name, "reads_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			bp.WritesPerSec,
			bp.Name, "writes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			bp.ReadBytesPerSec,
			bp.Name, "read_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BandwidthDesc,
			prometheus.GaugeValue,
			bp.WriteBytesPerSec,
			bp.Name, "write_bytes_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			bp.BytesPerOp,
			bp.Name, "bytes_per_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			bp.BytesPerRead,
			bp.Name, "bytes_per_read",
		)
		ch <- prometheus.MustNewConstMetric(
			c.AverageSizeDesc,
			prometheus.GaugeValue,
			bp.BytesPerWrite,
			bp.Name, "bytes_per_write",
		)
	}
}

func NewBucketsPerfCollector(fb *client.FBClient,
	b *client.BucketsList) *BucketsPerfCollector {
	return &BucketsPerfCollector{
		LatencyDesc: prometheus.NewDesc(
			"purefb_buckets_performance_latency_usec",
			"FlashBlade buckets latency",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		ThroughputDesc: prometheus.NewDesc(
			"purefb_buckets_performance_throughput_iops",
			"FlashBlade buckets throughput",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		BandwidthDesc: prometheus.NewDesc(
			"purefb_buckets_performance_bandwidth_bytes",
			"FlashBlade buckets bandwidth",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		AverageSizeDesc: prometheus.NewDesc(
			"purefb_buckets_performance_average_bytes",
			"FlashBlade buckets average operations size",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		Client:  fb,
		Buckets: b,
	}
}
