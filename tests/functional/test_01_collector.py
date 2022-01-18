from pure_fb_prometheus_exporter.flashblade_collector import flashblade_collector

def test_collector_array(fb_client):
    collector = flashblade_collector.FlashbladeCollector(fb_client, request='array')

    for s in collector.collect():
        print(type(s))
