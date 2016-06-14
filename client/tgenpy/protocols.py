from __future__ import unicode_literals
from . import layer


class Ethernet2(layer.Layer):

    fields = {
        'source': 'source',
        'destination': 'destination',
        'ethernet_type': 'ethernetType'
    }
    capnp_name = 'ethernet2'


class IPv4(layer.Layer):

    fields = {
        'version': 'version',
        'ihl': 'ihl',
        'tos': 'tos',
        'length': 'length',
        'id': 'id',
        'flags': 'flags',
        'frag_offset': 'fragOffset',
        'ttl': 'ttl',
        'protocol': 'protocol',
        'checksum': 'checksum',
        'source': 'source',
        'destination': 'destination',
        'options': 'options',
        'padding': 'padding',
    }
    capnp_name = 'ipv4'
