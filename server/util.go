package server

import "net"

func adjustSliceLength(length int, slice []byte) []byte {
	l := len(slice)
	if l == length {
		return slice
	}
	if l < length {
		// left padding with 0s
		return append(make([]byte, length-l), slice...)
	}
	// l > length so we truncate it, taking the right-most words
	return slice[l-length:]
}

func sub(a, b byte) (result byte, overflow bool) {
	if b > a {
		return 255 - ((b - a) - 1), true
	} else {
		return (a - b), false
	}
}

func add(a, b byte) (result, overflow byte) {
	res := uint16(a) + uint16(b)
	result = byte(res & 255)
	overflow = byte(res >> 8)
	return
}

func IPv4ToInt(ip net.IP) uint32 {
	return uint32(ip[0])<<24 + uint32(ip[1])<<16 + uint32(ip[2])<<8 + uint32(ip[3])
}

func IPv4FromInt(value uint32) net.IP {
	return net.IPv4(
		byte(value>>24),
		byte((value&0x00ff0000)>>16),
		byte((value&0x0000ff00)>>8),
		byte(value&0x000000ff))
}
