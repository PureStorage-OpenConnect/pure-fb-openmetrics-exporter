package collectors

import (
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type BucketReplicaLinksCollector struct {
	ObjectBacklogBytesDesc     *prometheus.Desc
	ObjectBacklogDeleteOpsDesc *prometheus.Desc
	ObjectBacklogOtherOpsDesc  *prometheus.Desc
	ObjectBacklogPutOpsDesc    *prometheus.Desc
	LagDesc                    *prometheus.Desc

	TotalObjectBacklogBytesDesc     *prometheus.Desc
	TotalObjectBacklogDeleteOpsDesc *prometheus.Desc
	TotalObjectBacklogOtherOpsDesc  *prometheus.Desc
	TotalObjectBacklogPutOpsDesc    *prometheus.Desc
	TotalLagDesc                    *prometheus.Desc

	Client *client.FBClient
}

func (c *BucketReplicaLinksCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.ObjectBacklogBytesDesc
	ch <- c.ObjectBacklogDeleteOpsDesc
	ch <- c.ObjectBacklogOtherOpsDesc
	ch <- c.ObjectBacklogPutOpsDesc
	ch <- c.LagDesc

	ch <- c.TotalObjectBacklogBytesDesc
	ch <- c.TotalObjectBacklogDeleteOpsDesc
	ch <- c.TotalObjectBacklogOtherOpsDesc
	ch <- c.TotalObjectBacklogPutOpsDesc
	ch <- c.TotalLagDesc
}

func (c *BucketReplicaLinksCollector) Collect(ch chan<- prometheus.Metric) {
	links := c.Client.GetBucketReplicaLinks()

	for i := range links.Items {
		item := links.Items[i]
		if item.ObjectBacklog == nil || item.LocalBucket == nil || item.Remote == nil {
			continue
		}

		localBucket := item.LocalBucket.Name
		remote := item.Remote.Name
		direction := item.Direction

		ch <- prometheus.MustNewConstMetric(
			c.ObjectBacklogBytesDesc,
			prometheus.GaugeValue,
			item.ObjectBacklog.BytesCount,
			localBucket, remote, direction,
		)
		ch <- prometheus.MustNewConstMetric(
			c.ObjectBacklogDeleteOpsDesc,
			prometheus.GaugeValue,
			item.ObjectBacklog.DeleteOpsCount,
			localBucket, remote, direction,
		)
		ch <- prometheus.MustNewConstMetric(
			c.ObjectBacklogOtherOpsDesc,
			prometheus.GaugeValue,
			item.ObjectBacklog.OtherOpsCount,
			localBucket, remote, direction,
		)
		ch <- prometheus.MustNewConstMetric(
			c.ObjectBacklogPutOpsDesc,
			prometheus.GaugeValue,
			item.ObjectBacklog.PutOpsCount,
			localBucket, remote, direction,
		)
		ch <- prometheus.MustNewConstMetric(
			c.LagDesc,
			prometheus.GaugeValue,
			item.Lag,
			localBucket, remote, direction,
		)
	}

	if links.Total == nil || links.Total.ObjectBacklog == nil {
		return
	}

	ch <- prometheus.MustNewConstMetric(
		c.TotalObjectBacklogBytesDesc,
		prometheus.GaugeValue,
		links.Total.ObjectBacklog.BytesCount,
	)
	ch <- prometheus.MustNewConstMetric(
		c.TotalObjectBacklogDeleteOpsDesc,
		prometheus.GaugeValue,
		links.Total.ObjectBacklog.DeleteOpsCount,
	)
	ch <- prometheus.MustNewConstMetric(
		c.TotalObjectBacklogOtherOpsDesc,
		prometheus.GaugeValue,
		links.Total.ObjectBacklog.OtherOpsCount,
	)
	ch <- prometheus.MustNewConstMetric(
		c.TotalObjectBacklogPutOpsDesc,
		prometheus.GaugeValue,
		links.Total.ObjectBacklog.PutOpsCount,
	)
	ch <- prometheus.MustNewConstMetric(
		c.TotalLagDesc,
		prometheus.GaugeValue,
		links.Total.Lag,
	)
}

func NewBucketReplicaLinksCollector(fb *client.FBClient) *BucketReplicaLinksCollector {
	return &BucketReplicaLinksCollector{
		ObjectBacklogBytesDesc: prometheus.NewDesc(
			"purefb_bucket_replica_links_object_backlog_bytes",
			"FlashBlade bucket replica links object backlog (bytes_count) per link",
			[]string{"local_bucket", "remote", "direction"},
			prometheus.Labels{},
		),
		ObjectBacklogDeleteOpsDesc: prometheus.NewDesc(
			"purefb_bucket_replica_links_object_backlog_delete_ops",
			"FlashBlade bucket replica links object backlog (delete_ops_count) per link",
			[]string{"local_bucket", "remote", "direction"},
			prometheus.Labels{},
		),
		ObjectBacklogOtherOpsDesc: prometheus.NewDesc(
			"purefb_bucket_replica_links_object_backlog_other_ops",
			"FlashBlade bucket replica links object backlog (other_ops_count) per link",
			[]string{"local_bucket", "remote", "direction"},
			prometheus.Labels{},
		),
		ObjectBacklogPutOpsDesc: prometheus.NewDesc(
			"purefb_bucket_replica_links_object_backlog_put_ops",
			"FlashBlade bucket replica links object backlog (put_ops_count) per link",
			[]string{"local_bucket", "remote", "direction"},
			prometheus.Labels{},
		),
		LagDesc: prometheus.NewDesc(
			"purefb_bucket_replica_links_lag",
			"FlashBlade bucket replica links lag per link (as returned by FlashBlade API)",
			[]string{"local_bucket", "remote", "direction"},
			prometheus.Labels{},
		),
		TotalObjectBacklogBytesDesc: prometheus.NewDesc(
			"purefb_bucket_replica_links_total_object_backlog_bytes",
			"FlashBlade bucket replica links total object backlog (bytes_count) (as returned by FlashBlade API)",
			[]string{},
			prometheus.Labels{},
		),
		TotalObjectBacklogDeleteOpsDesc: prometheus.NewDesc(
			"purefb_bucket_replica_links_total_object_backlog_delete_ops",
			"FlashBlade bucket replica links total object backlog (delete_ops_count) (as returned by FlashBlade API)",
			[]string{},
			prometheus.Labels{},
		),
		TotalObjectBacklogOtherOpsDesc: prometheus.NewDesc(
			"purefb_bucket_replica_links_total_object_backlog_other_ops",
			"FlashBlade bucket replica links total object backlog (other_ops_count) (as returned by FlashBlade API)",
			[]string{},
			prometheus.Labels{},
		),
		TotalObjectBacklogPutOpsDesc: prometheus.NewDesc(
			"purefb_bucket_replica_links_total_object_backlog_put_ops",
			"FlashBlade bucket replica links total object backlog (put_ops_count) (as returned by FlashBlade API)",
			[]string{},
			prometheus.Labels{},
		),
		TotalLagDesc: prometheus.NewDesc(
			"purefb_bucket_replica_links_total_lag",
			"FlashBlade bucket replica links total lag (as returned by FlashBlade API)",
			[]string{},
			prometheus.Labels{},
		),
		Client: fb,
	}
}


