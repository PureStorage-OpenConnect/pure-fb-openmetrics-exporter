from pure_fb_openmetrics_exporter.flashblade_collector.flashblade_metrics.array_nfs_specific_performance_metrics import ArrayNfsSpecificPerformanceMetrics


def test_array_nfs_perf_name(fb_client):
    array_nfs_perf = ArrayNfsSpecificPerformanceMetrics(fb_client)
    for m in array_nfs_perf.get_metrics():
        for s in m.samples:
            assert s.name in ['purefb_array_nfs_latency_usec',
                              'purefb_array_nfs_throughput_iops']

def test_array_nfs_perf_labels(fb_client):
    array_nfs_perf = ArrayNfsSpecificPerformanceMetrics(fb_client)
    for m in array_nfs_perf.get_metrics():
        for s in m.samples:
            if s.name == 'purefb_array_nfs_latency_usec':
                assert s.labels['dimension'] in [
                                  'aggregate_usec_per_file_metadata_create_op',
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
                                  'usec_per_write_op']
            if s.name == 'purefb_array_nfs_throughput_iops':
                assert s.labels['dimension'] in [
                                  'accesses_per_sec',
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
                                  'writes_per_sec']
         
def test_array_nfs_perf_val(fb_client):
    array_nfs_perf = ArrayNfsSpecificPerformanceMetrics(fb_client)
    for m in array_nfs_perf.get_metrics():
        for s in m.samples:
            assert s.value is not None
            assert s.value >= 0
