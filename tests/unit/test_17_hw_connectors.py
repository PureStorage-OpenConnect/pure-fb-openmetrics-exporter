from pure_fb_openmetrics_exporter.flashblade_collector.flashblade_metrics.hardware_connectors_performance_metrics import HardwareConnectorsPerformanceMetrics


def test_hw_connectors_perf_name(fb_client):
    hv_conn_perf = HardwareConnectorsPerformanceMetrics(fb_client)
    for m in hv_conn_perf.get_metrics():
        for s in m.samples:
            assert s.name in [
                     'purefb_hardware_connectors_performance_bandwidth_bytes',
                     'purefb_hardware_connectors_performance_throughput_pkts',
                     'purefb_hardware_connectors_performance_errors']


def test_hw_connectors_perf_labels(fb_client):
    hv_conn_perf = HardwareConnectorsPerformanceMetrics(fb_client)
    for m in hv_conn_perf.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            if s.name == 'purefb_hardware_connectors_performance_bandwidth_bytes':
                assert s.labels['dimension'] in ['received_bytes_per_sec',
                                                 'transmitted_bytes_per_sec']
            elif s.name == 'purefb_hardware_connectors_performance_throughput_pkts':
                assert s.labels['dimension'] in ['received_packets_per_sec',
                                                 'transmitted_packets_per_sec']
            elif s.name == 'purefb_hardware_connectors_performance_errors':
                assert s.labels['dimension'] in [
                                           'other_errors_per_sec',
                                           'received_crc_errors_per_sec',
                                           'received_frame_errors_per_sec',
                                           'transmitted_carrier_errors_per_sec',
                                           'transmitted_dropped_errors_per_sec',
                                           'total_errors_per_sec']
         
def test_hv_connectors_perf_val(fb_client):
    hv_conn_perf = HardwareConnectorsPerformanceMetrics(fb_client)
    for m in hv_conn_perf.get_metrics():
        for s in m.samples:
            assert s.value is not None
            assert s.value >= 0
