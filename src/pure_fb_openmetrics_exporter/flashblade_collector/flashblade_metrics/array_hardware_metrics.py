from prometheus_client.core import GaugeMetricFamily

class ArrayHardwareMetrics():
    """
    Base class for FlashBlade Prometheus hardware metrics
    """
    def __init__(self, fb_client):
        self.array_hardware_status = None
        self.hardware = fb_client.hardware()

    def _array_hardware_status(self):
        """
        Create metric of gauge types for components status.

        """
        self.array_hardware_status = GaugeMetricFamily(
                                  'purefb_hardware_health',
                                  'FlashBlade hardware component health status',
                                   labels=['type', 'name', 'index', 'slot'])

        for comp in self.hardware:
            if comp.status == 'not_installed':
                continue
            c_name = comp.name
            c_index = comp.index or ''
            c_slot = comp.slot or ''
            
            if comp.status == 'healthy':
                c_state = 1
            elif comp.status == 'unused':
                c_state = 2
            else:
                c_state = 0
            c_type = comp.type
            self.array_hardware_status.add_metric([c_type,
                                                   c_name,
                                                   str(c_index),
                                                   str(c_slot)], c_state)

    def get_metrics(self):
        self._array_hardware_status()
        yield self.array_hardware_status
