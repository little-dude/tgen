# -*- coding: utf-8 -*-

from __future__ import unicode_literals
import pytest
import tgenpy
from tgenpy import protocols
import logging
from . import utils
import netaddr

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


def test_long_field(controller):

    def check_defaults(field):
        assert field.count == 1
        assert field._count == 1

        assert field.mask == '0xffffffffffff'
        assert field._mask == b'\xff\xff\xff\xff\xff\xff'

        assert field.value == '0x000000000000'
        assert field._value == b'\x00\x00\x00\x00\x00\x00'

        assert field.mode == 'fixed'
        assert field._mode == 0

        assert field.step == '0x00'
        assert field._step == b''

    def check_new(field):
        assert field.count == 100
        assert field._count == 100

        assert field.mask == '0xfedcba987654'
        assert field._mask == b'\xfe\xdc\xba\x98\x76\x54'

        assert field.value == '0x0123456789ab'
        assert field._value == b'\x01\x23\x45\x67\x89\xab'

        assert field.step == '0x2a'
        assert field._step == b'\x2a'

    # create layer and check default values
    layer = protocols.Ethernet2()
    check_defaults(layer.source)
    stream = tgenpy.Stream(layers=[layer])

    # save layer and check default values
    controller.save_stream(stream)
    layer = stream.layers[0]
    check_defaults(layer.source)

    # set custom values and check
    field = layer.source
    field.count = 100
    field.mask = '0xfedcba987654'
    field.value = '0x0123456789ab'
    field.mode = 'increment'
    field.step = 42
    check_new(field)

    # save layer and check custom values
    controller.save_stream(stream)
    layer = stream.layers[0]
    check_new(layer.source)


@pytest.mark.usefixtures('create_ports')
def test_ethernet2(controller):

    eth = protocols.Ethernet2()

    eth.source.mode = 'increment'
    eth.source.step = 1
    eth.source.mask = 0x0000000000ff
    eth.source.count = 256

    eth.destination.value = 0xffffffffffff
    eth.destination.mode = 'decrement'
    eth.destination.step = 0x010000000000
    eth.destination.mask = 0xff0000000000
    eth.destination.count = 256

    # add dummy ip layer, otherwise pyshark refuses to load the frames
    ip = protocols.IPv4()

    stream = tgenpy.Stream(layers=[eth, ip])
    stream.count = 256
    controller.save_stream(stream)

    tx = controller.get_port('testveth0')
    rx = controller.get_port('testveth1')
    with utils.Capture(rx, count=256) as capture:
        tx.start_send([stream.id])
        packets = capture.get_packets(timeout=5000)

    for idx, packet in enumerate(packets):
        ethernet = packet.eth
        assert netaddr.EUI(ethernet.src).value == idx
        assert netaddr.EUI(ethernet.dst).value == 0xffffffffffff - (idx << (5*8))
        assert int(ethernet.type, 16) == 0x800
