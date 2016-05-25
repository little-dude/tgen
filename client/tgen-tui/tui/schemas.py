import os
import sys
import capnp

capnp.remove_import_hook()
sys.modules['schemas'] = capnp.load(
    os.path.join(
        os.path.dirname(os.path.realpath(__file__)), 'schemas', 'main.capnp'))
