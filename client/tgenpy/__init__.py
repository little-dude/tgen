from __future__ import unicode_literals
import capnp


def make_schemas_module():
    """
    small hack to make the schemas available in ``sys.modules``
    """
    import os
    import sys

    capnp.remove_import_hook()
    sys.modules['schemas'] = capnp.load(
        os.path.join(
            os.path.dirname(os.path.realpath(__file__)), 'schemas.capnp'))

make_schemas_module()

from .protocols import *
from .controller import Controller
from .stream import Stream

__all__ = ['Controller', 'Stream']
