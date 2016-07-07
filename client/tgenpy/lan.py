from __future__ import unicode_literals
from builtins import object
from . import utils
import schemas


class LAN(object):

    def __init__(self, capnp_port, capnp_lan):
        self._capnp_port = capnp_port
        self._capnp_lan = capnp_lan
        self.network = self._get_config().cidr

    def _get_config(self):
        return self._capnp_lan.getConfig().wait().config

    def get_devices(self):
        devices = []
        for dev in self._get_config().devices:
            devices.append(Device(capnp_device=dev))
        return devices

    def add_device(self, mac, ip):
        devices = self.get_devices()
        devices.append(Device(mac=mac, ip=ip))
        config = schemas.Lan.Config.new_message()
        config.init('devices', len(devices))
        for i, dev in enumerate(devices):
            dev.to_capnp(config.devices[i])
        self._capnp_lan.setConfig(config).wait()

    def start(self):
        self._capnp_lan.start().wait()

    def stop(self):
        self._capnp_lan.stop().wait()


class Device(object):

    def __init__(self, mac=None, ip=None, capnp_device=None):
        if capnp_device is None and (mac is None or ip is None):
            raise ValueError(
                'Provide either a capnp object or valid IP and MAC addresses')
        if capnp_device is None:
            self.mac = mac
            self.ip = ip
        else:
            self.from_capnp(capnp_device)

    @property
    def mac(self):
        return utils.bytes2mac(self._mac)

    @mac.setter
    def mac(self, value):
        self._mac = utils.mac2bytes(value)

    @property
    def ip(self):
        return utils.bytes2ip(self._ip)

    @ip.setter
    def ip(self, value):
        self._ip = utils.ip2bytes(value)

    def from_capnp(self, capnp_device):
        self._ip = capnp_device.ip
        self._mac = capnp_device.mac

    def to_capnp(self, capnp_device):
        capnp_device.ip = self._ip
        capnp_device.mac = self._mac
