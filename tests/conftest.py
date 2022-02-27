import pytest
from pypureclient import PureError
from pure_fb_openmetrics_exporter.flashblade_client import client


@pytest.fixture()
def fb_client(scope="session"):
    try:
        c = client.FlashbladeClient(
                             target = '10.11.12.1',  # put your fb mgnt ip here
                             api_token='T-xxxxxxxx-yyyy-zzzz-kkkk-jjjjjjjjjjjj', # your API token
                             disable_ssl_warn = True)
    except PureError as pe:
        pytest.fail("Could not connect to flashblade {0}".format(pe))
    yield c
