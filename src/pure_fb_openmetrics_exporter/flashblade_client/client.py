import urllib3
from pypureclient import flashblade, PureError

class FlashbladeClient():
    """
    This is a simple wrapper to the Pure REST API 2.x specifically meant
    to optimize the "scraping" of space and performance metrics by Prometheus.
    Each endpoint is scraped only once and the result cached internally, so
    that any subsequent call does not actually query the endpoint and uses
    instead the internal result.
    """
    def __init__(self, target, api_token, disable_ssl_warn=False):
        self._disable_ssl_warn = disable_ssl_warn
        self._arrays = None
        self._hardware = None
        self._alerts = None
        self._arrays_clients_performance = None
        self._arrays_performance = None
        self._arrays_http_specific_performance = None
        self._arrays_nfs_specific_performance = None
        self._arrays_s3_specific_performance = None
        self._arrays_space = None
        self._buckets = None
        self._buckets_performance = None
        self._buckets_s3_specific_performance = None
        self._bucket_replica_links = None
        self._file_system_replica_links = None
        self._file_systems = None
        self._file_systems_performance = None
        self._usage_groups = None
        self._usage_users = None
        self._hardware_connectors_performance = None
        if self._disable_ssl_warn:
            urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
        self.client = flashblade.Client(target=target, 
                                        api_token=api_token, 
                                        user_agent='Pure-FB-OpenMetrics-exporter/0.8')

    def __del__(self):
        del self._disable_ssl_warn 
        del self._arrays
        del self._hardware
        del self._alerts
        del self._arrays_clients_performance
        del self._arrays_performance
        del self._arrays_http_specific_performance
        del self._arrays_nfs_specific_performance
        del self._arrays_s3_specific_performance
        del self._arrays_space
        del self._buckets
        del self._buckets_performance
        del self._buckets_s3_specific_performance
        del self._bucket_replica_links
        del self._file_system_replica_links
        del self._file_systems
        del self._file_systems_performance
        del self._usage_groups
        del self._usage_users
        del self._hardware_connectors_performance
        del self.client

    def arrays(self):
        if self._arrays:
            return self._arrays
        self._arrays = []
        try:
            res = self.client.get_arrays()
            if isinstance(res, flashblade.ValidResponse):
                self._arrays = list(res.items)
        except:
            pass
        return self._arrays

    def hardware(self):
        if self._hardware:
            return self._hardware
        self._hardware = []
        try:
            res = self.client.get_hardware()
            if isinstance(res, flashblade.ValidResponse):
                self._hardware = list(res.items)
        except:
            pass
        return self._hardware

    def alerts(self):
        if self._alerts:
            return self._alerts
        self._alerts = []
        try:
            res = self.client.get_alerts(filter='state=\'open\'')
            if isinstance(res, flashblade.ValidResponse):
                self._alerts = list(res.items)
        except:
            pass
        return self._alerts

    def arrays_clients_performance(self):
        if self._arrays_clients_performance:
            return self._arrays_clients_performance
        self._arrays_clients_performance = []
        try:
            res = self.client.get_arrays_clients_performance()
            if isinstance(res, flashblade.ValidResponse):
                self._arrays_clients_performance = list(res.items)
        except:
            pass
        return self._arrays_clients_performance

    def arrays_performance(self):
        if self._arrays_performance:
            return self._arrays_performance
        array_perf = {}
        try:
            for p in ['all', 'http', 'nfs', 's3', 'smb']:
                array_perf[p] = None
                res = self.client.get_arrays_performance(protocol=p)
                if not isinstance(res, flashblade.ValidResponse):
                    continue
                array_perf[p] = next(res.items)
        except:
            pass
        self._arrays_performance = array_perf
        return self._arrays_performance

    def arrays_http_specific_performance(self):
        if self._arrays_http_specific_performance:
            return self._arrays_http_specific_performance
        self._arrays_http_specific_performance = []
        try:
            res = self.client.get_arrays_http_specific_performance()
            if isinstance(res, flashblade.ValidResponse):
               self._arrays_http_specific_performance = list(res.items)
        except:
            pass
        return self._arrays_http_specific_performance

    def arrays_nfs_specific_performance(self):
        if self._arrays_nfs_specific_performance:
            return self._arrays_nfs_specific_performance
        self._arrays_nfs_specific_performance = []
        try:
            res = self.client.get_arrays_nfs_specific_performance()
            if isinstance(res, flashblade.ValidResponse):
                self._arrays_nfs_specific_performance = list(res.items)
        except:
            pass
        return self._arrays_nfs_specific_performance

    def arrays_s3_specific_performance(self):
        if self._arrays_s3_specific_performance:
            return self._arrays_s3_specific_performance
        self._arrays_s3_specific_performance = []
        try:
            res = self.client.get_arrays_s3_specific_performance()
            if isinstance(res, flashblade.ValidResponse):
                self._arrays_s3_specific_performance = list(res.items)
        except:
            pass
        return self._arrays_s3_specific_performance

    def arrays_space(self):
        if self._arrays_space:
            return self._arrays_space
        array_space = {}
        try:
            for t in ['array', 'file-system', 'object-store']:
                array_space[t] = None
                res = self.client.get_arrays_space(type=t)
                if not isinstance(res, flashblade.ValidResponse):
                    continue
                array_space[t] = next(res.items)
        except:
            pass
        self._array_space = array_space
        return self._array_space

    def buckets(self):
        if self._buckets:
            return self._buckets
        self._buckets = []
        try:
            res = self.client.get_buckets()
            if isinstance(res, flashblade.ValidResponse):
                self._buckets = list(res.items)
        except:
            pass
        return self._buckets

    def buckets_performance(self):
        if self._buckets_performance:
            return self._buckets_performance
        buckets_perf = []
        try:
            buckets = self.buckets()
            b_names = []
            for b in buckets:
                b_names.append(b.name)
            # split buckets list into list of lists of 5 buckets each
            buckets_list = [b_names[i:i + 5] for i in range(0, len(b_names), 5)]
            for b_list in buckets_list:
                res = self.client.get_buckets_performance(names=b_list)
                if not isinstance(res, flashblade.ValidResponse):
                    continue
                buckets_perf.extend(res.items)
        except:
            pass
        self._buckets_performance = buckets_perf
        return self._buckets_performance

    def buckets_s3_specific_performance(self):
        if self._buckets_s3_specific_performance:
            return self._buckets_s3_specific_performance
        buckets_s3_perf = []
        try:
            buckets = self.buckets()
            b_names = []
            for b in buckets:
                b_names.append(b.name)
            # split buckets list into list of lists of 5 buckets each
            buckets_list = [b_names[i:i + 5] for i in range(0, len(b_names), 5)]
            for b_list in buckets_list:
                res = self.client.get_buckets_s3_specific_performance(names=b_list)
                if not isinstance(res, flashblade.ValidResponse):
                    continue
                buckets_s3_perf.extend(res.items)
        except:
            pass
        self._buckets_s3_specific_performance = buckets_s3_perf
        return self._buckets_s3_specific_performance

    def bucket_replica_links(self):
        if self._bucket_replica_links:
            return self._bucket_replica_links
        self._bucket_replica_links = []
        try:
            res = self.client.get_bucket_replica_links()
            if isinstance(res, flashblade.ValidResponse):
                self._bucket_replica_links = list(res.items)
        except:
            pass
        return self._bucket_replica_links


    def file_systems(self):
        if self._file_systems:
            return self._file_systems
        self._file_systems = []
        try:
            res = self.client.get_file_systems()
            if isinstance(res, flashblade.ValidResponse):
                self._file_systems = list(res.items)
        except:
            pass
        return self._file_systems

    def file_systems_performance(self):
        if self._file_systems_performance:
            return self._file_systems_performance
        file_systems_perf = []
        try:
            file_systems = self.file_systems()
            fs_names = []
            for fs in file_systems:
                # At this time only NFS performance is supported
                if fs.nfs.v3_enabled or fs.nfs.v4_1_enabled:
                    fs_names.append(fs.name)
            # Split file systems list into list of lists of 5 elements each
            file_systems_lists = [fs_names[i:i + 5] for i in range(0, len(fs_names), 5)]
            for fs_list in file_systems_lists:
                res = self.client.get_file_systems_performance(names=fs_list, protocol='nfs')
                if not isinstance(res, flashblade.ValidResponse):
                    continue
                file_systems_perf.extend(res.items)
        except:
            pass
        self._file_systems_performance = file_systems_perf
        return self._file_systems_performance

    def file_system_replica_links(self):
        if self._file_system_replica_links:
            return self._file_system_replica_links
        self._file_system_replica_links = []
        try:
            res = self.client.get_file_system_replica_links()
            if isinstance(res, flashblade.ValidResponse):
                self._file_system_replica_links = list(res.items)
        except:
            pass
        return self._file_system_replica_links

    def usage_groups(self):
        if self._usage_groups:
            return self._usage_groups
        usage_groups = []
        try: 
            file_systems = self.file_systems()
            fs_names = []
            for fs in file_systems:
                res = self.client.get_usage_groups(file_system_names=[fs.name])
                if not isinstance(res, flashblade.ValidResponse):
                    continue
                if len(res.items) == 0:
                    continue
                usage_groups.extend(list(res.items))
        except:
            pass
        self._usage_groups = usage_groups
        return self._usage_groups

    def usage_users(self):
        if self._usage_users:
            return self._usage_users
        usage_users = []
        try:
            file_systems = self.file_systems()
            fs_names = []
            for fs in file_systems:
                res = self.client.get_usage_users(file_system_names=[fs.name])
                if not isinstance(res, flashblade.ValidResponse):
                    continue
                if len(res.items) == 0:
                    continue
                usage_users.extend(list(res.items))
        except:
            pass
        self._usage_users = usage_users
        return self._usage_users

    def hardware_connectors_performance(self):
        if self._hardware_connectors_performance:
            return self._hardware_connectors_performance
        self._hardware_connectors_performance = []
        try:
            res = self.client.get_hardware_connectors_performance()
            if isinstance(res, flashblade.ValidResponse):
                self._hardware_connectors_performance = list(res.items)
        except:
            pass
        return self._hardware_connectors_performance
