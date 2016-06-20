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

    def wait_send(self, timeout=0):
        res = self._capnp_port.waitSend(timeout).wait()
        return (res.done, res.error)

    def start_capture(self, savePath, snapshot_length=9999, promiscuous=True,
                      packet_count=0, timeout=0):
        self._capnp_port.startCapture(savePath, snapshot_length, timeout,
                                      packet_count, promiscuous).wait()

    def wait_capture(self, timeout=0):
        res = self._capnp_port.waitCapture(timeout).wait()
        return (res.done, res.error)

    def stop_capture(self):
        self._capnp_port.stopCapture().wait()
