package collectors

import (
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type UsageGroupCollector struct {
	UsageGroupsDesc *prometheus.Desc
	Client          *client.FBClient
	FileSystems     *client.FileSystemsList
}

func (c *UsageGroupCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.UsageGroupsDesc
}

func (c *UsageGroupCollector) Collect(ch chan<- prometheus.Metric) {
	gid := ""
	ugroups := c.Client.GetUsageGroups(c.FileSystems)
	if len(ugroups.Items) > 0 {
		for _, usage := range ugroups.Items {
			gid = strconv.Itoa(usage.Group.Id)
			ch <- prometheus.MustNewConstMetric(
				c.UsageGroupsDesc,
				prometheus.GaugeValue,
				usage.Quota,
				usage.FileSystem.Name, usage.Group.Name, gid, "quota",
			)
			ch <- prometheus.MustNewConstMetric(
				c.UsageGroupsDesc,
				prometheus.GaugeValue,
				usage.Usage,
				usage.FileSystem.Name, usage.Group.Name, gid, "usage",
			)
		}
	}
}

func NewUsageGroupCollector(fb *client.FBClient,
	f *client.FileSystemsList) *UsageGroupCollector {
	return &UsageGroupCollector{
		UsageGroupsDesc: prometheus.NewDesc(
			"purefb_file_system_usage_groups_bytes",
			"FlashBlade file system groups usage",
			[]string{"file_system", "group_name", "id", "dimension"},
			prometheus.Labels{},
		),
		Client:      fb,
		FileSystems: f,
	}
}
