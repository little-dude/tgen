from __future__ import unicode_literals
from builtins import object
from . import protocols


class Stream(object):

    def __init__(self, capnp_stream=None):
        self.count = 1
        self.packets_per_sec = 1
        self.layers = []
        self.id = None
        if capnp_stream:
            self.from_capnp(capnp_stream)

    def from_capnp(self, capnp_stream):
        self.count = capnp_stream.count
        self.packets_per_sec = capnp_stream.packetsPerSec
        self.id = capnp_stream.id
        self.layers = []
        for capnp_layer in capnp_stream.layers:
            self.layers.append(protocols.new_layer(capnp_layer))

    def to_capnp(self, capnp_stream):
        capnp_stream.count = self.count
        capnp_stream.packetsPerSec = self.packets_per_sec
        layers = capnp_stream.init('layers', len(self.layers))
        for idx, layer in enumerate(self.layers):
            layer.to_capnp(layers[idx])


def new_layer(capnp_layer):
    which = capnp_layer.which()
    if which == 'ethernet2':
        return protocols.Ethernet2(capnp_layer.ethernet2)
    elif which == 'ipv4':
        return protocols.IPv4(capnp_layer.ipv4)
    else:
        raise Exception('unknown {}'.format(which))
