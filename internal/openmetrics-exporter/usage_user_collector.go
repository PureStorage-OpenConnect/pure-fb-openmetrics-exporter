package collectors

import (
	client "purestorage/fb-openmetrics-exporter/internal/rest-client"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type UsageUserCollector struct {
	UsageUsersDesc *prometheus.Desc
	Client         *client.FBClient
	FileSystems    *client.FileSystemsList
}

func (c *UsageUserCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.UsageUsersDesc
}

func (c *UsageUserCollector) Collect(ch chan<- prometheus.Metric) {
	uid := ""
	uusers := c.Client.GetUsageUsers(c.FileSystems)
	if len(uusers.Items) > 0 {
		for _, usage := range uusers.Items {
			uid = strconv.Itoa(usage.User.Id)
			ch <- prometheus.MustNewConstMetric(
				c.UsageUsersDesc,
				prometheus.GaugeValue,
				usage.Quota,
				usage.FileSystem.Name, usage.User.Name, uid, "quota",
			)
			ch <- prometheus.MustNewConstMetric(
				c.UsageUsersDesc,
				prometheus.GaugeValue,
				usage.Usage,
				usage.FileSystem.Name, usage.User.Name, uid, "usage",
			)
		}
	}
}

func NewUsageUserCollector(fb *client.FBClient,
	f *client.FileSystemsList) *UsageUserCollector {
	return &UsageUserCollector{
		UsageUsersDesc: prometheus.NewDesc(
			"purefb_file_system_usage_users_bytes",
			"FlashBlade file system users usage",
			[]string{"file_system", "user_name", "id", "dimension"},
			prometheus.Labels{},
		),
		Client:      fb,
		FileSystems: f,
	}
}
