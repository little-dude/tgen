import npyscreen
from . import connection


class TuiApp(npyscreen.NPSAppManaged):

    def onStart(self):
        self.registerForm("MAIN", connection.ConnectionForm())


def run():
    app = TuiApp()
    app.run()
