from prometheus_client.core import GaugeMetricFamily


class FileSystemsSpaceMetrics():
    """
    Base class for FlashBlade Prometheus file systems space metrics
    """
    def __init__(self, fb_client):
        self.reduction = None
        self.space = None
        self.file_systems = fb_client.file_systems()

    def _space(self) -> None:
        """
        Create metrics of gauge type for file systems space indicators.
        """
        self.reduction = GaugeMetricFamily(
                             'purefb_file_systems_space_data_reduction_ratio',
                             'FlashBlade file systems space data reduction',
                              labels=['name', 'nfs', 'smb'])
        self.space = GaugeMetricFamily(
                           'purefb_file_systems_space_bytes',
                           'FlashBlade file systems space in bytes',
                           labels=['name', 'nfs', 'smb', 'space'])
        for fs in self.file_systems:
            v3 = '3' if fs.nfs.v3_enabled else ''
            v4 = '41' if fs.nfs.v4_1_enabled else ''
            nfs = ','.join([v3, v4]).lstrip(',').rstrip(',')
            smb = '3' if fs.smb.enabled else ''
            self.reduction.add_metric([fs.name, nfs, smb],
                                      fs.space.data_reduction or 0)
            self.space.add_metric([fs.name, nfs, smb, 'provisioned'],
                                  fs.provisioned or 0)
            self.space.add_metric([fs.name, nfs, smb, 'snapshots'], 
                                  fs.space.snapshots)
            self.space.add_metric([fs.name, nfs, smb, 'total_physical'], 
                                  fs.space.total_physical)
            self.space.add_metric([fs.name, nfs, smb, 'unique'], 
                                  fs.space.unique)
            self.space.add_metric([fs.name, nfs, smb, 'virtual'], 
                                  fs.space.virtual)

    def get_metrics(self) -> None:
        self._space()
        yield self.reduction
        yield self.space
