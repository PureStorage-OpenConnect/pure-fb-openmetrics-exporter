from .flashblade_metrics.array_clients_performance_metrics import ArrayClientsPerformanceMetrics
from .flashblade_metrics.array_events_metrics import ArrayEventsMetrics
from .flashblade_metrics.array_hardware_metrics import ArrayHardwareMetrics
from .flashblade_metrics.array_info_metrics import ArrayInfoMetrics
from .flashblade_metrics.array_http_specific_performance_metrics import ArrayHttpSpecificPerformanceMetrics
from .flashblade_metrics.array_nfs_specific_performance_metrics import ArrayNfsSpecificPerformanceMetrics
from .flashblade_metrics.array_performance_metrics import ArrayPerformanceMetrics
from .flashblade_metrics.array_s3_specific_performance_metrics import ArrayS3SpecificPerformanceMetrics
from .flashblade_metrics.array_space_metrics import ArraySpaceMetrics
from .flashblade_metrics.bucket_replica_metrics import BucketReplicaMetrics
from .flashblade_metrics.buckets_performance_metrics import BucketsPerformanceMetrics
from .flashblade_metrics.buckets_s3_specific_performance_metrics import BucketsS3SpecificPerformanceMetrics
from .flashblade_metrics.buckets_space_metrics import BucketsSpaceMetrics
from .flashblade_metrics.file_system_replica_metrics import FileSystemReplicaMetrics
from .flashblade_metrics.file_systems_performance_metrics import FileSystemsPerformanceMetrics
from .flashblade_metrics.file_systems_space_metrics import FileSystemsSpaceMetrics
from .flashblade_metrics.usage_users_metrics import UsageUsersMetrics
from .flashblade_metrics.usage_groups_metrics import UsageGroupsMetrics
from .flashblade_metrics.hardware_connectors_performance_metrics import HardwareConnectorsPerformanceMetrics


class FlashbladeCollector():
    """
    Instantiates the collector's methods and properties to retrieve status,
    space occupancy and performance metrics from Puretorage FlasBlade.
    Provides also a 'collect' method to allow Prometheus client registry
    to work properly.
    :param target: IP address or domain name of the target array's management
                   interface.
    :type target: str
    :type api_token: str
    """
    def __init__(self, fb_client, request='all'):
        self.request = request
        self.fb_client = fb_client

    def collect(self):
        """Global collector method for all the collected array metrics."""
        if self.request in ['all', 'array']:
            yield from ArrayInfoMetrics(self.fb_client).get_metrics()
            yield from ArrayHardwareMetrics(self.fb_client).get_metrics()
            yield from ArrayEventsMetrics(self.fb_client).get_metrics()
            yield from ArrayPerformanceMetrics(self.fb_client).get_metrics()
            yield from ArrayHttpSpecificPerformanceMetrics(self.fb_client).get_metrics()
            yield from ArrayNfsSpecificPerformanceMetrics(self.fb_client).get_metrics()
            yield from ArrayS3SpecificPerformanceMetrics(self.fb_client).get_metrics()
            yield from ArraySpaceMetrics(self.fb_client).get_metrics()
            yield from BucketsPerformanceMetrics(self.fb_client).get_metrics()
            yield from BucketsS3SpecificPerformanceMetrics(self.fb_client).get_metrics()
            yield from BucketsSpaceMetrics(self.fb_client).get_metrics()
            yield from FileSystemsPerformanceMetrics(self.fb_client).get_metrics()
            yield from FileSystemsSpaceMetrics(self.fb_client).get_metrics()
            yield from BucketReplicaMetrics(self.fb_client).get_metrics()
            yield from FileSystemReplicaMetrics(self.fb_client).get_metrics()
            yield from HardwareConnectorsPerformanceMetrics(self.fb_client).get_metrics()
        if self.request in ['all', 'usage']:
            yield from UsageUsersMetrics(self.fb_client).get_metrics()
            yield from UsageGroupsMetrics(self.fb_client).get_metrics()
        if self.request in ['all', 'clients']:
            yield from ArrayClientsPerformanceMetrics(self.fb_client).get_metrics()
