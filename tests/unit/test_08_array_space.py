from pure_fb_prometheus_exporter.flashblade_collector.flashblade_metrics.array_space_metrics import ArraySpaceMetrics


def test_array_space_name(fb_client):
    array_space = ArraySpaceMetrics(fb_client)
    for m in array_space.get_metrics():
        for s in m.samples:
            assert s.name in ['purefb_array_space_data_reduction_ratio', 
                              'purefb_array_space_bytes',
                              'purefb_array_space_parity']

def test_array_space_labels(fb_client):
    array_space = ArraySpaceMetrics(fb_client)
    for m in array_space.get_metrics():
        for s in m.samples:
            assert s.labels['type'] in ['array', 'file-system', 'object-store']
            if s.name == 'purefb_array_space_bytes':
                assert s.labels['space'] in ['snapshots',
                                             'total_physical',
                                             'unique',
                                             'virtual',
                                             'capacity']
            elif s.name in ['purefb_array_space_data_reduction_ratio',
                            'purefb_array_space_parity']:
                assert list(s.labels.keys()) == ['type']
         
def test_array_space_val(fb_client):
    array_space = ArraySpaceMetrics(fb_client)
    for m in array_space.get_metrics():
        for s in m.samples:
            assert s.value is not None
            assert s.value >= 0
