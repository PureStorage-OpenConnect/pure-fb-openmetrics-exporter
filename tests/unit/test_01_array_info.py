from pure_fb_openmetrics_exporter.flashblade_collector.flashblade_metrics import array_info_metrics


def test_array_info(fb_client):
    """
    GIVEN a FlashArray
    WHEN information metric is requested
    THEN check the
    """
    array_info = array_info_metrics.ArrayInfoMetrics(fb_client)
    m =  next(array_info.get_metrics())
    for s in m.samples:
        assert s.name == 'purefb_info'
        assert len(s.labels['array_name']) > 0
        assert len(s.labels['system_id']) > 0
        assert len(s.labels['version']) > 0
        assert s.value == 1
