from prometheus_client.core import GaugeMetricFamily


class ArrayS3SpecificPerformanceMetrics():
    """
    Base class for FlashBlade Prometheus array S3 specific performance metrics
    """

    def __init__(self, fb_client):
        self.latency = None
        self.throughput = None
        self.s3_perf = fb_client.arrays_s3_specific_performance()

    def _performance(self):
        """
        Create array S3 specific performance metrics of gauge type.
        """
        self.latency = GaugeMetricFamily(
                            'purefb_array_s3_specific_performance_latency_usec',
                            'FlashBlade array S3 specific latency',
                            labels=['dimension'])
        self.throughput = GaugeMetricFamily(
                         'purefb_array_s3_specific_performance_throughput_iops',
                         'FlashBlade array S3 specific throughput',
                         labels=['dimension'])

        for p in self.s3_perf:
             self.latency.add_metric(['usec_per_read_bucket_op'],
                                     p.usec_per_read_bucket_op)
             self.latency.add_metric(['usec_per_write_bucket_op'],
                                     p.usec_per_write_bucket_op)
             self.latency.add_metric(['usec_per_read_object_op'],
                                     p.usec_per_read_object_op)
             self.latency.add_metric(['usec_per_write_object_op'],
                                     p.usec_per_write_object_op)
             self.latency.add_metric(['usec_per_other_op'],
                                     p.usec_per_other_op)
             self.throughput.add_metric(['others_per_sec'],
                                        p.others_per_sec)
             self.throughput.add_metric(['read_buckets_per_sec'],
                                        p.read_buckets_per_sec)
             self.throughput.add_metric(['read_objects_per_sec'],
                                        p.read_objects_per_sec)
             self.throughput.add_metric(['write_buckets_per_sec'],
                                        p.write_buckets_per_sec)
             self.throughput.add_metric(['write_objects_per_sec'],
                                        p.write_objects_per_sec)

    def get_metrics(self):
        self._performance()
        yield self.latency
        yield self.throughput
