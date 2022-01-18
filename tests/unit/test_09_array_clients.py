from pure_fb_prometheus_exporter.flashblade_collector.flashblade_metrics.array_clients_performance_metrics import ArrayClientsPerformanceMetrics


def test_array_clients_perf_name(fb_client):
    array_clients_perf = ArrayClientsPerformanceMetrics(fb_client)
    for m in array_clients_perf.get_metrics():
        for s in m.samples:
            assert s.name in [
                        'purefb_array_clients_performance_latency_usec',
                        'purefb_array_clients_performance_throughput_iops',
                        'purefb_array_clients_performance_bandwidth_bytes',
                        'purefb_array_clients_performance_avg_size_bytes']

def test_array_clients_perf_labels(fb_client):
    array_clients_perf = ArrayClientsPerformanceMetrics(fb_client)
    for m in array_clients_perf.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            if s.name == 'purefb_array_clients_performance_latency_msec':
                assert s.labels['dimension'] in ['usec_per_read_op',
                                                 'usec_per_write_op',
                                                 'usec_per_other_op']
            elif s.name == 'purefb_array_clients_performance_throughput_iops':
                assert s.labels['dimension'] in ['others_per_sec',
                                                 'reads_per_sec',
                                                 'writes_per_sec']
            elif s.name == 'purefb_array_clients_performance_bandwidth_bytes':
                assert s.labels['dimension'] in ['read_bytes_per_sec',
                                                 'write_bytes_per_sec']
            elif s.name == 'purefb_array_clients_performance_avg_size_bytes':
                assert s.labels['dimension'] in ['bytes_per_op',
                                                 'bytes_per_read',
                                                 'bytes_per_write']

def test_array_clients_perf_val(fb_client):
    array_clients_perf = ArrayClientsPerformanceMetrics(fb_client)
    for m in array_clients_perf.get_metrics():
        for s in m.samples:
            assert s.value is not None
            assert s.value >= 0
