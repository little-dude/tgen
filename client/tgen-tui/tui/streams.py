import npyscreen as nps
import schemas

class Form(nps.Form):

    def __init__(self, port, *args, **kwargs):
        self.port = port
        super(Form, self).__init__(*args, **kwargs)
        self.add_handlers({
            "q":               self.previous_view,
            "a":               self.add_stream,
            "d":               self.delete_stream,
            "c":               self.edit_stream
        })
        self.current_stream = None

    def create(self):
        self.w_streams = self.add(StreamsWg, max_width=30, max_height=30, rely=2)
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


class ConfigForm(nps.ActionPopup):
    def __init__(self, stream, *args, **kwargs):
        self.stream = stream
        super(ConfigForm, self).__init__(*args, **kwargs)

    def create(self):
        self.w_name = self.add(
            nps.TitleText, name='name', value=self.stream.name)
        self.w_loop = self.add(
            nps.Checkbox, name='loop', value=self.stream.loop)
        self.w_rate = self.add(
            nps.TitleText, name='packets/s', value=self.stream.packets_per_sec)

    def on_ok(self):
        self.stream.name = self.w_name.value
        self.stream.loop = self.w_loop.value
        self.stream.packets_per_sec = int(self.w_rate.value)
        self.stream.set_config()
        self.parentApp.switchForm("STREAMS")
