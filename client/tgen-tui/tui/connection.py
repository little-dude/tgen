import socket
import npyscreen as nps
import capnp
import schemas
from . import ports


class ConnectionForm(nps.ActionPopup):

    def __init__(self, *args, **kwargs):
        super(ConnectionForm, self).__init__(*args, **kwargs)

    def create(self):
        self.w_hostname = self.add(nps.TitleText, name='Controller hostname')
        self.w_port = self.add(nps.TitleText, name='Controller port', value='1234')

    def on_ok(self):
        self.connect()

    def connect(self, *args, **kwargs):
        address = '{}:{}'.format(self.w_hostname.value, self.w_port.value)
        try:
            client = capnp.TwoPartyClient(address)
        except socket.gaierror:
            err = 'Failed to connect to {}'.format(address)
            nps.notify_confirm(err, title='error')
        else:
            controller = client.bootstrap().cast_as(schemas.Controller)
            self.parentApp.addForm("PORTS", ports.Form, address, controller)
            self.parentApp.switchForm("PORTS")
