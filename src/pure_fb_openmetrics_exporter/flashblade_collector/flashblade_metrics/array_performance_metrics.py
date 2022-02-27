from prometheus_client.core import GaugeMetricFamily

class ArrayPerformanceMetrics():
    """
    Base class for FlashBlade Prometheus array performance metrics
    """

    def __init__(self, fb_client):
        self.latency = None
        self.throughput = None
        self.bandwidth = None
        self.average_size = None
        self.array_perf = fb_client.arrays_performance()

    def _performance(self):
        """
        Create array performance metrics of gauge type.
        """
        self.latency = GaugeMetricFamily(
                               'purefb_array_performance_latency_usec',
                               'FlashBlade array latency',
                               labels=['protocol', 'dimension'])
        self.throughput = GaugeMetricFamily(
                               'purefb_array_performance_throughput_iops',
                               'FlashBlade array throughput',
                               labels=['protocol', 'dimension'])
        self.bandwidth = GaugeMetricFamily(
                               'purefb_array_performance_bandwidth_bytes',
                               'FlashBlade array bandwidth',
                               labels=['protocol', 'dimension'])
        self.average_size = GaugeMetricFamily(
                                  'purefb_array_performance_average_bytes',
                                  'FlashBlade array average operations size',
                                  labels=['protocol', 'dimension'])
        for p in self.array_perf:
            self.latency.add_metric([p, 'usec_per_read_op'],
                                     self.array_perf[p].usec_per_read_op)
            self.latency.add_metric([p, 'usec_per_write_op'],
                                     self.array_perf[p].usec_per_write_op)
            self.latency.add_metric([p, 'usec_per_other_op'],
                                     self.array_perf[p].usec_per_other_op)
            self.throughput.add_metric([p, 'others_per_sec'],
                                       self.array_perf[p].others_per_sec)
            self.throughput.add_metric([p, 'reads_per_sec'],
                                       self.array_perf[p].reads_per_sec)
            self.throughput.add_metric([p, 'writes_per_sec'],
                                       self.array_perf[p].writes_per_sec)
            self.bandwidth.add_metric([p, 'read_bytes_per_sec'],
                                       self.array_perf[p].read_bytes_per_sec)
            self.bandwidth.add_metric([p, 'write_bytes_per_sec'],
                                       self.array_perf[p].write_bytes_per_sec)
            self.average_size.add_metric([p, 'bytes_per_op'],
                                         self.array_perf[p].bytes_per_op)
            self.average_size.add_metric([p, 'bytes_per_read'],
                                         self.array_perf[p].bytes_per_read)
            self.average_size.add_metric([p, 'bytes_per_write'],
                                         self.array_perf[p].bytes_per_write)

    def get_metrics(self):
        self._performance()
        yield self.latency
        yield self.throughput
        yield self.bandwidth
        yield self.average_size
