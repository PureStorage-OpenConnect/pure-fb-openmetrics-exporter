from prometheus_client.core import InfoMetricFamily

class ArrayInfoMetrics():
    """
    Base class for FlashBlade Prometheus array info
    """
    def __init__(self, fb_client):
        self.array_info = None
        self.array = fb_client.arrays()[0]

    def _array(self):
        """Assemble a simple information metric defining the scraped system."""

        self.array_info = InfoMetricFamily(
                                      'purefb',
                                      'FlashBlade system information',
                                      value={'array_name': self.array.name,
                                            'system_id': self.array.id,
                                            'os': self.array.os,
                                            'version': self.array.version
                                            })

    def get_metrics(self):
        self._array()
        yield self.array_info
