from __future__ import unicode_literals
from builtins import object


class Port(object):

    def __init__(self, capnp_controller, capnp_port):
        self._controller = capnp_controller
        self._capnp_port = capnp_port
        self.name = None

    def get_config(self):
        res = self._capnp_port.getConfig().wait()
        self.name = res.config.name

    def start_send(self, stream_ids):
        self._capnp_port.startSend(stream_ids).wait()

    def wait_send(self, timeout=1):
        res = self._capnp_port.waitSend(timeout).wait()
        return (res.done, res.error)
