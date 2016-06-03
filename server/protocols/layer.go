package protocols

import (
	schemas "github.com/little-dude/tgen/capnp"
	"github.com/little-dude/tgen/server/log"
	"zombiezen.com/go/capnproto2"
)

type Layer interface {
	ToCapnp(*capnp.Segment) schemas.Protocol
	FromCapnp(schemas.Protocol)
}

func NewLayer(capnpLayer schemas.Protocol) Layer {
	which := capnpLayer.Which()
	var layer Layer
	switch which {
	case schemas.Protocol_Which_ethernet2:
		layer = NewEthernet2()
	case schemas.Protocol_Which_ipv4:
		layer = NewIpv4()
	default:
		log.Error.Println("Unknown layer", which)
	}
	layer.FromCapnp(capnpLayer)
	return layer
}
