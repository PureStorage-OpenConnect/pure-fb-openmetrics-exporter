package collectors

import (
	"strconv"

	client "purestorage/fb-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type BucketsSpaceCollector struct {
	ReductionDesc         *prometheus.Desc
	SpaceDesc             *prometheus.Desc
	BucketQuotaDesc       *prometheus.Desc
	BucketObjectCountDesc *prometheus.Desc
	Buckets               *client.BucketsList
}

func (c *BucketsSpaceCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.ReductionDesc
	ch <- c.SpaceDesc
	ch <- c.BucketQuotaDesc
	ch <- c.BucketObjectCountDesc
}

func (c *BucketsSpaceCollector) Collect(ch chan<- prometheus.Metric) {
	if len(c.Buckets.Items) == 0 {
		return
	}
	for _, bucket := range c.Buckets.Items {
		if bucket.Space.DataReduction != nil {
			ch <- prometheus.MustNewConstMetric(
				c.ReductionDesc,
				prometheus.GaugeValue,
				float64(*bucket.Space.DataReduction),
				bucket.Name, bucket.Account.Name,
			)
		}
		if bucket.Space.Snapshots != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*bucket.Space.Snapshots),
				bucket.Name, bucket.Account.Name, "snapshots",
			)
		}
		if bucket.Space.TotalPhysical != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*bucket.Space.TotalPhysical),
				bucket.Name, bucket.Account.Name, "total_physical",
			)
		}
		if bucket.Space.Unique != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*bucket.Space.Unique),
				bucket.Name, bucket.Account.Name, "unique",
			)
		}
		if bucket.Space.Virtual != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*bucket.Space.Virtual),
				bucket.Name, bucket.Account.Name, "virtual",
			)
		}
		if bucket.Space.TotalProvisioned != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*bucket.Space.TotalProvisioned),
				bucket.Name, bucket.Account.Name, "total_provisioned",
			)
		}
		if bucket.Space.AvailableProvisioned != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*bucket.Space.AvailableProvisioned),
				bucket.Name, bucket.Account.Name, "available_provisioned",
			)
		}
		if bucket.Space.AvailableRatio != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*bucket.Space.AvailableRatio),
				bucket.Name, bucket.Account.Name, "available_ratio",
			)
		}
		if bucket.Space.Destroyed != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*bucket.Space.Destroyed),
				bucket.Name, bucket.Account.Name, "destroyed",
			)
		}
		if bucket.Space.DestroyedVirtual != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*bucket.Space.DestroyedVirtual),
				bucket.Name, bucket.Account.Name, "destroyed_virtual",
			)
		}
		if bucket.Space.Shared != nil {
			ch <- prometheus.MustNewConstMetric(
				c.SpaceDesc,
				prometheus.GaugeValue,
				float64(*bucket.Space.Shared),
				bucket.Name, bucket.Account.Name, "shared",
			)
		}
		ch <- prometheus.MustNewConstMetric(
			c.SpaceDesc,
			prometheus.GaugeValue,
			float64(bucket.ObjectCount),
			bucket.Name, bucket.Account.Name, "object_count",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BucketObjectCountDesc,
			prometheus.GaugeValue,
			float64(bucket.ObjectCount),
			bucket.Name, bucket.Account.Name,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BucketQuotaDesc,
			prometheus.GaugeValue,
			float64(bucket.QuotaLimit),
			bucket.Name, bucket.Account.Name, strconv.FormatBool(bucket.HardLimitEnabled),
		)
	}
}

func NewBucketsSpaceCollector(bl *client.BucketsList) *BucketsSpaceCollector {
	return &BucketsSpaceCollector{
		ReductionDesc: prometheus.NewDesc(
			"purefb_buckets_space_data_reduction_ratio",
			"FlashBlade buckets space data reduction",
			[]string{"name", "account"},
			prometheus.Labels{},
		),
		SpaceDesc: prometheus.NewDesc(
			"purefb_buckets_space_bytes",
			"FlashBlade buckets space in bytes",
			[]string{"name", "account", "space"},
			prometheus.Labels{},
		),
		BucketQuotaDesc: prometheus.NewDesc(
			"purefb_buckets_quota_space_bytes",
			"FlashBlade buckets quota space in bytes",
			[]string{"name", "account", "hard_limit_enabled"},
			prometheus.Labels{},
		),
		BucketObjectCountDesc: prometheus.NewDesc(
			"purefb_buckets_object_count",
			"FlashBlade buckets object count",
			[]string{"name", "account"},
			prometheus.Labels{},
		),
		Buckets: bl,
	}
}
