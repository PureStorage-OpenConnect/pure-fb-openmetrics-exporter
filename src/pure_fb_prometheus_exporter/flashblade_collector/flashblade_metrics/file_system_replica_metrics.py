from prometheus_client.core import GaugeMetricFamily

class FileSystemReplicaMetrics():
    """
    Base class for FlashBlade Prometheus filesystem replica link metrics
    """
    def __init__(self, fb_client):
        self.replica_links_lag = None
        self.file_system_replica_links = fb_client.file_system_replica_links()

    def _replica_links_lag(self):
        """
        Create metrics of gauge type for file system replica link lag, with the
        local filesystem name, replication direction, remote array name,
        remote filesystem name and replication status  as labels.
        """
        self.replica_links_lag = GaugeMetricFamily(
                                              'purefb_file_system_links_lag_msec',
                                              'FlashBlade filesystem links lag',
                                              labels=['name',
                                                      'direction',
                                                      'remote_name',
                                                      'remote_file_system_name',
                                                      'status'])
        for f in self.file_system_replica_links:
            self.replica_links_lag.add_metric([f.local_file_system.name,
                                               f.direction,
                                               f.remote.name,
                                               f.remote_file_system.name,
                                               f.status], -1 if f.lag is None else f.lag)

    def get_metrics(self):
        self._replica_links_lag()
        yield self.replica_links_lag
