import pytest
from pypureclient import PureError
from pure_fb_prometheus_exporter.flashblade_client import client


@pytest.fixture()
def fb_client(scope="session"):
    try:
        c = client.FlashbladeClient(
                             target = '10.225.112.69',
                             api_token='T-2b74f9eb-a35f-40d9-a6a6-33c13775a53c',
                             disable_ssl_warn = True)
    except PureError as pe:
        pytest.fail("Could not connect to flashblade {0}".format(pe))
    yield c
