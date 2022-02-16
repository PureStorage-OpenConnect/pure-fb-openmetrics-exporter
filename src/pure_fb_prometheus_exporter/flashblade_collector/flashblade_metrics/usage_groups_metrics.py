from prometheus_client.core import GaugeMetricFamily


class UsageGroupsMetrics():
    """
    Base class for FlashBlade Prometheus groups usage metrics
    """
    def __init__(self, fb_client):
        self.usage = None
        self.usage_groups = fb_client.usage_groups()

    def _usage(self):
        """
        Create metrics of gauge type for groups usage indicators.
        """
        self.usage = GaugeMetricFamily('purefb_file_system_usage_groups_bytes',
                                       'FlashBlade file system groups usage',
                                       labels=['file_system',
                                               'group_name',
                                               'uid',
                                               'dimension'])
        for gu in self.usage_groups:
            gname = gu.group.name or ''
            gid = str(gu.group.id)
            quota = gu.quota or 0
            usage = gu.usage or 0
            self.usage.add_metric([gu.file_system.name, gname, gid, 'quota'], 
                                  quota)
            self.usage.add_metric([gu.file_system.name, gname, gid, 'usage'],
                                  usage) 

    def get_metrics(self):
        self._usage()
        yield self.usage
