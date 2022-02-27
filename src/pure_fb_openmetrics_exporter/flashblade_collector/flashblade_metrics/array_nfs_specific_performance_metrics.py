from prometheus_client.core import GaugeMetricFamily

class ArrayNfsSpecificPerformanceMetrics():
    """
    Base class for FlashBlade Prometheus array NFS specific performance metrics
    """

    def __init__(self, fb_client):
        self.latency = None
        self.throughput = None
        self.nfs_perf = fb_client.arrays_nfs_specific_performance()


    def _performance(self):
        """
        Create array NFS specific performance metrics of gauge type.
        """
        self.latency = GaugeMetricFamily(
                               'purefb_array_nfs_latency_usec',
                               'FlashBlade array NFS latency',
                               labels=['dimension'])
        self.throughput = GaugeMetricFamily(
                               'purefb_array_nfs_throughput_iops',
                               'FlashBlade array NFS throughput',
                               labels=['dimension'])
        for p in self.nfs_perf:
            metric = p.to_dict()
            for d in ['aggregate_usec_per_file_metadata_create_op',
                      'aggregate_usec_per_file_metadata_modify_op',
                      'aggregate_usec_per_file_metadata_read_op',
                      'aggregate_usec_per_other_op',
                      'aggregate_usec_per_share_metadata_read_op',
                      'usec_per_access_op',
                      'usec_per_create_op',
                      'usec_per_fsinfo_op',
                      'usec_per_fsstat_op',
                      'usec_per_getattr_op',
                      'usec_per_link_op',
                      'usec_per_lookup_op',
                      'usec_per_mkdir_op',
                      'usec_per_pathconf_op',
                      'usec_per_read_op',
                      'usec_per_readdir_op',
                      'usec_per_readdirplus_op',
                      'usec_per_readlink_op',
                      'usec_per_remove_op',
                      'usec_per_rename_op',
                      'usec_per_rmdir_op',
                      'usec_per_setattr_op',
                      'usec_per_symlink_op',
                      'usec_per_write_op']:
                self.latency.add_metric([d], metric[d] or 0)
            for d in ['accesses_per_sec',
                      'aggregate_file_metadata_creates_per_sec',
                      'aggregate_file_metadata_modifies_per_sec',
                      'aggregate_file_metadata_reads_per_sec',
                      'aggregate_other_per_sec',
                      'aggregate_share_metadata_reads_per_sec',
                      'creates_per_sec',
                      'fsinfos_per_sec',
                      'fsstats_per_sec',
                      'getattrs_per_sec',
                      'links_per_sec',
                      'lookups_per_sec',
                      'mkdirs_per_sec',
                      'pathconfs_per_sec',
                      'readdirpluses_per_sec',
                      'readdirs_per_sec',
                      'readlinks_per_sec',
                      'reads_per_sec',
                      'removes_per_sec',
                      'renames_per_sec',
                      'rmdirs_per_sec',
                      'setattrs_per_sec',
                      'symlinks_per_sec',
                      'writes_per_sec']:
                self.throughput.add_metric([d], metric[d] or 0)

    def get_metrics(self):
        self._performance()
        yield self.latency
        yield self.throughput
