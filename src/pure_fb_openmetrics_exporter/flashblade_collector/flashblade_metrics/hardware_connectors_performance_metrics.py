from prometheus_client.core import GaugeMetricFamily

class HardwareConnectorsPerformanceMetrics():
    """
    Base class for FlashBlade Prometheus hardware connectors performace metrics
    """
    def __init__(self, fb_client):
        self.bandwidth = None
        self.throughput = None
        self.errors = None
        self.hardware_conn_perf = fb_client.hardware_connectors_performance()

    def _performance(self):
        """
        Create hardware connectors performance metrics of gauge type.
        """
        self.bandwidth = GaugeMetricFamily(
                       'purefb_hardware_connectors_performance_bandwidth_bytes',
                       'FlashBlade hardware connectors performance bandwidth',
                        labels=['name', 'dimension'])
        self.throughput = GaugeMetricFamily(
                       'purefb_hardware_connectors_performance_throughput_pkts',
                       'FlashBlade hardware connectors performance throughputh',
                        labels=['name', 'dimension'])
        self.errors = GaugeMetricFamily(
                       'purefb_hardware_connectors_performance_errors',
                       'FlashBlade hardware connectors performance errors per sec',
                        labels=['name', 'dimension'])
        for p in self.hardware_conn_perf:
            self.bandwidth.add_metric([p.name, 'received_bytes_per_sec'],
                                     p.received_bytes_per_sec)
            self.bandwidth.add_metric([p.name, 'transmitted_bytes_per_sec'],
                                     p.transmitted_bytes_per_sec)
            self.throughput.add_metric([p.name, 'received_packets_per_sec'],
                                     p.received_packets_per_sec)
            self.throughput.add_metric([p.name, 'transmitted_packets_per_sec'],
                                     p.transmitted_packets_per_sec)
            self.errors.add_metric([p.name, 'other_errors_per_sec'],
                                   p.other_errors_per_sec)
            self.errors.add_metric([p.name, 'received_crc_errors_per_sec'],
                                   p.received_crc_errors_per_sec)
            self.errors.add_metric([p.name, 'received_frame_errors_per_sec'],
                                   p.received_frame_errors_per_sec)
            self.errors.add_metric([p.name,
                                   'transmitted_carrier_errors_per_sec'],
                                   p.transmitted_carrier_errors_per_sec)
            self.errors.add_metric([p.name,
                                   'transmitted_dropped_errors_per_sec'],
                                   p.transmitted_dropped_errors_per_sec)
            self.errors.add_metric([p.name, 'total_errors_per_sec'],
                                   p.total_errors_per_sec)

    def get_metrics(self):
        self._performance()
        yield self.bandwidth
        yield self.throughput
        yield self.errors
