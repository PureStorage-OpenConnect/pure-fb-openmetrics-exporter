package collectors

import (
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"

	"github.com/prometheus/client_golang/prometheus"
)

type FileSystemsPerfCollector struct {
	LatencyDesc     *prometheus.Desc
	ThroughputDesc  *prometheus.Desc
	BandwidthDesc   *prometheus.Desc
	AverageSizeDesc *prometheus.Desc
	Client          *client.FBClient
	FileSystems     *client.FileSystemsList
}

func (c *FileSystemsPerfCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.LatencyDesc
	ch <- c.ThroughputDesc
	ch <- c.BandwidthDesc
	ch <- c.AverageSizeDesc
}

func (c *FileSystemsPerfCollector) Collect(ch chan<- prometheus.Metric) {
	if len(c.FileSystems.Items) == 0 {
		return
	}

	protocols := []string{"all", "SMB", "NFS"}
	for _, proto := range protocols {
		filesystemsperf := c.Client.GetFileSystemsPerformance(c.FileSystems, proto)
		for _, fp := range filesystemsperf.Items {
			ch <- prometheus.MustNewConstMetric(
				c.LatencyDesc,
				prometheus.GaugeValue,
				fp.UsecPerOtherOp,
				fp.Name, "usec_per_other_op", proto,
			)
			ch <- prometheus.MustNewConstMetric(
				c.LatencyDesc,
				prometheus.GaugeValue,
				fp.UsecPerReadOp,
				fp.Name, "usec_per_read_op", proto,
			)
			ch <- prometheus.MustNewConstMetric(
				c.LatencyDesc,
				prometheus.GaugeValue,
				fp.UsecPerWriteOp,
				fp.Name, "usec_per_write_op", proto,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ThroughputDesc,
				prometheus.GaugeValue,
				fp.OthersPerSec,
				fp.Name, "others_per_sec", proto,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ThroughputDesc,
				prometheus.GaugeValue,
				fp.ReadsPerSec,
				fp.Name, "reads_per_sec", proto,
			)
			ch <- prometheus.MustNewConstMetric(
				c.ThroughputDesc,
				prometheus.GaugeValue,
				fp.WritesPerSec,
				fp.Name, "writes_per_sec", proto,
			)
			ch <- prometheus.MustNewConstMetric(
				c.BandwidthDesc,
				prometheus.GaugeValue,
				fp.ReadBytesPerSec,
				fp.Name, "read_bytes_per_sec", proto,
			)
			ch <- prometheus.MustNewConstMetric(
				c.BandwidthDesc,
				prometheus.GaugeValue,
				fp.WriteBytesPerSec,
				fp.Name, "write_bytes_per_sec", proto,
			)
			ch <- prometheus.MustNewConstMetric(
				c.AverageSizeDesc,
				prometheus.GaugeValue,
				fp.BytesPerOp,
				fp.Name, "bytes_per_op", proto,
			)
			ch <- prometheus.MustNewConstMetric(
				c.AverageSizeDesc,
				prometheus.GaugeValue,
				fp.BytesPerRead,
				fp.Name, "bytes_per_read", proto,
			)
			ch <- prometheus.MustNewConstMetric(
				c.AverageSizeDesc,
				prometheus.GaugeValue,
				fp.BytesPerWrite,
				fp.Name, "bytes_per_write", proto,
			)
		}
	}
}

func NewFileSystemsPerfCollector(fb *client.FBClient,
	f *client.FileSystemsList) *FileSystemsPerfCollector {
	return &FileSystemsPerfCollector{
		LatencyDesc: prometheus.NewDesc(
			"purefb_file_systems_performance_latency_usec",
			"FlashBlade file systems latency",
			[]string{"name", "dimension", "protocol"},
			prometheus.Labels{},
		),
		ThroughputDesc: prometheus.NewDesc(
			"purefb_file_systems_performance_throughput_iops",
			"FlashBlade file systems throughput",
			[]string{"name", "dimension", "protocol"},
			prometheus.Labels{},
		),
		BandwidthDesc: prometheus.NewDesc(
			"purefb_file_systems_performance_bandwidth_bytes",
			"FlashBlade file systems bandwidth",
			[]string{"name", "dimension", "protocol"},
			prometheus.Labels{},
		),
		AverageSizeDesc: prometheus.NewDesc(
			"purefb_file_systems_performance_average_bytes",
			"FlashBlade file systems average operations size",
			[]string{"name", "dimension", "protocol"},
			prometheus.Labels{},
		),
		Client:      fb,
		FileSystems: f,
	}
}
