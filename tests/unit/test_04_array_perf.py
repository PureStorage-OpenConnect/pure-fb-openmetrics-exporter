from pure_fb_openmetrics_exporter.flashblade_collector.flashblade_metrics.array_performance_metrics import ArrayPerformanceMetrics


def test_array_perf_name(fb_client):
    array_perf = ArrayPerformanceMetrics(fb_client)
    for m in array_perf.get_metrics():
        for s in m.samples:
            assert s.name in ['purefb_array_performance_latency_usec',
                              'purefb_array_performance_throughput_iops',
                              'purefb_array_performance_bandwidth_bytes',
                              'purefb_array_performance_average_bytes']

def test_array_perf_labels(fb_client):
    array_perf = ArrayPerformanceMetrics(fb_client)
    for m in array_perf.get_metrics():
        for s in m.samples:
            assert s.labels['protocol'] in ['all', 'http', 'nfs', 's3', 'smb']
            if s.name == 'purefb_array_performance_latency_msec':
                assert s.labels['dimension'] in ['usec_per_read_op',
                                                 'usec_per_write_op',
                                                 'usec_per_other_op']
            elif s.name == 'purefb_array_performance_throughput_iops':
                assert s.labels['dimension'] in ['others_per_sec',
                                                 'reads_per_sec',
                                                 'writes_per_sec']
            elif s.name == 'purefb_array_performance_bandwidth_bytes':
                assert s.labels['dimension'] in ['read_bytes_per_sec',
                                                 'write_bytes_per_sec']
            elif s.name == 'purefb_array_performance_avg_size_bytes':
                assert s.labels['dimension'] in ['bytes_per_op',
                                                 'bytes_per_read',
                                                 'bytes_per_write']
         
def test_array_perf_val(fb_client):
    array_perf = ArrayPerformanceMetrics(fb_client)
    for m in array_perf.get_metrics():
        for s in m.samples:
            assert s.value is not None
            assert s.value >= 0
