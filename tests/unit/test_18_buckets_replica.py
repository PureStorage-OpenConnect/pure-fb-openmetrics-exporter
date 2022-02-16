from pure_fb_openmetrics_exporter.flashblade_collector.flashblade_metrics.bucket_replica_metrics import BucketReplicaMetrics


def test_bucket_replica_name(fb_client):
    bucket_replica = BucketReplicaMetrics(fb_client)
    for m in bucket_replica.get_metrics():
        for s in m.samples:
            assert s.name in ['purefb_bucket_replica_links_lag_msec'] 

def test_bucket_replica_labels(fb_client):
    bucket_replica = BucketReplicaMetrics(fb_client)
    for m in bucket_replica.get_metrics():
        for s in m.samples:
            assert s.labels['name'] is not None
            assert s.labels['direction'] is not None
            assert s.labels['remote_name'] is not None
            assert s.labels['remote_bucket_name'] is not None
            assert s.labels['remote_account'] is not None
            assert s.labels['status'] is not None
            assert len(s.labels['name']) > 0

def test_bucket_replica_val(fb_client):
    bucket_replica = BucketReplicaMetrics(fb_client)
    for m in bucket_replica.get_metrics():
        for s in m.samples:
            assert s.value is not None
            assert s.value >= 0
