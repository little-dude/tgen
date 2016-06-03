package protocols

import (
	schemas "github.com/little-dude/tgen/capnp"
	"github.com/little-dude/tgen/server/protocols/fields"
	"zombiezen.com/go/capnproto2"
)

type Ipv4 struct {
	source      fields.Field32
	destination fields.Field32
	version     fields.Field8
	ihl         fields.Field8
	tos         fields.Field8
	length      fields.Field16
	id          fields.Field16
	flags       fields.Field8
	fragOffset  fields.Field16
	ttl         fields.Field8
	protocol    fields.Field8
	checksum    fields.Field16
	options     fields.LongField
	padding     fields.LongField
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

func (ipv4 *Ipv4) ToCapnp(seg *capnp.Segment) (p schemas.Protocol) {
	protocol, _ := schemas.NewProtocol(seg)
	protocol.SetIpv4()
	ip := protocol.Ipv4()
	ip.SetSource(ipv4.source.ToCapnp(seg))
	ip.SetDestination(ipv4.destination.ToCapnp(seg))
	ip.SetVersion(ipv4.version.ToCapnp(seg))
	ip.SetIhl(ipv4.ihl.ToCapnp(seg))
	ip.SetTos(ipv4.tos.ToCapnp(seg))
	ip.SetLength(ipv4.length.ToCapnp(seg))
	ip.SetId(ipv4.id.ToCapnp(seg))
	ip.SetFlags(ipv4.flags.ToCapnp(seg))
	ip.SetFragOffset(ipv4.fragOffset.ToCapnp(seg))
	ip.SetTtl(ipv4.ttl.ToCapnp(seg))
	ip.SetProtocol(ipv4.protocol.ToCapnp(seg))
	ip.SetChecksum(ipv4.checksum.ToCapnp(seg))
	ip.SetOptions(ipv4.options.ToCapnp(seg))
	ip.SetPadding(ipv4.padding.ToCapnp(seg))
	return schemas.Protocol(ip)
}

func (ipv4 *Ipv4) FromCapnp(data schemas.Protocol) {
	ip := data.Ipv4()

	field, _ := ip.Source()
	ipv4.source.FromCapnp(&field)

	field, _ = ip.Destination()
	ipv4.destination.FromCapnp(&field)

	field, _ = ip.Version()
	ipv4.version.FromCapnp(&field)

	field, _ = ip.Ihl()
	ipv4.ihl.FromCapnp(&field)

	field, _ = ip.Tos()
	ipv4.tos.FromCapnp(&field)

	field, _ = ip.Length()
	ipv4.length.FromCapnp(&field)

	field, _ = ip.Id()
	ipv4.id.FromCapnp(&field)

	field, _ = ip.Flags()
	ipv4.flags.FromCapnp(&field)

	field, _ = ip.FragOffset()
	ipv4.fragOffset.FromCapnp(&field)

	field, _ = ip.Ttl()
	ipv4.ttl.FromCapnp(&field)

	field, _ = ip.Protocol()
	ipv4.protocol.FromCapnp(&field)

	field, _ = ip.Checksum()
	ipv4.checksum.FromCapnp(&field)

	field, _ = ip.Options()
	ipv4.options.FromCapnp(&field)

	field, _ = ip.Padding()
	ipv4.padding.FromCapnp(&field)
}
