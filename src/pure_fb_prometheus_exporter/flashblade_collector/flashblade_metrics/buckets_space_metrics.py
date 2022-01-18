from prometheus_client.core import GaugeMetricFamily


class BucketsSpaceMetrics():
    """
    Base class for FlashBlade Prometheus buckets space metrics
    """
    def __init__(self, fb_client):
        self.reduction = None
        self.space = None
        self.objects = None
        self.buckets = fb_client.buckets()

    def _space(self) -> None:
        """
        Create metrics of gauge type for buckets space indicators.
        """
        self.reduction = GaugeMetricFamily(
                             'purefb_buckets_space_data_reduction_ratio',
                             'FlashBlade buckets space data reduction',
                              labels=['name'])
        self.space = GaugeMetricFamily('purefb_buckets_space_bytes',
                                       'FlashBlade buckets space in bytes',
                                       labels=['name', 'space'])
        self.objects = GaugeMetricFamily('purefb_buckets_space_objects',
                                       'FlashBlade buckets objects count',
                                       labels=['name'])
        for b in self.buckets:
            self.reduction.add_metric([b.name], b.space.data_reduction or 0)
            self.space.add_metric([b.name, 'snapshots'], b.space.snapshots)
            self.space.add_metric([b.name, 'total_physical'],
                                  b.space.total_physical)
            self.space.add_metric([b.name, 'unique'], b.space.unique)
            self.space.add_metric([b.name, 'virtual'], b.space.virtual)
            self.objects.add_metric([b.name], b.object_count)

    def get_metrics(self) -> None:
        self._space()
        yield self.reduction
        yield self.space
        yield self.objects
