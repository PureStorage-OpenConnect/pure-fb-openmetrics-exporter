package collectors

import (
    "strconv"
    "purestorage/fb-openmetrics-exporter/internal/rest-client"
    "github.com/prometheus/client_golang/prometheus"
)

type UsageCollector struct {
    UsageUsersDesc     *prometheus.Desc
    UsageGroupsDesc    *prometheus.Desc
    Client             *client.FBClient
    FileSystems        *client.FileSystemsList
}

func (c *UsageCollector) Describe(ch chan<- *prometheus.Desc) {
    prometheus.DescribeByCollect(c, ch)
}

func (c *UsageCollector) Collect(ch chan<- prometheus.Metric) {
    uid := ""
    gid := ""
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

func NewUsageCollector(fb *client.FBClient, 
                       f *client.FileSystemsList) *UsageCollector {
    return &UsageCollector{
        UsageUsersDesc: prometheus.NewDesc(
            "purefb_file_system_usage_users_bytes",
            "FlashBlade file system users usage",
            []string{"file_system", "user_name", "id", "dimension"},
            prometheus.Labels{},
        ),
        UsageGroupsDesc: prometheus.NewDesc(
            "purefb_file_system_usage_groups_bytes",
            "FlashBlade file system groups usage",
            []string{"file_system", "group_name", "id", "dimension"},
            prometheus.Labels{},
        ),
        Client: fb,
        FileSystems: f,
    }
}
