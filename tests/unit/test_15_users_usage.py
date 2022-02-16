from pure_fb_openmetrics_exporter.flashblade_collector.flashblade_metrics.usage_users_metrics import UsageUsersMetrics


def test_usage_users_name(fb_client):
    usage_users = UsageUsersMetrics(fb_client)
    for m in usage_users.get_metrics():
        for s in m.samples:
            assert s.name in ['purefb_filesystem_usage_users_bytes']

def test_usage_users_labels(fb_client):
    usage_users = UsageUsersMetrics(fb_client)
    dim = ['quota', 'usage']
    for m in usage_users.get_metrics():
        for s in m.samples:
            assert len(s.labels['file_system']) > 0
            assert int(s.labels['uid']) >= 0
            assert s.labels['dimension'] in dim
         
def test_usage_users_val(fb_client):
    usage_users = UsageUsersMetrics(fb_client)
    for m in usage_users.get_metrics():
        for s in m.samples:
            assert s.value is not None
            assert s.value >= 0
