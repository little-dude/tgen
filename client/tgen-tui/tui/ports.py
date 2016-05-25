import curses
import npyscreen as nps
from . import streams


class Form(nps.Form):

    def __init__(self, address, controller, *args, **kwargs):
        self.controller = controller
        self.address = address
        super(Form, self).__init__(*args, **kwargs)
        self.add_handlers({
            curses.ascii.NL:   self.edit_port_config,
            curses.ascii.CR:   self.edit_port_config,
            "q":               self.quit,
            "s":               self.edit_streams,
        })

    def create(self):
        self.w_ports = self.add(PortsWg, max_width=30, max_height=30)
        self.w_help = self.add(HelpWg, rely=32, max_width=30)
        self.w_config = self.add(ConfigWg, max_width=30, max_height=30, rely=2, relx=32)
        self.w_stats = self.add(StatsWg, max_width=30, rely=32, relx=32)
        self.w_streams = self.add(StreamsWg, rely=2, relx=62)
        self.update(fetch_ports=True)

    def fetch_ports(self):
        ports_res = self.controller.getPorts().wait()
        self.ports = []
        for port in ports_res.ports:
            self.ports.append(Port(port))

    def edit_port_config(self, *args, **kwargs):
        pass

    def edit_streams(self, *args, **kwargs):
        self.parentApp.addForm("STREAMS", streams.Form, self.current_port)
        self.parentApp.switchForm("STREAMS")

    def quit(self, *args, **kwargs):
        self.parentApp.switchForm(None)

    def update(self, fetch_ports=False):
        if fetch_ports:
            self.fetch_ports()
            self.w_ports.update()
        self.get_current_port()
        self.display()

    def get_current_port(self):
        wg = self.w_ports.entry_widget
        self.current_port = None
        try:
            name = wg.values[wg.cursor_line]
        except IndexError:
            return
        for port in self.ports:
            if port.name == name:
                self.current_port = port
                return


class StreamsWg(nps.BoxTitle):

    """
    Streams box widget. It shows the streams for the current port.
    """

    def __init__(self, *args, **kwargs):
        super(StreamsWg, self).__init__(*args, **kwargs)
        self.name = "streams"

    def update(self, *args, **kwargs):
        self.values = []
        super(StreamsWg, self).update(*args, **kwargs)


class HelpWg(nps.BoxTitle):

    """
    Help box widget. It shows the available keyboard shortcuts.
    """

    def __init__(self, *args, **kwargs):
        super(HelpWg, self).__init__(*args, **kwargs)
        self.name = "help"
        self.values = [
            'R: start (R)eceiving',
            'r: stop  (r)eceiving',
            'S: start (S)ending',
            's: stop  (s)ending',
            'C: (C)lear stats',
            'F5: refresh',
            'e: Edit port configuration',
            'undecided: Edit streams',
            '^q: Quit']


class StatsWg(nps.BoxTitle):

    """
    Stats box widget. It holds the current port statistics.
    """

    def __init__(self, *args, **kwargs):
        super(StatsWg, self).__init__(*args, **kwargs)
        self.name = "stats"

    def update(self, *args, **kwargs):
        self.values = [
            'tx bytes: unknown',
            'tx packets: unknown',
            'rx bytes: unknown',
            'rx packets: unknown']
        super(StatsWg, self).update(*args, **kwargs)


class ConfigWg(nps.BoxTitle):

    """
    Port config widget. It holds the current port config.
    """

    def __init__(self, *args, **kwargs):
        super(ConfigWg, self).__init__(*args, **kwargs)
        self.name = "config"

    def update(self, *args, **kwargs):
        self.values = [
            'name: {}'.format(self.parent.current_port.name),
            'promiscuous: unknown']
        super(ConfigWg, self).update(*args, **kwargs)


class PortsWg(nps.BoxTitle):

    """
    Ports widget. It list the available ports. The other widgets are updated
    when the cursos moves within this widget.
    """

    def __init__(self, *args, **kwargs):
        super(PortsWg, self).__init__(*args, **kwargs)
        self.name = "ports"

    def update(self, *args, **kwargs):
        self.values = []
        for port in self.parent.ports:
            self.values.append(port.name)
        super(PortsWg, self).update(*args, **kwargs)

    def when_cursor_moved(self):
        super(PortsWg, self).when_cursor_moved()
        self.parent.update()


class Port(object):

    def __init__(self, capnp_port):
        self.capnp_port = capnp_port
        self.get_config()
        self.get_streams()

    def get_config(self):
        res = self.capnp_port.getConfig().wait()
        self.name = res.config.name

    def get_streams(self):
        self.streams = []
        res = self.capnp_port.getStreams().wait()
        for stream in res.streams:
            self.streams.append(streams.Stream(stream))

    def new_stream(self):
        res = self.capnp_port.newStream().wait()
        self.streams.append(streams.Stream(res.stream))

    def del_stream(self, name):
        target = None
        for idx, stream in enumerate(self.streams):
            if stream.name == name:
                target = self.streams.pop(idx)
                break
        if target is not None:
            self.capnp_port.delStream(name)
