package collectors

import (
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type S3PerfCollector struct {
	LatencyDesc    *prometheus.Desc
	ThroughputDesc *prometheus.Desc
	Client         *client.FBClient
}

func (c *S3PerfCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.LatencyDesc
	ch <- c.ThroughputDesc
}

func (c *S3PerfCollector) Collect(ch chan<- prometheus.Metric) {
	arrayss3perf := c.Client.GetArraysS3Performance()

	if len(arrayss3perf.Items) == 0 {
		return
	}

	s3p := arrayss3perf.Items[0]
	ch <- prometheus.MustNewConstMetric(
		c.ThroughputDesc,
		prometheus.GaugeValue,
		s3p.OthersPerSec,
		"others_per_sec",
	)
	ch <- prometheus.MustNewConstMetric(
		c.ThroughputDesc,
		prometheus.GaugeValue,
		s3p.ReadBucketsPerSec,
		"read_buckets_per_sec",
	)
	ch <- prometheus.MustNewConstMetric(
		c.ThroughputDesc,
		prometheus.GaugeValue,
		s3p.ReadObjectsPerSec,
		"read_objects_per_sec",
	)
	ch <- prometheus.MustNewConstMetric(
		c.ThroughputDesc,
		prometheus.GaugeValue,
		s3p.WriteBucketsPerSec,
		"write_buckets_per_sec",
	)
	ch <- prometheus.MustNewConstMetric(
		c.ThroughputDesc,
		prometheus.GaugeValue,
		s3p.WriteObjectsPerSec,
		"write_objects_per_sec",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		s3p.UsecPerOtherOp,
		"usec_per_other_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		s3p.UsecPerReadBucketOp,
		"usec_per_read_bucket_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		s3p.UsecPerReadObjectOp,
		"usec_per_read_object_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		s3p.UsecPerWriteBucketOp,
		"usec_per_write_bucket_op",
	)
	ch <- prometheus.MustNewConstMetric(
		c.LatencyDesc,
		prometheus.GaugeValue,
		s3p.UsecPerWriteObjectOp,
		"usec_per_write_object_op",
	)

}

func NewS3PerfCollector(fb *client.FBClient) *S3PerfCollector {
	return &S3PerfCollector{
		LatencyDesc: prometheus.NewDesc(
			"purefb_array_s3_performance_latency_usec",
			"FlashBlade array latency",
			[]string{"dimension"},
			prometheus.Labels{},
		),
		ThroughputDesc: prometheus.NewDesc(
			"purefb_array_s3_performance_throughput_iops",
			"FlashBlade array throughput",
			[]string{"dimension"},
			prometheus.Labels{},
		),
		Client: fb,
	}
}
