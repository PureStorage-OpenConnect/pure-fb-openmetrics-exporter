package collectors

import (
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type BucketsS3PerfCollector struct {
	LatencyDesc    *prometheus.Desc
	ThroughputDesc *prometheus.Desc
	Client         *client.FBClient
	Buckets        *client.BucketsList
}

func (c *BucketsS3PerfCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.LatencyDesc
	ch <- c.ThroughputDesc
}

func (c *BucketsS3PerfCollector) Collect(ch chan<- prometheus.Metric) {
	bucketsperf := c.Client.GetBucketsS3Performance(c.Buckets)
	if len(bucketsperf.Items) == 0 {
		return
	}

	for _, bp := range bucketsperf.Items {
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			bp.OthersPerSec,
			bp.Name, "others_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			bp.ReadBucketsPerSec,
			bp.Name, "read_buckets_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			bp.ReadObjectsPerSec,
			bp.Name, "read_objects_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			bp.WriteBucketsPerSec,
			bp.Name, "write_buckets_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.ThroughputDesc,
			prometheus.GaugeValue,
			bp.WriteObjectsPerSec,
			bp.Name, "write_objects_per_sec",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			bp.UsecPerOtherOp,
			bp.Name, "usec_per_other_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			bp.UsecPerReadBucketOp,
			bp.Name, "usec_per_read_bucket_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			bp.UsecPerReadObjectOp,
			bp.Name, "usec_per_read_object_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			bp.UsecPerWriteBucketOp,
			bp.Name, "usec_per_write_bucket_op",
		)
		ch <- prometheus.MustNewConstMetric(
			c.LatencyDesc,
			prometheus.GaugeValue,
			bp.UsecPerWriteObjectOp,
			bp.Name, "usec_per_write_object_op",
		)
	}
}

func NewBucketsS3PerfCollector(fb *client.FBClient,
	b *client.BucketsList) *BucketsS3PerfCollector {
	return &BucketsS3PerfCollector{
		LatencyDesc: prometheus.NewDesc(
			"purefb_buckets_s3_specific_performance_latency_usec",
			"FlashBlade buckets S3 specific latency",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		ThroughputDesc: prometheus.NewDesc(
			"purefb_buckets_s3_specific_performance_throughput_iops",
			"FlashBlade buckets S3 specific throughput",
			[]string{"name", "dimension"},
			prometheus.Labels{},
		),
		Client:  fb,
		Buckets: b,
	}
}
