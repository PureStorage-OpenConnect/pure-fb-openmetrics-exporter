from prometheus_client.core import GaugeMetricFamily

class ArrayEventsMetrics():
    """
    Base class for FlashBlade Prometheus events metrics
    """
    def __init__(self, fb_client):
        self.open_events = None
        self.alerts = fb_client.alerts()

    def _open_events(self) -> None:
        """
        Create a metric of gauge type for the number of open alerts:
        critical, warning and info, with the severity as label.
        """

        self.open_events = GaugeMetricFamily('purefb_alerts_open',
                                             'Open alert events',
                                             labels=['severity',
                                                     'component_type',
                                                     'component_name'])

        for a in self.alerts:
            self.open_events.add_metric([a.severity,
                                         a.component_type,
                                         a.component_name], 1.0)
    def get_metrics(self) -> None:
        self._open_events()
        yield self.open_events
