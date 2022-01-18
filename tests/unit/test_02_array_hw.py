from pure_fb_prometheus_exporter.flashblade_collector.flashblade_metrics import array_hardware_metrics

def test_array_hw_name(fb_client):
    array_hw = array_hardware_metrics.ArrayHardwareMetrics(fb_client)
    for m in array_hw.get_metrics():
        for s in m.samples:
            assert s.name in ['purefb_hardware_health']

def test_array_hw_labels(fb_client):
    array_hw = array_hardware_metrics.ArrayHardwareMetrics(fb_client)
    for m in array_hw.get_metrics():
        for s in m.samples:
            assert len(s.labels['chassis']) > 0
            assert len(s.labels['fabric_module']) >= 0
            assert len(s.labels['type']) > 0
            assert len(s.labels['name']) > 0
            assert len(s.labels['index']) >= 0
            assert len(s.labels['slot']) >= 0

def test_array_hw_val(fb_client):
    array_hw = array_hardware_metrics.ArrayHardwareMetrics(fb_client)
    for m in array_hw.get_metrics():
        for s in m.samples:
            assert (s.value is not None)
