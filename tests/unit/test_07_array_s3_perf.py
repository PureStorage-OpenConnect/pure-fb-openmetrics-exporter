from pure_fb_prometheus_exporter.flashblade_collector.flashblade_metrics.array_s3_specific_performance_metrics import ArrayS3SpecificPerformanceMetrics


def test_array_s3_perf_name(fb_client):
    array_s3_perf = ArrayS3SpecificPerformanceMetrics(fb_client)
    for m in array_s3_perf.get_metrics():
        for s in m.samples:
            assert s.name in [
                         'purefb_array_s3_specific_performance_latency_usec',
                         'purefb_array_s3_specific_performance_throughput_iops']

def test_array_s3_perf_labels(fb_client):
    array_s3_perf = ArrayS3SpecificPerformanceMetrics(fb_client)
    for m in array_s3_perf.get_metrics():
        for s in m.samples:
            if s.name == 'purefb_array_s3_specific_performance_latency_usec':
                assert s.labels['dimension'] in ['usec_per_read_bucket_op',
                                                 'usec_per_read_object_op',
                                                 'usec_per_write_bucket_op',
                                                 'usec_per_write_object_op',
                                                 'usec_per_other_op']
            elif s.name == 'purefb_array_s3_specific_performance_throughput_iops':
                assert s.labels['dimension'] in ['others_per_sec',
                                                 'read_buckets_per_sec',
                                                 'read_objects_per_sec',
                                                 'write_buckets_per_sec',
                                                 'write_objects_per_sec']
         
def test_array_s3_perf_val(fb_client):
    array_s3_perf = ArrayS3SpecificPerformanceMetrics(fb_client)
    for m in array_s3_perf.get_metrics():
        for s in m.samples:
            assert s.value is not None
            assert s.value >= 0
