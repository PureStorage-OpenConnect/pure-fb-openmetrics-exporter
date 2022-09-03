package collectors

import (
	"github.com/prometheus/client_golang/prometheus"
	"purestorage/fb-openmetrics-exporter/internal/rest-client"
)

type BucketsSpaceCollector struct {
	ReductionDesc *prometheus.Desc
	SpaceDesc     *prometheus.Desc
	Buckets       *client.BucketsList
}

func (c *BucketsSpaceCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(c, ch)
}

func (c *BucketsSpaceCollector) Collect(ch chan<- prometheus.Metric) {
	if len(c.Buckets.Items) == 0 {
		return
	}
	for _, bucket := range c.Buckets.Items {
		ch <- prometheus.MustNewConstMetric(
			c.ReductionDesc,
			prometheus.GaugeValue,
			bucket.Space.DataReduction,
			bucket.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			bucket.Space.Snapshots,
			bucket.Name, "snapshots",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			bucket.Space.TotalPhysical,
			bucket.Name, "total_physical",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			bucket.Space.Unique,
			bucket.Name, "unique",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			bucket.Space.Virtual,
			bucket.Name, "virtual",
		)
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(bucket.ObjectCount),
			bucket.Name, "object_count",
		)
	}
}

func NewBucketsSpaceCollector(bl *client.BucketsList) *BucketsSpaceCollector {
	return &BucketsSpaceCollector{
		ReductionDesc: prometheus.NewDesc(
			"purefb_buckets_space_data_reduction_ratio",
			"FlashBlade buckets space data reduction",
			[]string{"name"},
			prometheus.Labels{},
		),
		SpaceDesc: prometheus.NewDesc(
			"purefb_buckets_space_bytes",
			"FlashBlade buckets space in bytes",
			[]string{"name", "space"},
			prometheus.Labels{},
		),
		Buckets: bl,
	}
}
