package collectors

import (
    "strings"
    "purestorage/fb-openmetrics-exporter/internal/rest-client"
    "github.com/prometheus/client_golang/prometheus"
)

type FileSystemsSpaceCollector struct {
    ReductionDesc    *prometheus.Desc
    SpaceDesc        *prometheus.Desc
    FileSystems      *client.FileSystemsList
}

func (c *FileSystemsSpaceCollector) Describe(ch chan<- *prometheus.Desc) {
    prometheus.DescribeByCollect(c, ch)
}

func (c *FileSystemsSpaceCollector) Collect(ch chan<- prometheus.Metric) {
    if len(c.FileSystems.Items) == 0 {
        return
    }
    v3 := ""
    v4 := ""
    nfs := ""
    smb := ""
    for _, fs := range c.FileSystems.Items {
        if (fs.Nfs.V3Enabled) {
           v3 = "3"
        } else {
           v3 = ""
        }
        if (fs.Nfs.V41Enabled) {
           v4 = "41"
        } else {
           v4 = ""
        }
        nfs = strings.TrimRight(strings.TrimLeft(strings.Join([]string{v3 , v4}, ","), ","), ",")
        if (fs.Smb.Enabled) {
           smb = "3"
        } else {
           smb = ""
        }
        ch <- prometheus.MustNewConstMetric(
                c.ReductionDesc,
                prometheus.GaugeValue,
                fs.Space.DataReduction,
                fs.Name, nfs, smb,
        )
        ch <- prometheus.MustNewConstMetric(
                c.SpaceDesc,
                prometheus.GaugeValue,
                float64(fs.Provisioned),
                fs.Name, nfs, smb, "provisioned",
        )
        ch <- prometheus.MustNewConstMetric(
                c.SpaceDesc,
                prometheus.GaugeValue,
                fs.Space.Snapshots,
                fs.Name, nfs, smb, "snapshots",
        )
        ch <- prometheus.MustNewConstMetric(
                c.SpaceDesc,
                prometheus.GaugeValue,
                fs.Space.TotalPhysical,
                fs.Name, nfs, smb, "total_physical",
        )
        ch <- prometheus.MustNewConstMetric(
                c.SpaceDesc,
                prometheus.GaugeValue,
                fs.Space.Unique,
                fs.Name, nfs, smb, "unique",
        )
        ch <- prometheus.MustNewConstMetric(
                c.SpaceDesc,
                prometheus.GaugeValue,
                fs.Space.Virtual,
                fs.Name, nfs, smb, "virtual",
        )
    }
}

func NewFileSystemsSpaceCollector(fl *client.FileSystemsList) *FileSystemsSpaceCollector {
    return &FileSystemsSpaceCollector{
        ReductionDesc: prometheus.NewDesc(
            "purefb_file_systems_space_data_reduction_ratio",
            "FlashBlade file systems space data reduction",
            []string{"name", "nfs", "smb"},
            prometheus.Labels{},
        ),
        SpaceDesc: prometheus.NewDesc(
            "purefb_file_systems_space_bytes",
            "FlashBlade file systems space in bytes",
            []string{"name", "nfs", "smb", "space"},
            prometheus.Labels{},
        ),
        FileSystems: fl,
    }
}
