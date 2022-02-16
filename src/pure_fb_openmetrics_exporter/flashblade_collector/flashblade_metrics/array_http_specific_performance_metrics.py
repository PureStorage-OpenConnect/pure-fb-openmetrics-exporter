from prometheus_client.core import GaugeMetricFamily


class ArrayHttpSpecificPerformanceMetrics():
    """
    Base class for FlashBlade Prometheus array HTTP specific performance metrics
    """
    def __init__(self, fb_client):
        self.latency = None
        self.throughput = None
        self.http_perf = fb_client.arrays_http_specific_performance()

    def _performance(self):
        """
        Create array HTTP specific performance metrics of gauge type.
        """
        self.latency = GaugeMetricFamily(
                            'purefb_array_http_specific_performance_latency_usec',
                            'FlashBlade array HTTP specific latency',
                            labels=['dimension'])
        self.throughput = GaugeMetricFamily(
                         'purefb_array_http_specific_performance_throughput_iops',
                         'FlashBlade array HTTP specific throughput',
                         labels=['dimension'])
        for p in self.http_perf:
            self.latency.add_metric(['usec_per_other_op'],
                                    p.usec_per_other_op)
            self.latency.add_metric(['usec_per_read_dir_op'],
                                    p.usec_per_read_dir_op)
            self.latency.add_metric(['usec_per_read_file_op'],
                                    p.usec_per_read_file_op)
            self.latency.add_metric(['usec_per_write_dir_op'],
                                    p.usec_per_write_dir_op)
            self.latency.add_metric(['usec_per_write_file_op'],
                                    p.usec_per_write_file_op)
            self.throughput.add_metric(['others_per_sec'],
                                    p.others_per_sec)
            self.throughput.add_metric(['read_dirs_per_sec'],
                                    p.read_dirs_per_sec)
            self.throughput.add_metric(['read_files_per_sec'],
                                    p.read_files_per_sec)
            self.throughput.add_metric(['write_dirs_per_sec'],
                                    p.write_dirs_per_sec)
            self.throughput.add_metric(['write_files_per_sec'],
                                    p.write_files_per_sec)

    def get_metrics(self):
        self._performance()
        yield self.latency
        yield self.throughput
