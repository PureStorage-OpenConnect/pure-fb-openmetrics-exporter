from pure_fb_openmetrics_exporter.flashblade_collector.flashblade_metrics.buckets_s3_specific_performance_metrics import BucketsS3SpecificPerformanceMetrics


def test_buckets_s3_perf_name(fb_client):
    buckets_s3perf = BucketsS3SpecificPerformanceMetrics(fb_client)
    for m in buckets_s3perf.get_metrics():
        for s in m.samples:
            assert s.name in [
                       'purefb_buckets_s3_specific_performance_latency_usec',
                       'purefb_buckets_s3_specific_performance_throughput_iops']

def test_buckets_perf_labels(fb_client):
    buckets_s3perf = BucketsS3SpecificPerformanceMetrics(fb_client)
    for m in buckets_s3perf.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            if s.name == 'purefb_buckets_s3_specific_performance_latency_usec':
                assert s.labels['dimension'] in ['usec_per_read_bucket_op',
                                                 'usec_per_read_object_op',
                                                 'usec_per_write_bucket_op',
                                                 'usec_per_write_object_op',
                                                 'usec_per_other_op']
            elif s.name == 'purefb_buckets_s3_specific_performance_throughput_iops':
                assert s.labels['dimension'] in ['others_per_sec',
                                                 'read_buckets_per_sec',
                                                 'read_objects_per_sec',
                                                 'write_buckets_per_sec',
                                                 'write_objects_per_sec']
         
def test_buckets_perf_val(fb_client):
    buckets_s3perf = BucketsS3SpecificPerformanceMetrics(fb_client)
    for m in buckets_s3perf.get_metrics():
        for s in m.samples:
            assert s.value is not None
            assert s.value >= 0
