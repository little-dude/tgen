# -*- coding: utf-8 -*-

from __future__ import unicode_literals
import capnp
from . import utils


def make_schemas_module():
    """
    small hack to make the schemas available in ``sys.modules``
    """
    import os
    import sys

    capnp.remove_import_hook()
    pkg_path = os.path.dirname(os.path.realpath(__file__))
    schemas_path = os.path.join(pkg_path, 'schemas.capnp')
    sys.modules['schemas'] = capnp.load(utils.ensure_native_str(schemas_path))

make_schemas_module()

from .protocols import *
from .controller import Controller
from .stream import Stream

__all__ = ['Controller', 'Stream']
