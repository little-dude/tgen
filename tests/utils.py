# -*- coding: utf-8 -*-

from __future__ import unicode_literals
import atexit
import time
import platform
import logging
import os
import sys
import re
import six
import pyshark
from pyroute2 import IPRoute
from threading import Thread
# from multiprocessing import Process
from tgenpy import Controller

try:
    from Queue import Queue, Empty
except ImportError:
    from queue import Queue, Empty  # python 3.x

if os.name == 'posix' and sys.version_info[0] < 3:
    from subprocess32 import Popen, PIPE, TimeoutExpired
else:
    from subprocess import Popen, PIPE, TimeoutExpired


LOG = logging.getLogger(__name__)


def link_down(name):
    ip = IPRoute()
    ip.link('set', index=ip.link_lookup(ifname=name)[0], state='down')


def link_up(name):
    ip = IPRoute()
    ip.link('set', index=ip.link_lookup(ifname=name)[0], state='up')


def create_veth_pair(name):
    ip = IPRoute()
    peers = ('{}0'.format(name), '{}1'.format(name))
    LOG.info('creating veth pair {}'.format(peers))
    ip.link('add', kind='veth', ifname=peers[0], peer=peers[1])
    link_up(peers[0])
    link_up(peers[1])


def delete_veth_pair(name):
    ip = IPRoute()
    peers = ('{}0'.format(name), '{}1'.format(name))
    LOG.info('deleting veth pair {}'.format(peers))
    link_down(peers[0])
    link_down(peers[1])
    ip.link('del', index=ip.link_lookup(ifname=peers[0])[0])


def create_port(name):
    LOG.info('creating port {}'.format(name))
    ip = IPRoute()
    ip.link('add', ifname=name, kind='dummy')
    link_up(name)


def delete_port(name):
    LOG.info('deleting port {}'.format(name))
    ip = IPRoute()
    link_down(name)
    ip.link('del', index=ip.link_lookup(ifname=name)[0])


def is_pypy():
    if platform.python_implementation() == 'PyPy':
        return True
    return False


def send_and_receive(tx, rx, duration=1, clear_stats=True, save_as=None):
    raise NotImplemented()


TGEN_ERR_Q = None
TGEN_OUT_Q = None
TGEN = None


def enqueue_output(out, queue):
    for line in iter(out.readline, b''):
        # convert the input to a unicode object
        line = ensure_native_str(line).strip('\n')
        LOG.debug(line)
        queue.put(line)
    out.close()


def readlines(q):
    lines = []
    while True:
        try:
            lines.append(q.get_nowait())
        except Empty:
            return lines


def start_tgen():
    global TGEN, TGEN_ERR_Q, TGEN_OUT_Q
    LOG.info('starting tgen')
    TGEN = Popen(['./tgen'], stdout=PIPE, stderr=PIPE, bufsize=1)
    TGEN_OUT_Q = Queue()
    TGEN_ERR_Q = Queue()
    for stream, q in [(TGEN.stdout, TGEN_OUT_Q), (TGEN.stderr, TGEN_ERR_Q)]:
        t = Thread(target=enqueue_output, args=(stream, q))
        t.daemon = True
        t.start()
    start_time = time.time()
    stdout = []
    while time.time() - start_time < 5:
        err = readlines(TGEN_ERR_Q)
        if err:
            raise Exception('Output on stderr: {}'.format(err))
        out = readlines(TGEN_OUT_Q)
        for line in out:
            if re.match('.*Waiting for connections.*', line, re.UNICODE):
                LOG.info('tgen is ready')
                return
        stdout.extend(out)
    raise Exception('tgen is not ready: {}'.format(stdout))


def kill_tgen():
    LOG.info('trying to stop tgen gracefully')
    TGEN.terminate()
    try:
        TGEN.wait(timeout=10)
        LOG.info('tgen exited gracefully')
    except TimeoutExpired:
        LOG.info('could not terminate tgen properly, kill it.')
        TGEN.kill()
        TGEN.wait(timeout=10)
        LOG.info('tgen has been killed')


def restart_tgen():
    kill_tgen()
    start_tgen()
    return Controller('localhost')


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


class Capture(object):

    def __init__(self, port, count=0, timeout=0):
        self.port = port
        self.capture_file = 'test.pcap'
        self.count = count
        self.timeout = timeout

    def __enter__(self):
        self.port.start_capture(
            self.capture_file,
            timeout=self.timeout,
            packet_count=self.count)
        return self

    def get_packets(self, timeout=0):
        done, error = self.port.wait_capture(timeout=timeout)
        assert done is True
        assert error == ''
        capture = pyshark.FileCapture(self.capture_file)
        capture.load_packets()
        return capture

    def __exit__(self, type, value, traceback):
        # TODO: not sure if we need to do anything here
        pass

# class Capture(object):
#
#     def __init__(self, interface, packets=0, timeout=0):
#         self.interface = interface
#         self.packets = packets
#         self.timeout = timeout
#
#     def __enter__(self):
#         self.start_capture()
#         return self
#
#     def __exit__(self, type, value, traceback):
#         # TODO: not sure if we need to do anything here
#         pass
#
#     def start_capture(self):
#         self.capture = pyshark.LiveCapture(interface=self.interface)
#         self.capture.set_debug()
#
#         def load_packets(*args, **kwargs):
#             sys.stdout = open(str(os.getpid()) + ".out", "w")
#             return self.capture.load_packets(*args, **kwargs)
#
#         self.capture_process = Process(
#             target=load_packets,
#             kwargs={
#                 'packet_count': self.packets,
#                 'timeout': self.timeout
#             }
#         )
#         self.capture_process.start()
#
#     def wait(self):
#         self.capture_process.join()
#         return self.capture


def cleanup():
    global TGEN
    if isinstance(TGEN, Popen):
        if TGEN.poll() is None:
            kill_tgen()


atexit.register(cleanup)
