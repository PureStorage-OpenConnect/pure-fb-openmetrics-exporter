from prometheus_client.core import GaugeMetricFamily

class BucketsS3SpecificPerformanceMetrics():
    """
    Base class for FlashBlade Prometheus buckets S3 specific performace metrics
    """
    def __init__(self, fb_client):
        self.latency = None
        self.throughput = None
        self.buckets_s3_perf = fb_client.buckets_s3_specific_performance()

    def _performance(self):
        """
        Create buckets S3 performance metrics of gauge type.
        """
        self.latency = GaugeMetricFamily(
                            'purefb_buckets_s3_specific_performance_latency_usec',
                            'FlashBlade buckets S3 specific latency',
                            labels=['name', 'dimension'])
        self.throughput = GaugeMetricFamily(
                         'purefb_buckets_s3_specific_performance_throughput_iops',
                         'FlashBlade buckets S3 specific throughput',
                         labels=['name', 'dimension'])
        for p in self.buckets_s3_perf:
             self.latency.add_metric([p.name, 'usec_per_read_bucket_op'],
                                     p.usec_per_read_bucket_op)
             self.latency.add_metric([p.name, 'usec_per_write_bucket_op'],
                                     p.usec_per_write_bucket_op)
             self.latency.add_metric([p.name, 'usec_per_read_object_op'],
                                     p.usec_per_read_object_op)
             self.latency.add_metric([p.name, 'usec_per_write_object_op'],
                                     p.usec_per_write_object_op)
             self.latency.add_metric([p.name, 'usec_per_other_op'],
                                     p.usec_per_other_op)
             self.throughput.add_metric([p.name, 'others_per_sec'],
                                        p.others_per_sec)
             self.throughput.add_metric([p.name, 'read_buckets_per_sec'],
                                        p.read_buckets_per_sec)
             self.throughput.add_metric([p.name, 'read_objects_per_sec'],
                                        p.read_objects_per_sec)
             self.throughput.add_metric([p.name, 'write_buckets_per_sec'],
                                        p.write_buckets_per_sec)
             self.throughput.add_metric([p.name, 'write_objects_per_sec'],
                                        p.write_objects_per_sec)

    def get_metrics(self):
        self._performance()
        yield self.latency
        yield self.throughput
