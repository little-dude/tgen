import os
import binascii
import schemas
import sys
import capnp

capnp.remove_import_hook()
sys.modules['schemas'] = capnp.load(
    os.path.join(
        os.path.dirname(os.path.realpath(__file__)), 'schemas.capnp'))


class Controller():

    def __init__(self, hostname, port):
        address = '{}:{}'.format(hostname, port)
        client = capnp.TwoPartyClient(address)
        self._controller = client.bootstrap().cast_as(schemas.Controller)
        self.ports = []
        self.streams = []

    def fetch_ports(self):
        res = self._controller.getPorts().wait()
        self.ports = []
        for port in res.ports:
            self.ports.append(Port(port))

    def fetch_streams(self):
        res = self._controller.listStreams().wait()
        self.streams = []
        for stream_id in res.ids:
            self.streams.append(self.fetch_stream(stream_id))

    def fetch_stream(self, stream_id):
        res = self._controller.fetchStream(stream_id).wait()
        return Stream(capnp_stream=res.stream)

    def create_stream(self, stream):
        capnp_stream = schemas.Stream.new_message()
        stream.to_capnp(capnp_stream)
        stream.id = self._controller.createStream(capnp_stream).wait().id
        self.streams.append(stream)
        return stream

    def delete_stream(self, stream_id):
        for idx, stream in enumerate(self.streams):
            if stream.id == stream_id:
                self.streams.pop(idx)
                self._controller.deleteStream(stream_id).wait()


class TgenError(Exception):
    pass


class Port():

    def __init__(self, capnp_port):
        self._capnp_port = capnp_port
        self.streams = []

    def get_config(self):
        res = self._capnp_port.getConfig().wait()
        self.name = res.config.name


class Stream(object):

    def __init__(self, capnp_stream=None):
        if capnp_stream is not None:
            self.from_capnp(capnp_stream)
        else:
            self.count = 1
            self.packets_per_sec = 1
            self.layers = []

    def from_capnp(self, capnp_stream):
        self._capnp_stream = capnp_stream
        self.count = capnp_stream.count
        self.packets_per_sec = capnp_stream.packetsPerSec
        self.id = capnp_stream.id
        self.layers = []
        for capnp_layer in capnp_stream.layers:
            self.layers.append(Layer(capnp_layer=capnp_layer))

    def to_capnp(self, capnp_stream):
        capnp_stream.count = self.count
        capnp_stream.packetsPerSec = self.packets_per_sec
        layers = capnp_stream.init('layers', len(self.layers))
        for idx, layer in enumerate(self.layers):
            layer.to_capnp(layers[idx])
        # self._capnp_stream.setConfig(cfg).wait()

    # def set_layers(self):
    #     capnp_layers = []
    #     for layer in self.layers:
    #         msg = schemas.Protocol.new_message()
    #         layer.to_capnp(msg)
    #         capnp_layers.append(msg)
    #     self._capnp_stream.setLayers(capnp_layers).wait()

    # def add_layer(self, position, layer):
    #     self.layers.insert(position, layer)
    #     self.set_layers()

    # def del_layer(self, layer):
    #     self.layers.remove(layer)
    #     self.set_layers()

    # def get_layers(self):
    #     self.layers = []
    #     res = self._capnp_stream.getLayers().wait()
    #     for capnp_layer in res.layers:
    #         self.layers.append(layer_factory(capnp_layer))


class Layer(object):

    def __init__(self, capnp_layer=None):
        if capnp_layer is None:
            # create a capnp_layer, since it has default values that we can use
            # to initialize the object
            capnp_layer = schemas.Protocol.new_message().init(self.capnp_name)
        self.from_capnp(capnp_layer)

    def from_capnp(self, capnp_layer):
        for field_name, capnp_field_name in self.fields.items():
            capnp_field = getattr(capnp_layer, capnp_field_name)
            setattr(self, field_name, Field(capnp_field))

    def to_capnp(self, capnp_layer):
        real_msg = capnp_layer.init(self.capnp_name)
        for field_name, capnp_field_name in self.fields.items():
            capnp_field = getattr(real_msg, capnp_field_name)
            getattr(self, field_name).to_capnp(capnp_field)

    @property
    def name(self):
        return self.__class__.__name__


class Field(object):

    def __init__(self, capnp_field=None):
        if capnp_field is None:
            self.value = 0
            self.step = 0
            self.mask = 0
            self.count = 0
            self.mode = 'fixed'
        else:
            self.from_capnp(capnp_field)

    def from_capnp(self, capnp_field):
        self._value = capnp_field.value
        self._step = capnp_field.step
        self._mask = capnp_field.mask
        self._count = capnp_field.count
        self._mode = capnp_field.mode

    def to_capnp(self, capnp_field):
        capnp_field.value = self._value
        capnp_field.step = self._step
        capnp_field.mask = self._mask
        capnp_field.count = self._count
        capnp_field.mode = self._mode

    @property
    def value(self):
        return '0x{}'.format(self.data2str(self._value))

    @value.setter
    def value(self, data):
        self._value = self.parse_data_input(data)

    @property
    def step(self):
        return '0x{}'.format(self.data2str(self._step))

    @step.setter
    def step(self, data):
        self._step = self.parse_data_input(data)

    @property
    def mask(self):
        return '0x{}'.format(self.data2str(self._mask))

    @mask.setter
    def mask(self, data):
        self._mask = self.parse_data_input(data)

    @staticmethod
    def parse_data_input(data):
        if isinstance(data, str):
            data = str2int(data)
        elif isinstance(data, (tuple, list)):
            data = iter2int(data)
        if not isinstance(data, int):
            raise ValueError('Cannot parse {}'.format(data))
        data = hex(data)[2:]
        if len(data) % 2 == 1:
            data = '0{}'.format(data)
        try:
            return binascii.unhexlify(data)
        except:
            raise Exception(data)

    @property
    def mode(self):
        for key, value in MODES_MAPPING.items():
            if self._mode == value:
                return key
        raise ValueError('Unknown mode {}'.format(self._mode))

    @mode.setter
    def mode(self, data):
        if data not in MODES_MAPPING:
            self._mode = FieldMode.FIXED
        self._mode = MODES_MAPPING[data]

    @property
    def count(self):
        return self._count

    @count.setter
    def count(self, data):
        if isinstance(data, str):
            data = str2int(data)
        if not isinstance(data, int):
            raise ValueError('Expected "int" got "{}"'.format(type(data)))
        if data < 0 or data > 18446744073709551615:
            raise ValueError('"count" must be between 0 and 2^64-1')
        self._count = data

    @staticmethod
    def data2str(data):
        string = binascii.hexlify(data).decode()
        if string == '':
            string = '00'
        return string


class Ethernet2(Layer):

    fields = {
        'source': 'source',
        'destination': 'destination',
        'ethernet_type': 'ethernetType'
    }
    capnp_name = 'ethernet2'


class IPv4(Layer):

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


def layer_factory(capnp_layer):
    which = capnp_layer.which()
    if which == 'ethernet2':
        return Ethernet2(capnp_layer.ethernet2)
    elif which == 'ipv4':
        return IPv4(capnp_layer.ipv4)
    else:
        raise Exception('unknown {}'.format(which))


class FieldMode(object):

    FIXED = 0
    INCREMENT = 1
    DECREMENT = 2
    RANDOMIZE = 3
    AUTO = 4

MODES_MAPPING = {
    'fixed': FieldMode.FIXED,
    'auto': FieldMode.AUTO,
    'increment': FieldMode.INCREMENT,
    'decrement': FieldMode.DECREMENT,
    'randomize': FieldMode.RANDOMIZE,
}


def str2int(string):
    if string.startswith('0x'):
        return int(string, 16)
    elif string.startswith('0b'):
        return int(string, 2)
    else:
        return int(string)


def iter2int(iterable):
    data = []
    for item in iterable:
        if isinstance(item, int):
            data.append(str(item))
        elif isinstance(item, str):
            data.append(str2int(item))
        else:
            raise ValueError(
                'Cannot parse {}'.format(iterable))
    return ''.data
