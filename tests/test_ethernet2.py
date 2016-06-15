# -*- coding: utf-8 -*-

import pytest
import tgenpy
import logging

logging.basicConfig(level=logging.INFO)


@pytest.mark.usefixtures('start_tgen')
def test_connect():
    tgenpy.Controller('localhost', 1234)


@pytest.mark.usefixtures('create_ports')
def test_fetch_ports(controller):
    ports = controller.fetch_ports()
    assert len(ports) >= 2
    veth_found = 0
    for port in ports:
        port.get_config()
        if port.name in ['testveth0', 'testveth1']:
            veth_found += 1
    assert veth_found == 2
