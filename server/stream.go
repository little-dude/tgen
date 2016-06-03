package server

import (
	"fmt"
	schemas "github.com/little-dude/tgen/capnp"
	"zombiezen.com/go/capnproto2"
)

type Stream struct {
	name          string
	loop          bool
	repeat        uint32
	packetsPerSec uint32
	layers        []Layer
}

type Layer interface {
	ToCapnp(*capnp.Segment) schemas.Protocol
	FromCapnp(schemas.Protocol)
}

type Ethernet2 struct {
	source       LongField
	destination  LongField
	ethernetType Field16
}

func (ethernet2 *Ethernet2) String() string {
	return fmt.Sprint(
		"source: ", ethernet2.source,
		", destination: ", ethernet2.destination,
		", ethernet type: ", ethernet2.ethernetType)
}

func NewEthernet2() *Ethernet2 {
	ethernet2 := Ethernet2{}
	ethernet2.source.fullMask = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	ethernet2.destination.fullMask = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	ethernet2.ethernetType.fullMask = uint16(0xffff)
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
	Error.Println(ethernet2.source)
	ethernet2.source.FromCapnp(&field)
	field, _ = eth.Destination()
	ethernet2.destination.FromCapnp(&field)
	field, _ = eth.EthernetType()
	ethernet2.ethernetType.FromCapnp(&field)
}

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

func NewIpv4() *Ipv4 {
	ipv4 := Ipv4{}
	ipv4.source.fullMask = uint32(0xffffffff)
	ipv4.destination.fullMask = uint32(0xffffffff)
	ipv4.version.fullMask = uint8(0x0f) // uint8(0xf0)
	ipv4.ihl.fullMask = uint8(0x0f)
	ipv4.tos.fullMask = uint8(0xff)
	ipv4.length.fullMask = uint16(0xffff)
	ipv4.id.fullMask = uint16(0xffff)
	ipv4.flags.fullMask = uint8(0x07) // uint8(0xe0)
	ipv4.fragOffset.fullMask = uint16(0x1fff)
	ipv4.ttl.fullMask = uint8(0xff)
	ipv4.protocol.fullMask = uint8(0xff)
	ipv4.checksum.fullMask = uint16(0xffff)
	ipv4.options.fullMask = []byte{}
	ipv4.padding.fullMask = []byte{}
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

func (stream *Stream) GetConfig(call schemas.Stream_getConfig) error {
	Trace.Println("GetConfig called on ", stream)
	seg := call.Results.Segment()
	capnpStream, _ := schemas.NewStream_Config(seg)
	capnpStream.SetName(stream.name)
	capnpStream.SetLoop(stream.loop)
	capnpStream.SetRepeat(stream.repeat)
	capnpStream.SetPacketsPerSec(stream.packetsPerSec)
	return call.Results.SetConfig(capnpStream)
	return nil
}

func (stream *Stream) SetConfig(call schemas.Stream_setConfig) error {
	Trace.Println("SetConfig called on ", stream)
	capnpStream, _ := call.Params.Config()
	if capnpStream.HasName() == true {
		stream.name, _ = capnpStream.Name()
	}
	stream.loop = capnpStream.Loop()
	stream.repeat = capnpStream.Repeat()
	stream.packetsPerSec = capnpStream.PacketsPerSec()
	return nil
}

func (stream *Stream) GetLayers(call schemas.Stream_getLayers) error {
	Trace.Println("GetLayers called on ", stream)
	seg := call.Results.Segment()
	capnpLayers, _ := schemas.NewProtocol_List(seg, int32(len(stream.layers)))
	for idx, layer := range stream.layers {
		Trace.Println("layer ", layer)
		capnpLayers.Set(idx, layer.ToCapnp(seg))
	}
	return call.Results.SetLayers(capnpLayers)
	return nil
}

func (stream *Stream) SetLayers(call schemas.Stream_setLayers) error {
	Trace.Println("SetLayers called on ", stream)
	capnpLayers, _ := call.Params.Layers()
	stream.layers = make([]Layer, capnpLayers.Len())
	for i := 0; i < capnpLayers.Len(); i++ {
		capnpLayer := capnpLayers.At(i)
		stream.layers[i] = NewLayer(capnpLayer)
		Trace.Println("layer ", stream.layers[i])
	}
	return nil
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
		Error.Println("Unknown layer", which)
	}
	layer.FromCapnp(capnpLayer)
	return layer
}
