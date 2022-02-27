from prometheus_client.core import GaugeMetricFamily

class ArraySpaceMetrics():
    """
    Base class for FlashBlade Prometheus array space metrics
    """
    def __init__(self, fb_client):
        self.reduction = None
        self.space = None
        self.parity = None
        self.array_space = fb_client.arrays_space()

    def _space(self) -> None:
        """
        Create metrics of gauge type for array space indicators.
        """
        self.reduction = GaugeMetricFamily(
                                       'purefb_array_space_data_reduction_ratio',
                                       'FlashBlade space data reduction',
                                       labels=['type'])
        self.space = GaugeMetricFamily('purefb_array_space_bytes',
                                       'FlashBlade space in bytes',
                                       labels=['type', 'space'])
        self.parity = GaugeMetricFamily('purefb_array_space_parity',
                                       'FlashBlade space parity',
                                       labels=['type'])

        for t in self.array_space:
            self.reduction.add_metric([t], 
                                 self.array_space[t].space.data_reduction or 0)
            self.space.add_metric([t, 'snapshots'],
                                  self.array_space[t].space.snapshots)
            self.space.add_metric([t, 'total_physical'],
                                  self.array_space[t].space.total_physical)
            self.space.add_metric([t, 'unique'],
                                  self.array_space[t].space.unique)
            self.space.add_metric([t, 'virtual'],
                                  self.array_space[t].space.virtual)
            self.space.add_metric([t, 'capacity'],
                                  self.array_space[t].capacity)
            self.parity.add_metric([t], self.array_space[t].parity)

    def get_metrics(self):
        self._space()
        yield self.reduction
        yield self.space
        yield self.parity
