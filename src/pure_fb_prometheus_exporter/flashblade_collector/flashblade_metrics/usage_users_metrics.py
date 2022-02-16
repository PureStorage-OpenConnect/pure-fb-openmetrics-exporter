from prometheus_client.core import GaugeMetricFamily


class UsageUsersMetrics():
    """
    Base class for FlashBlade Prometheus users quota metrics
    """
    def __init__(self, fb_client):
        self.usage = None
        self.usage_users = fb_client.usage_users()

    def _usage(self):
        """
        Create metrics of gauge type for users usage indicators.
        """
        self.usage = GaugeMetricFamily('purefb_file_system_usage_users_bytes',
                                       'FlashBlade filesystem users usage',
                                       labels=['file_system',
                                               'user_name',
                                               'uid',
                                               'dimension'])
        for uu in self.usage_users:
            uname = uu.user.name or ''
            uid = str(uu.user.id)
            quota = uu.quota or 0
            usage = uu.usage or 0
            self.usage.add_metric([uu.file_system.name, uname, uid, 'quota'], 
                                  quota)
            self.usage.add_metric([uu.file_system.name, uname, uid, 'usage'],
                                  usage) 

    def get_metrics(self):
        self._usage()
        yield self.usage
