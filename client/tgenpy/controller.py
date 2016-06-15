# -*- coding: utf-8

from __future__ import unicode_literals
from builtins import object
import capnp
import schemas
from .port import Port
from .stream import Stream
from . import utils


class Controller(object):
    """
    Represent a remote tgen controller
    """

    def __init__(self, hostname, port):
        address = '{}:{}'.format(hostname, port)
        client = capnp.TwoPartyClient(utils.ensure_native_str(address))
        self._controller = client.bootstrap().cast_as(schemas.Controller)

    def fetch_ports(self):
        """
        Return the ports available on the tgen controller
        """
        res = self._controller.getPorts().wait()
        ports = []
        for port in res.ports:
            ports.append(Port(self._controller, port))
        return ports

    def fetch_streams(self):
        """
        Return the streams configured on the tgen controller
        """
        streams = []
        ids = self._controller.listStreams().wait().ids
        for stream_id in ids:
            res = self._controller.fetchStream(stream_id).wait()
            stream = Stream(capnp_stream=res.stream)
            streams.append(stream)
        return streams

    def save_stream(self, stream):
        capnp_stream = schemas.Stream.new_message()
        stream.to_capnp(capnp_stream)
        stream.id = self._controller.saveStream(capnp_stream).wait().id

    def delete_stream(self, stream_id):
        """
        Delete a stream by ID.
        """
        self._controller.deleteStream(stream_id).wait()
