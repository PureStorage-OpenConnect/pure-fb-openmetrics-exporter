package collectors

import (
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type BucketsReplicaLinksCollector struct {
	BucketReplicaLagDesc              *prometheus.Desc
	BucketReplicaStatusDesc           *prometheus.Desc
	BucketReplicaBacklogBytesSizeDesc *prometheus.Desc
	BucketReplicaBacklogOpsSizeDesc   *prometheus.Desc
	Client                            *client.FBClient
}

func (c *BucketsReplicaLinksCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.BucketReplicaLagDesc
}

func (c *BucketsReplicaLinksCollector) Collect(ch chan<- prometheus.Metric) {
	brl := c.Client.GetBucketsReplicaLinks()
	blstate := 0.0
	for _, bucketlink := range brl.Items {
		switch status := bucketlink.Status; status {
		case "unhealthy":
			continue
		case "replicating":
			blstate = 1.0
		case "paused":
			blstate = 2.0
		default:
			blstate = 0.0
		}
		ch <- prometheus.MustNewConstMetric(
			c.BucketReplicaLagDesc,
			prometheus.GaugeValue,
			float64(bucketlink.Lag),
			bucketlink.Direction, bucketlink.LocalBucket.Name, bucketlink.Remote.Name, bucketlink.RemoteBucket.Name, bucketlink.Status,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BucketReplicaBacklogBytesSizeDesc,
			prometheus.GaugeValue,
			float64(bucketlink.ObjectBacklog.BytesCount),
			bucketlink.Direction, bucketlink.LocalBucket.Name, bucketlink.Remote.Name, bucketlink.RemoteBucket.Name, bucketlink.Status,
		)
		ch <- prometheus.MustNewConstMetric(
			c.BucketReplicaBacklogOpsSizeDesc,
			prometheus.GaugeValue,
			float64(bucketlink.ObjectBacklog.PutOpsCount),
			bucketlink.Direction, bucketlink.LocalBucket.Name, bucketlink.Remote.Name, bucketlink.RemoteBucket.Name, bucketlink.Status, "put_ops_count",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BucketReplicaBacklogOpsSizeDesc,
			prometheus.GaugeValue,
			float64(bucketlink.ObjectBacklog.DeleteOpsCount),
			bucketlink.Direction, bucketlink.LocalBucket.Name, bucketlink.Remote.Name, bucketlink.RemoteBucket.Name, bucketlink.Status, "delete_ops_count",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BucketReplicaBacklogOpsSizeDesc,
			prometheus.GaugeValue,
			float64(bucketlink.ObjectBacklog.OtherOpsCount),
			bucketlink.Direction, bucketlink.LocalBucket.Name, bucketlink.Remote.Name, bucketlink.RemoteBucket.Name, bucketlink.Status, "other_ops_count",
		)
		ch <- prometheus.MustNewConstMetric(
			c.BucketReplicaStatusDesc,
			prometheus.GaugeValue,
			blstate,
			bucketlink.Direction, bucketlink.LocalBucket.Name, bucketlink.Remote.Name, bucketlink.RemoteBucket.Name, bucketlink.Status, bucketlink.StatusDetails,
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
		BucketReplicaBacklogBytesSizeDesc: prometheus.NewDesc(
			"purefb_buckets_replica_backlog_bytes",
			"FlashBlade size of the objects in bytes that need to be replicated. This does not include the size of custom metadata.",
			[]string{"direction", "local_bucket_name", "remote_name", "remote_bucket_name", "status"},
			prometheus.Labels{},
		),
		BucketReplicaBacklogOpsSizeDesc: prometheus.NewDesc(
			"purefb_buckets_replica_backlog_ops",
			"FlashBlade number of other operations that need to be replicated.",
			[]string{"direction", "local_bucket_name", "remote_name", "remote_bucket_name", "status", "dimension"},
			prometheus.Labels{},
		),
		BucketReplicaStatusDesc: prometheus.NewDesc(
			"purefb_buckets_replica_status",
			"FlashBlade status of the replica link",
			[]string{"direction", "local_bucket_name", "remote_name", "remote_bucket_name", "status", "status_details"},
			prometheus.Labels{},
		),
		Client: fb,
	}
}
