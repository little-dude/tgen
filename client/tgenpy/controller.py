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
        # HACK: There seems to be the weirdest bug here.  If fetchStream is
        # called directly in the loop that iterates over the ids, the ids take
        # random values. To avoid this we copy all the ids in a separate list
        # and iterate over this new list.
        #
        # TODO: reproduce and raise an issue on github
        stream_ids = []
        for stream_id in self._controller.listStreams().wait().ids:
            stream_ids.append(stream_id)
        for stream_id in stream_ids:
            res = self._controller.fetchStream(stream_id).wait()
            stream = Stream(capnp_stream=res.stream)
            streams.append(stream)
        return streams

    def get_port(self, name):
        ports = self.fetch_ports()
        for port in ports:
            port.get_config()
            if port.name == name:
                return port
        raise ValueError('Port {} not found'.format(name))

    def save_stream(self, stream):
        capnp_stream = schemas.Stream.new_message()
        stream.to_capnp(capnp_stream)
        stream.id = self._controller.saveStream(capnp_stream).wait().id

    def delete_stream(self, stream_id):
        """
        Delete a stream by ID.
        """
        self._controller.deleteStream(stream_id).wait()
