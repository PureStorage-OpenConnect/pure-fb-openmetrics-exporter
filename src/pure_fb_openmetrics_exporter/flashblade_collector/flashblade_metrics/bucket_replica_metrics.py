from prometheus_client.core import GaugeMetricFamily

class BucketReplicaMetrics():
    """
    Base class for FlashBlade Prometheus buckets replication metrics
    """
    def __init__(self, fb_client):
        self.replica_links = None
        self.bucket_replica_links = fb_client.bucket_replica_links()

    def _replica_links(self):
        """
        Create metrics of gauge type for bucket  indicators, with the
        account name and the bucket name as labels.
        """
        self.replica_links = GaugeMetricFamily(
                                       'purefb_bucket_replica_links_lag_msec',
                                       'FlashBlade bucket replica links lag',
                                       labels=['name',
                                               'direction',
                                               'remote_name',
                                               'remote_bucket_name',
                                               'remote_account',
                                               'status'])
        for l in self.bucket_replica_links:
            self.replica_links.add_metric([l.local_bucket.name,
                                           l.direction,
                                           l.remote.name,
                                           l.remote_bucket.name, 
                                           l.remote_credentials.name,
                                           l.status], -1 if l.lag is None else l.lag)

    def get_metrics(self):
        self._replica_links()
        yield self.replica_links
