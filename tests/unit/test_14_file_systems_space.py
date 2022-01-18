from pure_fb_prometheus_exporter.flashblade_collector.flashblade_metrics.file_systems_space_metrics import FileSystemsSpaceMetrics


def test_file_systems_space_name(fb_client):
    file_systems_space = FileSystemsSpaceMetrics(fb_client)
    for m in file_systems_space.get_metrics():
        for s in m.samples:
            assert s.name in ['purefb_file_systems_space_data_reduction_ratio', 
                              'purefb_file_systems_space_bytes']

def test_file_systems_space_labels(fb_client):
    file_systems_space = FileSystemsSpaceMetrics(fb_client)
    for m in file_systems_space.get_metrics():
        for s in m.samples:
            assert 'nfs' in list(s.labels.keys()) 
            assert 'smb' in list(s.labels.keys()) 
            if s.name == 'purefb_file_systems_space_bytes':
                assert s.labels['space'] in ['provisioned',
                                             'snapshots',
                                             'total_physical',
                                             'unique',
                                             'virtual']
            if s.name == 'purefb_file_systems_space_data_reduction_ratio':
                for k in ['provisioned',
                          'snapshots',
                          'total_physical', 
                          'unique', 
                          'virtual']:
                    assert k not in list(s.labels.keys()) 
            assert len(s.labels['name']) > 0
         
def test_file_systems_space_val(fb_client):
    file_systems_space = FileSystemsSpaceMetrics(fb_client)
    for m in file_systems_space.get_metrics():
        for s in m.samples:
            assert s.value is not None
            assert s.value >= 0
