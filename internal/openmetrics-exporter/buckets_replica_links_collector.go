package collectors

import (
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type BucketsReplicaLinksCollector struct {
	BucketReplicaLagDesc *prometheus.Desc
	Client               *client.FBClient
}

func (c *BucketsReplicaLinksCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.BucketReplicaLagDesc
}

func (c *BucketsReplicaLinksCollector) Collect(ch chan<- prometheus.Metric) {
	brl := c.Client.GetBucketsReplicaLinks()
	for _, bucketlink := range brl.Items {
		ch <- prometheus.MustNewConstMetric(
			c.BucketReplicaLagDesc,
			prometheus.GaugeValue,
			float64(bucketlink.Lag),
			bucketlink.Direction, bucketlink.LocalBucket.Name, bucketlink.Remote.Name, bucketlink.RemoteBucket.Name, bucketlink.Status,
		)
	}
}

func NewBucketsReplicaLinksCollector(fb *client.FBClient) *BucketsReplicaLinksCollector {
	return &BucketsReplicaLinksCollector{
		BucketReplicaLagDesc: prometheus.NewDesc(
			"purefb_buckets_replica_lag_msec",
			"FlashBlade duration in milliseconds that represents how far behind the replication target is from the source",
			[]string{"direction", "local_bucket_name", "remote_name", "remote_bucket_name", "status"},
			prometheus.Labels{},
		),
		Client: fb,
	}
}
