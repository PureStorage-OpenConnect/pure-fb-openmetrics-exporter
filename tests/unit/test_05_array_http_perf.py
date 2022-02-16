from pure_fb_openmetrics_exporter.flashblade_collector.flashblade_metrics.array_http_specific_performance_metrics import ArrayHttpSpecificPerformanceMetrics


def test_array_http_perf_name(fb_client):
    array_http_perf = ArrayHttpSpecificPerformanceMetrics(fb_client)
    for m in array_http_perf.get_metrics():
        for s in m.samples:
            assert s.name in [
                       'purefb_array_http_specific_performance_latency_usec',
                       'purefb_array_http_specific_performance_throughput_iops']

def test_array_http_perf_labels(fb_client):
    array_http_perf = ArrayHttpSpecificPerformanceMetrics(fb_client)
    for m in array_http_perf.get_metrics():
        for s in m.samples:
            if s.name == 'purefb_array_http_specific_latency_msec':
                assert s.labels['dimension'] in ['usec_per_other_op',
                                                 'usec_per_read_dir_op',
                                                 'usec_per_read_file_op',
                                                 'usec_per_write_dir_op',
                                                 'usec_per_write_file_op']
            elif s.name == 'purefb_array_http_specific_throughput':
                assert s.labels['dimension'] in ['others_per_sec',
                                                 'read_dirs_per_sec',
                                                 'read_files_per_sec',
                                                 'write_dirs_per_sec',
                                                 'write_files_per_sec']
         
def test_array_http_perf_val(fb_client):
    array_http_perf = ArrayHttpSpecificPerformanceMetrics(fb_client)
    for m in array_http_perf.get_metrics():
        for s in m.samples:
            assert s.value is not None
            assert s.value >= 0
