package collectors

import (
    "purestorage/fb-openmetrics-exporter/internal/rest-client"
    "github.com/prometheus/client_golang/prometheus"
)

type NfsPerfCollector struct {
    LatencyDesc *prometheus.Desc
    ThroughputDesc *prometheus.Desc
    Client      *client.FBClient
}

func (c *NfsPerfCollector) Describe(ch chan<- *prometheus.Desc) {
    prometheus.DescribeByCollect(c, ch)
}

func (c *NfsPerfCollector) Collect(ch chan<- prometheus.Metric) {
    arraysnfsperf := c.Client.GetArraysNfsPerformance()
    if len(arraysnfsperf.Items) == 0 {
        return
    }

    np := arraysnfsperf.Items[0]
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.AccessesPerSec,
            "accesses_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.AggregateFileMetadataCreatesPerSec,
            "aggregate_file_metadata_creates_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.AggregateFileMetadataModifiesPerSec,
            "aggregate_file_metadata_modifies_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.AggregateFileMetadataReadsPerSec,
            "aggregate_file_metadata_reads_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.AggregateOtherPerSec,
            "aggregate_other_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.AggregateShareMetadataReadsPerSec,
            "aggregate_share_metadata_reads_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.CreatesPerSec,
            "creates_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.FsinfosPerSec,
            "fsinfos_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.FsstatsPerSec,
            "fsstats_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.GetattrsPerSec,
            "getattrs_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.LinksPerSec,
            "links_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.LookupsPerSec,
            "lookups_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.MkdirsPerSec,
            "mkdirs_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.PathconfsPerSec,
            "pathconfs_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.ReadsPerSec,
            "reads_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.ReaddirsPerSec,
            "readdirs_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.ReaddirplusesPerSec,
            "readdirpluses_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.ReadlinksPerSec,
            "readlinks_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.RemovesPerSec,
            "removes_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.RenamesPerSec,
            "renames_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.RmdirsPerSec,
            "rmdirs_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.SetattrsPerSec,
            "setattrs_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.SymlinksPerSec,
            "symlinks_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.ThroughputDesc,
            prometheus.GaugeValue,
            np.WritesPerSec,
            "writes_per_sec",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.AggregateUsecPerFileMetadataCreateOp,
            "aggregate_usec_per_file_metadata_create_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.AggregateUsecPerFileMetadataModifyOp,
            "aggregate_usec_per_file_metadata_modify_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.AggregateUsecPerFileMetadataReadOp,
            "aggregate_usec_per_file_metadata_read_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.AggregateUsecPerOtherOp,
            "aggregate_usec_per_other_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.AggregateUsecPerShareMetadataReadOp,
            "aggregate_usec_per_share_metadata_read_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerAccessOp,
            "usec_per_access_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerCreateOp,
            "usec_per_create_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerFsinfoOp,
            "usec_per_fsinfo_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerFsstatOp,
            "usec_per_fsstat_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerGetattrOp,
            "usec_per_getattr_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerLinkOp,
            "usec_per_link_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerLookupOp,
            "usec_per_lookup_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerMkdirOp,
            "usec_per_mkdir_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerPathconfOp,
            "usec_per_pathconf_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerReadOp,
            "usec_per_read_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerReaddirOp,
            "usec_per_readdir_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerReaddirplusOp,
            "usec_per_readdirplus_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerReadlinkOp,
            "usec_per_readlink_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerRemoveOp,
            "usec_per_remove_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerRenameOp,
            "usec_per_rename_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerRmdirOp,
            "usec_per_rmdir_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerSetattrOp,
            "usec_per_setattr_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerSymlinkOp,
            "usec_per_symlink_op",
    )
    ch <- prometheus.MustNewConstMetric(
            c.LatencyDesc,
            prometheus.GaugeValue,
            np.UsecPerWriteOp,
            "usec_per_write_op",
    )
}

func NewNfsPerfCollector(fb *client.FBClient) *NfsPerfCollector {
    return &NfsPerfCollector{
        LatencyDesc: prometheus.NewDesc(
            "purefb_array_nfs_specific_performance_latency_usec",
            "FlashBlade array NFS specific latency",
            []string{"dimension"},
            prometheus.Labels{},
        ),
        ThroughputDesc: prometheus.NewDesc(
            "purefb_array_nfs_specific_performance_throughput_iops",
            "FlashBlade array NFS specific throughput",
            []string{"dimension"},
            prometheus.Labels{},
        ),
        Client: fb,
    }
}
