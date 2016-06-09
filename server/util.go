package server

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
