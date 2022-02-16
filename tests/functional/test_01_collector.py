from pure_fb_openmetrics_exporter.flashblade_collector import collector

def test_collector_array(fb_client):
    coll = collector.FlashbladeCollector(fb_client, request='array')

    for s in coll.collect():
        print(type(s))
