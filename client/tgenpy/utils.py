# -*- coding: utf-8 -*-

from __future__ import unicode_literals, print_function
from builtins import hex
from builtins import str
import binascii
import netaddr


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


def to_bytes(data):
    """
    Convert int, str, and lists to bytes, which correspond to capn'proto `Data`
    type
    """
    if isinstance(data, int):
        return int2bytes(data)
    elif isinstance(data, str):
        return int2bytes(str2int(data))
    elif isinstance(data, (tuple, list)):
        data = int2bytes(iter2int(data))
    else:
        raise ValueError('Cannot parse {}'.format(data))


def int2bytes(value, length=0):
    hex_str = hex(value)[2:]
    if len(hex_str) % 2 == 1:
        hex_str = '0{}'.format(hex_str)
    byte_str = binascii.unhexlify(hex_str)
    if length > 0:
        while len(byte_str) < length:
            byte_str = b'\x00' + byte_str
    return byte_str


def bytes2str(value):
    string = binascii.hexlify(value).decode()
    if string == '':
        string = '00'
    return string


def ip2bytes(ip_str):
    """
    convert strings representing ip addresses (e.g. 10.150.0.1) to bytes
    """
    return int2bytes(netaddr.IPAddress(ip_str).value, length=4)


def bytes2ip(ip_bytes):
    ip_int = 0
    # from http://stackoverflow.com/q/444591/1836144
    # note that int.from_bytes does this but requires python >= 3.2
    for i, byte in enumerate(ip_bytes[::-1]):
        ip_int += ord(byte) << i*8
    return str(netaddr.IPAddress(ip_int))


def mac2bytes(mac_str):
    return int2bytes(netaddr.EUI(mac_str).value, length=6)


def bytes2mac(mac_bytes):
    mac_int = 0
    # from http://stackoverflow.com/q/444591/1836144
    # note that int.from_bytes does this but requires python >= 3.2
    for i, byte in enumerate(mac_bytes[::-1]):
        mac_int += ord(byte) << i*8
    return str(netaddr.EUI(mac_int))
