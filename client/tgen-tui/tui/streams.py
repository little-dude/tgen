import npyscreen as nps
import schemas
import binascii


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


class Form(nps.Form):

    def __init__(self, port, *args, **kwargs):
        self.port = port
        super(Form, self).__init__(*args, **kwargs)
        self.add_handlers({
            "q":               self.previous_view,
            "a":               self.add_stream,
            "d":               self.delete_stream,
            "e":               self.edit_stream,
        })
        self.current_stream = None

    def create(self):
        self.w_streams = self.add(
            StreamsWg, max_width=30, max_height=30, rely=2)
        self.w_help = self.add(HelpWg, rely=32, max_width=30)
        self.w_config = self.add(ConfigWg, max_width=30, rely=2, relx=32)
        self.w_layers = self.add(LayersWg, rely=2, relx=62)
        self.fetch_streams(display=False)

    def fetch_streams(self, display=True):
        self.port.get_streams()
        # self.w_streams.populate()
        self.get_current_stream()
        if display:
            self.display()

    def previous_view(self, *args, **kwargs):
        self.parentApp.switchForm("PORTS")

    def get_current_stream(self):
        wg = self.w_streams.entry_widget
        self.current_stream = None
        try:
            name = wg.values[wg.cursor_line]
        except IndexError:
            return
        for stream in self.port.streams:
            if stream.name == name:
                self.current_stream = stream
                return

    def edit_stream(self, *args, **kwargs):
        self.parentApp.addForm("EDIT_STREAM", ConfigForm, self.current_stream)
        self.parentApp.switchForm("EDIT_STREAM")

    def save_stream(self):
        pass

    def delete_stream(self, *args, **kwargs):
        if self.current_stream is not None:
            self.port.del_stream(self.current_stream.name)
        self.fetch_streams()

    def add_stream(self, *args, **kwargs):
        self.port.new_stream()
        self.fetch_streams()

    def display(self, *args, **kwargs):
        self.get_current_stream()
        super(Form, self).display(*args, **kwargs)


class HelpWg(nps.BoxTitle):

    def __init__(self, *args, **kwargs):
        super(HelpWg, self).__init__(*args, **kwargs)
        self.name = "help"
        self.values = []


class LayersWg(nps.BoxTitle):

    def __init__(self, *args, **kwargs):
        super(LayersWg, self).__init__(*args, **kwargs)
        self.name = "layers"
        self.add_handlers({
            "q":               self.parent.previous_view,
            "a":               self.add_layer,
        })

    def add_layer(self, *args, **kwargs):
        stream = self.parent.current_stream
        stream.add_layer(0, Ethernet2())

    def update(self, *args, **kwargs):
        stream = self.parent.current_stream
        if stream is not None:
            self.values = []
            for layer in stream.layers:
                self.values.append(layer.__class__.__name__)
        else:
            self.values = []
        super(LayersWg, self).update(*args, **kwargs)


class ConfigWg(nps.BoxTitle):

    def __init__(self, *args, **kwargs):
        super(ConfigWg, self).__init__(*args, **kwargs)
        self.name = "config"

    def update(self, *args, **kwargs):
        stream = self.parent.current_stream
        if stream is not None:
            self.values = [
                'name: {}'.format(stream.name),
                'loop: {}'.format(stream.loop),
                'packets/s: {}'.format(stream.packets_per_sec),
            ]
        else:
            self.values = []
        super(ConfigWg, self).update(*args, **kwargs)


class StreamsWg(nps.BoxTitle):

    def __init__(self, *args, **kwargs):
        self.name = "streams"
        super(StreamsWg, self).__init__(*args, **kwargs)

    def update(self, *args, **kwargs):
        try:
            self.populate()
        except AttributeError:
            pass
        super(StreamsWg, self).update(*args, **kwargs)

    def populate(self):
        self.values = []
        for stream in self.parent.port.streams:
            self.values.append(stream.name)

    def when_cursor_moved(self):
        super(StreamsWg, self).when_cursor_moved()
        self.parent.display()


class Stream(object):

    def __init__(self, capnp_stream):
        self.capnp_stream = capnp_stream
        self.get_config()
        self.get_layers()

    def get_config(self):
        res = self.capnp_stream.getConfig().wait()
        self.name = res.config.name
        self.loop = res.config.loop
        self.packets_per_sec = res.config.packetsPerSec

    def set_config(self):
        cfg = schemas.Stream.Config.new_message()
        cfg.name = self.name
        cfg.loop = self.loop
        cfg.packetsPerSec = self.packets_per_sec
        self.capnp_stream.setConfig(cfg).wait()

    def set_layers(self):
        capnp_layers = []
        for layer in self.layers:
            msg = schemas.Protocol.new_message()
            layer.to_capnp(msg)
            capnp_layers.append(msg)
        self.capnp_stream.setLayers(capnp_layers).wait()

    def add_layer(self, position, layer):
        self.layers.insert(position, layer)
        self.set_layers()

    def del_layer(self, layer):
        self.layers.remove(layer)
        self.set_layers()

    def get_layers(self):
        self.layers = []
        res = self.capnp_stream.getLayers().wait()
        for capnp_layer in res.layers:
            self.layers.append(layer_factory(capnp_layer))


class Layer(object):

    def __init__(self, capnp_layer=None):
        if capnp_layer is None:
            capnp_layer = schemas.Protocol.new_message().init(self.capnp_name)
        for field_name, capnp_field_name in self.fields.items():
            capnp_field = getattr(capnp_layer, capnp_field_name)
            setattr(self, field_name, Field(field_name, capnp_field))

    def to_capnp(self, protocol_msg):
        real_msg = protocol_msg.init(self.capnp_name)
        for field_name, capnp_field_name in self.fields.items():
            capnp_field = getattr(real_msg, capnp_field_name)
            getattr(self, field_name).to_capnp(capnp_field)

    @property
    def name(self):
        return self.__class__.__name__


class Field(object):

    def __init__(self, field_name, capnp_field):
        self._value = capnp_field.value
        self._step = capnp_field.step
        self._mask = capnp_field.mask
        self._count = capnp_field.count
        self._mode = capnp_field.mode
        self.name = field_name

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


class FieldPopup(nps.fmPopup.ActionPopup):

    def __init__(self, field, *args, **kwargs):
        self.valid = False
        self.field = field
        self.name = field.name
        super(FieldPopup, self).__init__(*args, **kwargs)

    def create(self):
        self.w_value = self.add(
            nps.TitleText, name="value", value=self.field.value)
        self.w_mask = self.add(
            nps.TitleText, name="mask", value=self.field.mask)
        self.w_step = self.add(
            nps.TitleText, name="step", value=self.field.step)
        self.w_mode = self.add(
            nps.TitleText, name="mode", value=self.field.mode)
        self.w_count = self.add(
            nps.TitleText, name="count", value=str(self.field.count))

    def on_ok(self):
        self.valid = True

    def on_cancel(self):
        self.valid = False


class LayerPopup(nps.fmPopup.ActionPopup):

    protocols = ['Ethernet2', 'IPv4']

    def __init__(self, *args, **kwargs):
        self.valid = False
        super(LayerPopup, self).__init__(*args, **kwargs)

    def create(self):
        self.w_protocol = self.add(
            nps.TitleSelectOne, name="Protocol", values=self.protocols)

    def get_protocol(self):
        return self.protocols[self.w_protocol.value[0]]

    def on_ok(self):
        self.valid = True

    def on_cancel(self):
        self.valid = False


class ConfigForm(nps.ActionFormWithMenus):

    def __init__(self, stream, *args, **kwargs):
        self.stream = stream
        super(ConfigForm, self).__init__(*args, **kwargs)
        self.make_main_menu()

    def configure_field(self, field):
        form = FieldPopup(field)
        # form.preserve_selected_widget = True
        form.edit()
        if form.valid is True:
            field.value = form.w_value.value
            field.step = form.w_step.value
            field.mask = form.w_mask.value
            field.count = form.w_count.value
            field.mode = form.w_mode.value
        self.stream.set_layers()

    def add_layer(self):
        form = LayerPopup()
        form.preserve_selected_widget = True
        form.edit()
        if form.valid is False:
            return
        protocol = form.get_protocol()
        if protocol == 'Ethernet2':
            layer = Ethernet2()
        elif protocol == 'IPv4':
            layer = IPv4()
        self.stream.add_layer(0, layer)

    def make_main_menu(self):
        self.main_menu = self.new_menu(
            name='layers configuration', shortcut='^X')
        self.make_layers_menu()
        self.main_menu.addItem('Add layer', onSelect=self.add_layer)
        self.make_delete_menu()

    def make_layers_menu(self):
        for layer in self.stream.layers:
            menu = self.main_menu.addNewSubmenu(name=layer.name)
            for field_name in layer.fields.keys():
                menu.addItem(field_name,
                             onSelect=self.configure_field,
                             arguments=(getattr(layer, field_name),))

    def make_delete_menu(self):
        menu = self.main_menu.addNewSubmenu(name='Delete layer')
        for layer in self.stream.layers:
            menu.addItem(layer.name,
                         onSelect=self.stream.del_layer,
                         arguments=(layer,))

    def create(self):
        self.w_name = self.add(
            nps.TitleText, name='name', value=self.stream.name)
        self.w_loop = self.add(
            nps.Checkbox, name='loop', value=self.stream.loop)
        self.w_rate = self.add(
            nps.TitleText, name='packets/s',
            value=str(self.stream.packets_per_sec))

    def on_ok(self):
        self.stream.name = self.w_name.value
        self.stream.loop = self.w_loop.value
        self.stream.packets_per_sec = int(self.w_rate.value)
        self.stream.set_config()
        self.parentApp.switchForm("STREAMS")


def monkey_patch_npyscreen_menu():
    """
    we want to be able to generate menu contents dynamically but the default
    menu class does not allow it, and it's difficult to make the forms use a
    custom menu class, so we monkey patch it here.
    """
    def setItems(self, items):
        self._menuList = []
        self.addItemsFromList(items)

    setattr(nps.muNewMenu.NewMenu, "setItems", setItems)

monkey_patch_npyscreen_menu()
