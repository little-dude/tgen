import capnp


def make_schemas_module():
    """
    small hack to make the schemas available in ``sys.modules``
    """
    import os
    import sys

    capnp.remove_import_hook()
    sys.modules['schemas'] = capnp.load(
        os.path.join(
            os.path.dirname(os.path.realpath(__file__)), 'schemas.capnp'))

make_schemas_module()
import schemas
from . import protocols


class Controller():
    """
    Represent a remote tgen controller
    """

    def __init__(self, hostname, port):
        address = '{}:{}'.format(hostname, port)
        client = capnp.TwoPartyClient(address)
        self._controller = client.bootstrap().cast_as(schemas.Controller)
        self.ports = []

    def fetch_ports(self):
        """
        Return the ports available on the tgen controller
        """
        res = self._controller.getPorts().wait()
        self.ports = []
        for port in res.ports:
            self.ports.append(Port(self._controller, port))

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


class Port():

    def __init__(self, capnp_controller, capnp_port):
        self._controller = capnp_controller
        self._capnp_port = capnp_port

    def get_config(self):
        res = self._capnp_port.getConfig().wait()
        self.name = res.config.name

    def start_send(self, stream_ids):
        self.promise = self._capnp_port.startSend(stream_ids)

    def wait(self):
        self.promise.wait()


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
