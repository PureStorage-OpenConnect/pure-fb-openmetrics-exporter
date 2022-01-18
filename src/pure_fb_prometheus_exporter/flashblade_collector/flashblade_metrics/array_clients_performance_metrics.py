from prometheus_client.core import GaugeMetricFamily


class ArrayClientsPerformanceMetrics():
    """
    Base class for FlashBlade Prometheus array clients performance metrics
    """

    def __init__(self, fb_client):
        self.array_clients_perf = None
        self.latency = None
        self.throughput = None
        self.bandwidth = None
        self.average_size = None
        self.array_clients_perf = fb_client.arrays_clients_performance()

    def _performance(self):
        """
        Create array clients performance metrics of gauge type.
        """
        self.latency = GaugeMetricFamily(
                             'purefb_array_clients_performance_latency_usec',
                             'FlashBlade array clients latency',
                             labels=['name', 'dimension'])
        self.throughput = GaugeMetricFamily(
                             'purefb_array_clients_performance_throughput_iops',
                             'FlashBlade array clients throughput',
                             labels=['name', 'dimension'])
        self.bandwidth = GaugeMetricFamily(
                             'purefb_array_clients_performance_bandwidth_bytes',
                             'FlashBlade array clients bandwidth',
                             labels=['name', 'dimension'])
        self.average_size = GaugeMetricFamily(
                              'purefb_array_clients_performance_avg_size_bytes',
                              'FlashBlade array clients average operations size',
                              labels=['name', 'dimension'])
        for p in self.array_clients_perf:
            self.latency.add_metric([p.name, 'usec_per_read_op'],
                                     p.usec_per_read_op)
            self.latency.add_metric([p.name, 'usec_per_write_op'],
                                     p.usec_per_write_op)
            self.latency.add_metric([p.name, 'usec_per_other_op'],
                                     p.usec_per_other_op)
            self.throughput.add_metric([p.name, 'others_per_sec'],
                                       p.others_per_sec)
            self.throughput.add_metric([p.name, 'reads_per_sec'],
                                       p.reads_per_sec)
            self.throughput.add_metric([p.name, 'writes_per_sec'],
                                       p.writes_per_sec)
            self.bandwidth.add_metric([p.name, 'read_bytes_per_sec'],
                                       p.read_bytes_per_sec)
            self.bandwidth.add_metric([p.name, 'write_bytes_per_sec'],
                                       p.write_bytes_per_sec)
            self.average_size.add_metric([p.name, 'bytes_per_op'],
                                         p.bytes_per_op)
            self.average_size.add_metric([p.name, 'bytes_per_read'],
                                         p.bytes_per_read)
            self.average_size.add_metric([p.name, 'bytes_per_write'],
                                         p.bytes_per_write)

    def get_metrics(self):
        self._performance()
        yield self.latency
        yield self.throughput
        yield self.bandwidth
        yield self.average_size
