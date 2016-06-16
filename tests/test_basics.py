# -*- coding: utf-8 -*-

import pytest
import tgenpy
from tgenpy import protocols
import logging

# logging.basicConfig(level=logging.INFO)
logging.basicConfig(level=logging.DEBUG)


@pytest.mark.usefixtures('start_tgen')
def test_connect():
    """
    Connect to the controller on the default port
    """
    tgenpy.Controller('localhost', 1234)


@pytest.mark.usefixtures('create_ports')
def test_fetch_ports(controller):
    """
    Fetch the ports and their configuration
    """
    ports = controller.fetch_ports()
    assert len(ports) >= 2
    veth_found = 0
    for port in ports:
        port.get_config()
        if port.name in ['testveth0', 'testveth1']:
            veth_found += 1
    assert veth_found == 2


def test_single_stream(controller):
    """
    CRUD operations on a simple stream
    """
    # create a stream
    stream = tgenpy.Stream()
    assert stream.count == 1
    assert stream.packets_per_sec == 1
    assert stream.layers == []
    stream.layers = [tgenpy.Ethernet2()]
    controller.save_stream(stream)

    # fetch stream
    streams = controller.fetch_streams()
    assert len(streams) == 1
    stream = streams[0]
    assert stream.id == 1
    assert stream.count == 1
    assert stream.packets_per_sec == 1
    assert len(stream.layers) == 1
    assert isinstance(stream.layers[0], protocols.Ethernet2)

    # update stream
    stream.count = 2
    stream.packets_per_sec = 100
    stream.layers.append(tgenpy.IPv4())
    controller.save_stream(stream)

    # fetch stream
    streams = controller.fetch_streams()
    assert len(streams) == 1
    stream = streams[0]
    # it's important to check the ID didn't change to make sure the stream has
    # been updated, instead of a new one being created
    assert stream.id == 1
    assert stream.count == 2
    assert stream.packets_per_sec == 100
    assert len(stream.layers) == 2
    assert isinstance(stream.layers[0], protocols.Ethernet2)
    assert isinstance(stream.layers[1], protocols.IPv4)

    # delete stream
    controller.delete_stream(1)
    streams = controller.fetch_streams()
    assert len(streams) == 0


def test_multiple_streams(controller):
    s1 = tgenpy.Stream(layers=[protocols.Ethernet2()])
    s2 = tgenpy.Stream(layers=[protocols.Ethernet2(), protocols.IPv4()])
    controller.save_stream(s1)
    controller.save_stream(s2)
    streams = controller.fetch_streams()
    assert len(streams) == 2
    if streams[0].id == 1:
        assert streams[1].id == 2
    else:
        assert streams[0].id == 2
        assert streams[1].id == 1

    for i in range(0, 8):
        controller.save_stream(tgenpy.Stream(layers=[protocols.Ethernet2()]))
    ids = controller._controller.listStreams().wait().ids
    assert 0 not in ids
    assert len(controller.fetch_streams()) == 10
