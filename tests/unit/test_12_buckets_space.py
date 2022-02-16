from pure_fb_openmetrics_exporter.flashblade_collector.flashblade_metrics.buckets_space_metrics import BucketsSpaceMetrics


def test_buckets_space_name(fb_client):
    buckets_space = BucketsSpaceMetrics(fb_client)
    for m in buckets_space.get_metrics():
        for s in m.samples:
            assert s.name in ['purefb_buckets_space_data_reduction_ratio', 
                              'purefb_buckets_space_bytes',
                              'purefb_buckets_space_objects']

def test_buckets_space_labels(fb_client):
    buckets_space = BucketsSpaceMetrics(fb_client)
    for m in buckets_space.get_metrics():
        for s in m.samples:
            assert len(s.labels['name']) > 0
            if s.name == 'purefb_buckets_space_bytes':
                assert s.labels['space'] in ['snapshots',
                                             'total_physical',
                                             'unique',
                                             'virtual']
            elif s.name in ['purefb_buckets_space_data_reduction_ratio',
                            'purefb_buckets_space_objects']:
                assert list(s.labels.keys()) == ['name']
         
def test_buckets_space_val(fb_client):
    buckets_space = BucketsSpaceMetrics(fb_client)
    for m in buckets_space.get_metrics():
        for s in m.samples:
            assert s.value is not None
            assert s.value >= 0
