package server

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/little-dude/tgen/schemas"
	"net"
)

type Ipv4 struct {
	source      Field32
	destination Field32
	version     Field8
	ihl         Field8
	tos         Field8
	length      Field16
	id          Field16
	flags       Field8
	fragOffset  Field16
	ttl         Field8
	protocol    Field8
	checksum    Field16
	options     LongField
	padding     LongField
}

func (ipv4 *Ipv4) MinCount() uint32 {
	return lcm(
		ipv4.destination.Count,
		ipv4.source.Count,
		ipv4.version.Count,
		ipv4.ihl.Count,
		ipv4.tos.Count,
		ipv4.length.Count,
		ipv4.id.Count,
		ipv4.flags.Count,
		ipv4.fragOffset.Count,
		ipv4.ttl.Count,
		ipv4.protocol.Count,
		ipv4.checksum.Count)
	// ipv4.options.Count,
	// ipv4.padding.Count,
}

func (ipv4 *Ipv4) SetFields(i uint) {
	ipv4.source.SetCurrentValue(i)
	ipv4.destination.SetCurrentValue(i)
	ipv4.version.SetCurrentValue(i)
	ipv4.ihl.SetCurrentValue(i)
	ipv4.tos.SetCurrentValue(i)
	ipv4.length.SetCurrentValue(i)
	ipv4.id.SetCurrentValue(i)
	ipv4.flags.SetCurrentValue(i)
	ipv4.fragOffset.SetCurrentValue(i)
	ipv4.ttl.SetCurrentValue(i)
	ipv4.protocol.SetCurrentValue(i)
	ipv4.checksum.SetCurrentValue(i)
}

func (ipv4 *Ipv4) ToPackets() []gopacket.SerializableLayer {
	ipv4.source.FirstValue = ipv4.source.Value
	ipv4.destination.FirstValue = ipv4.destination.Value
	ipv4.version.FirstValue = ipv4.version.Value
	ipv4.ihl.FirstValue = ipv4.ihl.Value
	ipv4.tos.FirstValue = ipv4.tos.Value
	ipv4.length.FirstValue = ipv4.length.Value
	ipv4.id.FirstValue = ipv4.id.Value
	ipv4.flags.FirstValue = ipv4.flags.Value
	ipv4.fragOffset.FirstValue = ipv4.fragOffset.Value
	ipv4.ttl.FirstValue = ipv4.ttl.Value
	ipv4.protocol.FirstValue = ipv4.protocol.Value
	ipv4.checksum.FirstValue = ipv4.checksum.Value
	ipv4.options.FirstValue = ipv4.options.Value
	ipv4.padding.FirstValue = ipv4.padding.Value

	count := ipv4.MinCount()
	res := make([]gopacket.SerializableLayer, count)
	for i := uint32(0); i < count; i++ {
		ipv4.SetFields(uint(i))
		res[i] = &layers.IPv4{
			SrcIP:      net.IP(ipv4.source.GetValue()),
			DstIP:      net.IP(ipv4.destination.GetValue()),
			Version:    ipv4.version.Value,
			IHL:        ipv4.ihl.Value,
			TOS:        ipv4.tos.Value,
			Length:     ipv4.length.Value,
			Id:         ipv4.id.Value,
			Flags:      layers.IPv4Flag(ipv4.flags.Value),
			FragOffset: ipv4.fragOffset.Value,
			TTL:        ipv4.ttl.Value,
			Protocol:   layers.IPProtocol(ipv4.protocol.Value),
			Checksum:   ipv4.checksum.Value,
			// Options: ipv4.options.Value,
			// Padding: ipv4.padding.Value,
		}
	}
	return res
}

func (ipv4 *Ipv4) String() string {
	return fmt.Sprint(
		"source: ", ipv4.source,
		", destination: ", ipv4.destination,
		", version: ", ipv4.version,
		", ihl: ", ipv4.ihl,
		", tos: ", ipv4.tos,
		", length:", ipv4.length,
		", id:", ipv4.id,
		", flags:", ipv4.flags,
		", fragOffset:", ipv4.fragOffset,
		", ttl:", ipv4.ttl,
		", protocol:", ipv4.protocol,
		", checksum:", ipv4.checksum,
		", options:", ipv4.options,
		", padding:", ipv4.padding)
}

func NewIpv4() *Ipv4 {
	ipv4 := Ipv4{}
	ipv4.source.FullMask = uint32(0xffffffff)
	ipv4.destination.FullMask = uint32(0xffffffff)
	ipv4.version.FullMask = uint8(0x0f) // uint8(0xf0)
	ipv4.ihl.FullMask = uint8(0x0f)
	ipv4.tos.FullMask = uint8(0xff)
	ipv4.length.FullMask = uint16(0xffff)
	ipv4.id.FullMask = uint16(0xffff)
	ipv4.flags.FullMask = uint8(0x07) // uint8(0xe0)
	ipv4.fragOffset.FullMask = uint16(0x1fff)
	ipv4.ttl.FullMask = uint8(0xff)
	ipv4.protocol.FullMask = uint8(0xff)
	ipv4.checksum.FullMask = uint16(0xffff)
	ipv4.options.FullMask = []byte{}
	ipv4.padding.FullMask = []byte{}
	return &ipv4
}

func (ipv4 *Ipv4) ToCapnp(capnpProtocol *schemas.Protocol) error {
	var field schemas.Field
	var e error

	capnpProtocol.SetIpv4()
	ip := capnpProtocol.Ipv4()

	field, e = ip.NewSource()
	if e != nil {
		return e
	}
	ipv4.source.ToCapnp(&field)

	field, e = ip.NewDestination()
	if e != nil {
		return e
	}
	ipv4.destination.ToCapnp(&field)

	field, e = ip.NewVersion()
	if e != nil {
		return e
	}
	ipv4.version.ToCapnp(&field)

	field, e = ip.NewIhl()
	if e != nil {
		return e
	}
	ipv4.ihl.ToCapnp(&field)

	field, e = ip.NewTos()
	if e != nil {
		return e
	}
	ipv4.tos.ToCapnp(&field)

	field, e = ip.NewLength()
	if e != nil {
		return e
	}
	ipv4.length.ToCapnp(&field)

	field, e = ip.NewId()
	if e != nil {
		return e
	}
	ipv4.id.ToCapnp(&field)

	field, e = ip.NewFlags()
	if e != nil {
		return e
	}
	ipv4.flags.ToCapnp(&field)

	field, e = ip.NewFragOffset()
	if e != nil {
		return e
	}
	ipv4.fragOffset.ToCapnp(&field)

	field, e = ip.NewTtl()
	if e != nil {
		return e
	}
	ipv4.ttl.ToCapnp(&field)

	field, e = ip.NewProtocol()
	if e != nil {
		return e
	}
	ipv4.protocol.ToCapnp(&field)

	field, e = ip.NewChecksum()
	if e != nil {
		return e
	}
	ipv4.checksum.ToCapnp(&field)

	// field, e = ip.NewOptions()
	// if e != nil {
	// 	return e
	// }
	// ipv4.options.ToCapnp(&field)
	//
	// field, e = ip.NewPadding()
	// if e != nil {
	// 	return e
	// }
	// ipv4.options.ToCapnp(&field)
	return nil
}

func (ipv4 *Ipv4) FromCapnp(capnpProtocol *schemas.Protocol) error {
	var field schemas.Field
	var e error

	ip := capnpProtocol.Ipv4()

	field, e = ip.Source()
	if e != nil {
		return e
	}
	ipv4.source.FromCapnp(&field)

	field, e = ip.Destination()
	if e != nil {
		return e
	}
	ipv4.destination.FromCapnp(&field)

	field, e = ip.Version()
	if e != nil {
		return e
	}
	ipv4.version.FromCapnp(&field)

	field, e = ip.Ihl()
	if e != nil {
		return e
	}
	ipv4.ihl.FromCapnp(&field)

	field, e = ip.Tos()
	if e != nil {
		return e
	}
	ipv4.tos.FromCapnp(&field)

	field, e = ip.Length()
	if e != nil {
		return e
	}
	ipv4.length.FromCapnp(&field)

	field, e = ip.Id()
	if e != nil {
		return e
	}
	ipv4.id.FromCapnp(&field)

	field, e = ip.Flags()
	if e != nil {
		return e
	}
	ipv4.flags.FromCapnp(&field)

	field, e = ip.FragOffset()
	if e != nil {
		return e
	}
	ipv4.fragOffset.FromCapnp(&field)

	field, e = ip.Ttl()
	if e != nil {
		return e
	}
	ipv4.ttl.FromCapnp(&field)

	field, e = ip.Protocol()
	if e != nil {
		return e
	}
	ipv4.protocol.FromCapnp(&field)

	field, e = ip.Checksum()
	if e != nil {
		return e
	}
	ipv4.checksum.FromCapnp(&field)

	// field, e = ip.Options()
	// if e != nil {
	// 	return e
	// }
	// ipv4.options.FromCapnp(&field)

	// field, _ = ip.Padding()
	// if e != nil {
	// 	return e
	// }
	// ipv4.padding.FromCapnp(&field)
	return nil
}
