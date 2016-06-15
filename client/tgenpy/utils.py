from __future__ import unicode_literals
import six


def ensure_text(data, encoding='utf-8'):
    # copy-pasted from:
    # https://github.com/jparyani/pycapnp/issues/92#issue-138016674
    #
    # for encoding issues in general, see this very good (but in french)
    # article:
    # http://sametmax.com/lencoding-en-python-une-bonne-fois-pour-toute/
    if isinstance(data, six.text_type):
        return data
    elif isinstance(data, six.binary_type):
        return data.decode(encoding)
    raise ValueError('cannot ensure_text from type %r' % (type(data,)))


def ensure_native_str(data, encoding='utf-8'):
    # copy-pasted from:
    # https://github.com/jparyani/pycapnp/issues/92#issue-138016674
    #
    # for encoding issues in general, see this very good (but in french)
    # article:
    # http://sametmax.com/lencoding-en-python-une-bonne-fois-pour-toute/
    if isinstance(data, str):
        return data
    elif six.PY2 and isinstance(data, six.text_type):  # py2 "unicode"
        return data.encode(encoding)
    elif six.PY3 and isinstance(data, six.binary_type):  # py3 "bytes"
        return data.decode(encoding)
    raise ValueError('cannot ensure_native_str from type %r' % (type(data,)))
