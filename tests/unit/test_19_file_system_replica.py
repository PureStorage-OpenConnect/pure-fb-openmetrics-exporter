from pure_fb_prometheus_exporter.flashblade_collector.flashblade_metrics.file_system_replica_metrics import FileSystemReplicaMetrics


def test_file_system_replica_name(fb_client):
    file_system_replica = FileSystemReplicaMetrics(fb_client)
    for m in file_system_replica.get_metrics():
        for s in m.samples:
            assert s.name in ['purefb_file_system_links_lag_msec'] 

def test_file_system_replica_labels(fb_client):
    file_system_replica = FileSystemReplicaMetrics(fb_client)
    for m in file_system_replica.get_metrics():
        for s in m.samples:
            assert s.labels['name'] is not None
            assert s.labels['direction'] is not None
            assert s.labels['remote_name'] is not None
            assert s.labels['remote_file_system_name'] is not None
            assert s.labels['status'] is not None
            assert len(s.labels['name']) > 0

def test_file_system_replica_val(fb_client):
    file_system_replica = FileSystemReplicaMetrics(fb_client)
    for m in file_system_replica.get_metrics():
        for s in m.samples:
            assert s.value is not None
            assert s.value >= 0
