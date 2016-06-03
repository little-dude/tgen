package protocols

import (
	"fmt"
	schemas "github.com/little-dude/tgen/capnp"
	"github.com/little-dude/tgen/server/protocols/fields"
	"zombiezen.com/go/capnproto2"
)

type Ethernet2 struct {
	source       fields.LongField
	destination  fields.LongField
	ethernetType fields.Field16
}

func (ethernet2 *Ethernet2) String() string {
	return fmt.Sprint(
		"source: ", ethernet2.source,
		", destination: ", ethernet2.destination,
		", ethernet type: ", ethernet2.ethernetType)
}

func NewEthernet2() *Ethernet2 {
	ethernet2 := Ethernet2{}
	ethernet2.source.FullMask = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	ethernet2.destination.FullMask = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	ethernet2.ethernetType.FullMask = uint16(0xffff)
	return &ethernet2
}

func (ethernet2 *Ethernet2) ToCapnp(seg *capnp.Segment) (p schemas.Protocol) {
	protocol, _ := schemas.NewProtocol(seg)
	protocol.SetEthernet2()
	eth := protocol.Ethernet2()
	eth.SetSource(ethernet2.source.ToCapnp(seg))
	eth.SetDestination(ethernet2.destination.ToCapnp(seg))
	eth.SetEthernetType(ethernet2.ethernetType.ToCapnp(seg))
	return schemas.Protocol(eth)
}

func (ethernet2 *Ethernet2) FromCapnp(data schemas.Protocol) {
	eth := data.Ethernet2()
	field, _ := eth.Source()
	ethernet2.source.FromCapnp(&field)
	field, _ = eth.Destination()
	ethernet2.destination.FromCapnp(&field)
	field, _ = eth.EthernetType()
	ethernet2.ethernetType.FromCapnp(&field)
}
