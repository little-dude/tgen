import tgenpy
import pytest
from . import utils


@pytest.fixture(scope="session")
def create_ports(request):
    def teardown():
        utils.delete_veth_pair('testveth')

    utils.create_veth_pair('testveth')
    request.addfinalizer(teardown)


@pytest.fixture()
def start_tgen(request):

    def teardown():
        utils.kill_tgen()

    utils.start_tgen()
    request.addfinalizer(teardown)


@pytest.fixture()
def controller(start_tgen):
    return tgenpy.Controller('localhost', 1234)
