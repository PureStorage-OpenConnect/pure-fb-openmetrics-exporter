from pure_fb_openmetrics_exporter.flashblade_collector.flashblade_metrics import array_events_metrics

def test_array_events_name(fb_client):
    array_ev = array_events_metrics.ArrayEventsMetrics(fb_client)
    for m in array_ev.get_metrics():
        for s in m.samples:
            assert s.name == 'purefb_alerts_open'

def test_array_events_labels(fb_client):
    array_ev = array_events_metrics.ArrayEventsMetrics(fb_client)
    for m in array_ev.get_metrics():
        for s in m.samples:
            assert len(s.labels['severity']) > 0
            assert len(s.labels['component_type']) > 0
            assert len(s.labels['component_name']) > 0

def test_array_events_value(fb_client):
    array_ev = array_events_metrics.ArrayEventsMetrics(fb_client)
    for m in array_ev.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
