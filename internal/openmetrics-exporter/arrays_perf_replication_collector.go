package collectors

import (
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type PerfReplicationCollector struct {
	ThroughputDesc        *prometheus.Desc
	ObjectBacklogBytes    *prometheus.Desc
	ObjectBacklogDeleteOp *prometheus.Desc
	ObjectBacklogOtherOp  *prometheus.Desc
	ObjectBacklogPutOp    *prometheus.Desc
	Client                *client.FBClient
}

func (c *PerfReplicationCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.ThroughputDesc
	ch <- c.ObjectBacklogBytes
	ch <- c.ObjectBacklogDeleteOp
	ch <- c.ObjectBacklogOtherOp
	ch <- c.ObjectBacklogPutOp
}

func (c *PerfReplicationCollector) Collect(ch chan<- prometheus.Metric) {
	arraysreplperf := c.Client.GetArraysPerformanceReplication()
	if len(arraysreplperf.Items) == 0 {
		return
	}

	arp := arraysreplperf.Items[0]
	cont := arp.GetContinuous()
	if cont == nil {
		return
	}
	ch <- prometheus.MustNewConstMetric(
		c.ThroughputDesc,
		prometheus.GaugeValue,
		cont.TransmittedBytesPerSec,
		"transmitted_bytes_per_sec",
	)
	ch <- prometheus.MustNewConstMetric(
		c.ThroughputDesc,
		prometheus.GaugeValue,
		cont.ReceivedBytesPerSec,
		"received_bytes_per_sec",
	)

	var (
		sumBytes  float64
		sumDelete float64
		sumOther  float64
		sumPut    float64
		found     bool
	)
	for i := range arraysreplperf.Items {
		item := arraysreplperf.Items[i]
		citem := item.GetContinuous()
		if citem == nil || citem.ObjectBacklog == nil {
			continue
		}
		found = true
		sumBytes += citem.ObjectBacklog.BytesCount
		sumDelete += citem.ObjectBacklog.DeleteOpsCount
		sumOther += citem.ObjectBacklog.OtherOpsCount
		sumPut += citem.ObjectBacklog.PutOpsCount
	}
	if !found {
		return
	}
	ch <- prometheus.MustNewConstMetric(
		c.ObjectBacklogBytes,
		prometheus.GaugeValue,
		sumBytes,
	)
	ch <- prometheus.MustNewConstMetric(
		c.ObjectBacklogDeleteOp,
		prometheus.GaugeValue,
		sumDelete,
	)
	ch <- prometheus.MustNewConstMetric(
		c.ObjectBacklogOtherOp,
		prometheus.GaugeValue,
		sumOther,
	)
	ch <- prometheus.MustNewConstMetric(
		c.ObjectBacklogPutOp,
		prometheus.GaugeValue,
		sumPut,
	)
}

func NewPerfReplicationCollector(fb *client.FBClient) *PerfReplicationCollector {
	return &PerfReplicationCollector{
		ThroughputDesc: prometheus.NewDesc(
			"purefb_array_performance_replication",
			"FlashBlade array replication throughput",
			[]string{"dimension"},
			prometheus.Labels{},
		),
		ObjectBacklogBytes: prometheus.NewDesc(
			"purefb_array_performance_replication_object_backlog_bytes",
			"FlashBlade replication object backlog (bytes)",
			[]string{},
			prometheus.Labels{},
		),
		ObjectBacklogDeleteOp: prometheus.NewDesc(
			"purefb_array_performance_replication_object_backlog_delete_ops",
			"FlashBlade replication object backlog (delete operations count)",
			[]string{},
			prometheus.Labels{},
		),
		ObjectBacklogOtherOp: prometheus.NewDesc(
			"purefb_array_performance_replication_object_backlog_other_ops",
			"FlashBlade replication object backlog (other operations count)",
			[]string{},
			prometheus.Labels{},
		),
		ObjectBacklogPutOp: prometheus.NewDesc(
			"purefb_array_performance_replication_object_backlog_put_ops",
			"FlashBlade replication object backlog (put operations count)",
			[]string{},
			prometheus.Labels{},
		),
		Client: fb,
	}
}
