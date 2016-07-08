package stateless

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/little-dude/tgen/server/schemas"
	"net"
)

// THA = Target Hardware Address
// TIA = Target Internet Address
// SHA = Sender Hardware Address
// SIA = Sender Internet Address

type ARP struct {
	HardwareType          Field16
	ProtocolType          Field16
	HardwareLength        Field8
	ProtocolLength        Field8
	Operation             Field16
	SenderHardwareAddress Field64
	SenderProtocolAddress Field32
	TargetHardwareAddress Field64
	TargetProtocolAddress Field32
}

func (arp *ARP) MinCount() uint32 {
	return lcm(
		arp.HardwareType.Count,
		arp.ProtocolType.Count,
		arp.HardwareLength.Count,
		arp.ProtocolLength.Count,
		arp.Operation.Count,
		arp.SenderHardwareAddress.Count,
		arp.SenderProtocolAddress.Count,
		arp.TargetHardwareAddress.Count,
		arp.TargetProtocolAddress.Count)
}

func (arp *ARP) SetFields(i uint) {
	arp.HardwareType.SetCurrentValue(i)
	arp.ProtocolType.SetCurrentValue(i)
	arp.HardwareLength.SetCurrentValue(i)
	arp.ProtocolLength.SetCurrentValue(i)
	arp.Operation.SetCurrentValue(i)
	arp.SenderHardwareAddress.SetCurrentValue(i)
	arp.SenderProtocolAddress.SetCurrentValue(i)
	arp.TargetHardwareAddress.SetCurrentValue(i)
	arp.TargetProtocolAddress.SetCurrentValue(i)
}

func (arp *ARP) ToPackets() []gopacket.SerializableLayer {
	arp.HardwareType.FirstValue = arp.HardwareType.Value
	arp.ProtocolType.FirstValue = arp.ProtocolType.Value
	arp.HardwareLength.FirstValue = arp.HardwareLength.Value
	arp.ProtocolLength.FirstValue = arp.ProtocolLength.Value
	arp.Operation.FirstValue = arp.Operation.Value
	arp.SenderHardwareAddress.FirstValue = arp.SenderHardwareAddress.Value
	arp.SenderProtocolAddress.FirstValue = arp.SenderProtocolAddress.Value
	arp.TargetHardwareAddress.FirstValue = arp.TargetHardwareAddress.Value
	arp.TargetProtocolAddress.FirstValue = arp.TargetProtocolAddress.Value

	count := arp.MinCount()
	res := make([]gopacket.SerializableLayer, count)
	for i := uint32(0); i < count; i++ {
		arp.SetFields(uint(i))
		res[i] = &layers.ARP{
			AddrType:          layers.LinkType(arp.HardwareType.Value),
			Protocol:          layers.EthernetType(arp.ProtocolType.Value),
			HwAddressSize:     arp.HardwareLength.Value,
			ProtAddressSize:   arp.ProtocolLength.Value,
			Operation:         arp.Operation.Value,
			SourceHwAddress:   net.HardwareAddr(arp.SenderHardwareAddress.GetValue()[2:]),
			SourceProtAddress: net.IP(arp.SenderProtocolAddress.GetValue()),
			DstHwAddress:      net.HardwareAddr(arp.TargetHardwareAddress.GetValue()[2:]),
			DstProtAddress:    net.IP(arp.TargetProtocolAddress.GetValue()),
		}
	}
	return res
}

func (arp *ARP) String() string {
	return fmt.Sprint(
		"hardware type:", arp.HardwareType,
		"protocol type:", arp.ProtocolType,
		"hardware length", arp.HardwareLength,
		"protocol length", arp.ProtocolLength,
		"operation", arp.Operation,
		"sender hardware address", arp.SenderHardwareAddress,
		"sender protocol address", arp.SenderProtocolAddress,
		"target hardware address", arp.TargetHardwareAddress,
		"target protocol address", arp.TargetProtocolAddress)
}

func NewARP() *ARP {
	arp := ARP{}
	arp.HardwareType.FullMask = uint16(0xffff)
	arp.HardwareType.Value = uint16(layers.LinkTypeEthernet)
	arp.ProtocolType.FullMask = uint16(0xffff)
	arp.ProtocolType.Value = uint16(layers.EthernetTypeARP)
	arp.HardwareLength.FullMask = uint8(0xff)
	arp.HardwareLength.Value = uint8(0x06)
	arp.ProtocolLength.FullMask = uint8(0xff)
	arp.ProtocolLength.Value = uint8(0x04)
	arp.Operation.FullMask = uint16(0xffff)
	arp.SenderHardwareAddress.FullMask = uint64(0x0000ffffffffffff)
	arp.SenderProtocolAddress.FullMask = uint32(0xffffffff)
	arp.TargetHardwareAddress.FullMask = uint64(0x0000ffffffffffff)
	arp.TargetProtocolAddress.FullMask = uint32(0xffffffff)
	return &arp
}

func (arp *ARP) ToCapnp(capnpProtocol *schemas.Protocol) error {
	var field schemas.Field
	var e error

	capnpProtocol.SetArp()
	capnpArp := capnpProtocol.Arp()

	field, e = capnpArp.NewHardwareType()
	if e != nil {
		return e
	}
	arp.HardwareType.ToCapnp(&field)

	field, e = capnpArp.NewProtocolType()
	if e != nil {
		return e
	}
	arp.ProtocolType.ToCapnp(&field)

	field, e = capnpArp.NewHardwareLength()
	if e != nil {
		return e
	}
	arp.ProtocolType.ToCapnp(&field)

	field, e = capnpArp.NewProtocolLength()
	if e != nil {
		return e
	}
	arp.ProtocolLength.ToCapnp(&field)

	field, e = capnpArp.NewOperation()
	if e != nil {
		return e
	}
	arp.Operation.ToCapnp(&field)

	field, e = capnpArp.NewSenderHardwareAddress()
	if e != nil {
		return e
	}
	arp.SenderHardwareAddress.ToCapnp(&field)

	field, e = capnpArp.NewSenderProtocolAddress()
	if e != nil {
		return e
	}
	arp.SenderProtocolAddress.ToCapnp(&field)

	field, e = capnpArp.NewTargetHardwareAddress()
	if e != nil {
		return e
	}
	arp.TargetHardwareAddress.ToCapnp(&field)

	field, e = capnpArp.NewTargetProtocolAddress()
	if e != nil {
		return e
	}
	arp.TargetProtocolAddress.ToCapnp(&field)

	return nil
}

func (arp *ARP) FromCapnp(capnpProtocol *schemas.Protocol) error {
	var field schemas.Field
	var e error

	capnpArp := capnpProtocol.Arp()

	field, e = capnpArp.HardwareType()
	if e != nil {
		return e
	}
	arp.HardwareType.FromCapnp(&field)

	field, e = capnpArp.ProtocolType()
	if e != nil {
		return e
	}
	arp.ProtocolType.FromCapnp(&field)

	field, e = capnpArp.HardwareLength()
	if e != nil {
		return e
	}
	arp.HardwareLength.FromCapnp(&field)

	field, e = capnpArp.ProtocolLength()
	if e != nil {
		return e
	}
	arp.ProtocolLength.FromCapnp(&field)

	field, e = capnpArp.Operation()
	if e != nil {
		return e
	}
	arp.Operation.FromCapnp(&field)

	field, e = capnpArp.SenderHardwareAddress()
	if e != nil {
		return e
	}
	arp.SenderHardwareAddress.FromCapnp(&field)

	field, e = capnpArp.SenderProtocolAddress()
	if e != nil {
		return e
	}
	arp.SenderProtocolAddress.FromCapnp(&field)

	field, e = capnpArp.TargetHardwareAddress()
	if e != nil {
		return e
	}
	arp.TargetHardwareAddress.FromCapnp(&field)

	field, e = capnpArp.TargetProtocolAddress()
	if e != nil {
		return e
	}
	arp.TargetProtocolAddress.FromCapnp(&field)

	return nil
}
