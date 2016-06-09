package server

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/little-dude/tgen/schemas"
	"net"
	// "zombiezen.com/go/capnproto2"
)

type Ethernet2 struct {
	source       LongField
	destination  LongField
	ethernetType Field16
}

func (ethernet2 *Ethernet2) MinCount() uint32 {
	return lcm(
		ethernet2.destination.Count,
		ethernet2.source.Count,
		ethernet2.ethernetType.Count)
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

func (ethernet2 *Ethernet2) ToCapnp(capnpProtocol *schemas.Protocol) error {
	var field schemas.Field
	var e error

	capnpProtocol.SetEthernet2()
	eth := capnpProtocol.Ethernet2()

	field, e = eth.NewSource()
	if e != nil {
		return e
	}
	ethernet2.source.ToCapnp(&field)

	field, e = eth.NewDestination()
	if e != nil {
		return e
	}
	ethernet2.destination.ToCapnp(&field)

	field, e = eth.NewEthernetType()
	if e != nil {
		return e
	}
	ethernet2.ethernetType.ToCapnp(&field)

	return nil
}

func (ethernet2 *Ethernet2) FromCapnp(capnpProtocol *schemas.Protocol) error {
	var field schemas.Field
	var e error

	eth := capnpProtocol.Ethernet2()

	field, e = eth.Source()
	if e != nil {
		return e
	}
	ethernet2.source.FromCapnp(&field)

	field, e = eth.Destination()
	if e != nil {
		return e
	}
	ethernet2.destination.FromCapnp(&field)

	field, e = eth.EthernetType()
	if e != nil {
		return e
	}
	ethernet2.ethernetType.FromCapnp(&field)

	return nil
}

func (ethernet2 *Ethernet2) SetFields(i uint) {
	ethernet2.source.SetCurrentValue(i)
	ethernet2.destination.SetCurrentValue(i)
	ethernet2.ethernetType.SetCurrentValue(i)
}

func (ethernet2 *Ethernet2) ToPackets() []gopacket.SerializableLayer {
	ethernet2.source.FirstValue = ethernet2.source.Value
	ethernet2.destination.FirstValue = ethernet2.destination.Value
	ethernet2.ethernetType.FirstValue = ethernet2.ethernetType.Value

	count := ethernet2.MinCount()
	res := make([]gopacket.SerializableLayer, count)
	for i := uint32(0); i < count; i++ {
		ethernet2.SetFields(uint(i))
		res[i] = &layers.Ethernet{
			SrcMAC:       net.HardwareAddr(ethernet2.source.Value),
			DstMAC:       net.HardwareAddr(ethernet2.destination.Value),
			EthernetType: layers.EthernetType(ethernet2.ethernetType.Value),
		}
	}
	return res
}
