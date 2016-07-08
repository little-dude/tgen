package stateless

import (
	"github.com/google/gopacket"
	"github.com/little-dude/tgen/server/schemas"
)

type Layer interface {
	ToCapnp(*schemas.Protocol) error
	FromCapnp(*schemas.Protocol) error
	ToPackets() []gopacket.SerializableLayer
	MinCount() uint32
}

func NewLayer(capnpLayer *schemas.Protocol) (Layer, error) {
	which := capnpLayer.Which()
	var layer Layer
	switch which {
	case schemas.Protocol_Which_ethernet2:
		layer = NewEthernet2()
	case schemas.Protocol_Which_ipv4:
		layer = NewIpv4()
	default:
		// log.log.Error.Println("Unknown layer", which)
	}
	e := layer.FromCapnp(capnpLayer)
	return layer, e
}

func lcm(ints ...uint16) uint32 {
	if len(ints) == 1 {
		return uint32(ints[0])
	}
	res := uint32(1)
	for _, integer := range ints {
		res = _lcm(res, uint32(integer))
	}
	return res
}

func _lcm(n, m uint32) uint32 {
	if m == 0 || n == 0 {
		// FIXME
		// not sure what to do here. should probably avoid calling this
		// function with 0. for now, just panic, we'll see later.
		panic("don't call me with a null argument please")
	}
	if m > n {
		return _lcm(m, n)
	}
	// LCM = (n * m) / GCD(n, m)
	// keep this around for later since we'll mutate a and b
	// FIXME: this can silently overflow. Not much we can do for this.
	p := n * m

	// find GCD with the basic euclide algo. there is probably smarter than
	// this.
	r := m
	for r != 0 {
		r = n % m
		n = m
		m = r
	}
	// n is now the GCD (latest non-zero remainder)
	return p / n
}
