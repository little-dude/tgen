from __future__ import unicode_literals
from builtins import hex
from builtins import str
from builtins import object
import schemas
import binascii


class Layer(object):

    def __init__(self, capnp_layer=None):
        if capnp_layer is None:
            # create a capnp_layer, since it has default values that we can use
            # to initialize the object
            capnp_layer = schemas.Protocol.new_message().init(self.capnp_name)
        self.from_capnp(capnp_layer)

    def from_capnp(self, capnp_layer):
        for field_name, capnp_field_name in list(self.fields.items()):
            capnp_field = getattr(capnp_layer, capnp_field_name)
            setattr(self, field_name, Field(capnp_field))

    def to_capnp(self, capnp_layer):
        real_msg = capnp_layer.init(self.capnp_name)
        for field_name, capnp_field_name in list(self.fields.items()):
            capnp_field = getattr(real_msg, capnp_field_name)
            getattr(self, field_name).to_capnp(capnp_field)

    @property
    def name(self):
        return self.__class__.__name__


class Field(object):

    def __init__(self, capnp_field=None):
        if capnp_field is None:
            # create a capnp_field, since it has default values that we can use
            # to initialize the object
            capnp_field = schemas.Field.new_message()
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
        for key, value in list(MODES_MAPPING.items()):
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
