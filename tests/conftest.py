import pytest
from pypureclient import PureError
from pure_fb_openmetrics_exporter.flashblade_client import client
from pure_fb_openmetrics_exporter import pure_fb_exporter
from sanic_testing.testing import SanicTestClient


@pytest.fixture()
def fb_client(scope="session"):
    try:
        c = client.FlashbladeClient(
                             target = '10.225.112.69',  # put your fb mgnt ip here
                             api_token='T-2b74f9eb-a35f-40d9-a6a6-33c13775a53c', # your API token
                             disable_ssl_warn = True)
    except PureError as pe:
        pytest.fail("Could not connect to flashblade {0}".format(pe))
    yield c

@pytest.fixture()
def api_token():
    return '' 

@pytest.fixture()
def endpoint():
    return '10.225.112.69' 

@pytest.fixture(scope="session")
def app_client():
    app = pure_fb_exporter.app
    app.ctx.disable_cert_warn = True
    client = SanicTestClient(app)
    return client
