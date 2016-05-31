package server

import (
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
	source       Field
	destination  Field
	ethernetType Field
}

func (ethernet2 *Ethernet2) ToCapnp(seg *capnp.Segment) (p schemas.Protocol) {
	protocol, _ := schemas.NewProtocol(seg)
	eth := protocol.Ethernet2()
	eth.SetSource(ethernet2.source.ToCapnp(seg))
	eth.SetDestination(ethernet2.destination.ToCapnp(seg))
	eth.SetEthernetType(ethernet2.ethernetType.ToCapnp(seg))
	return schemas.Protocol(eth)
}

func (ethernet2 *Ethernet2) FromCapnp(data schemas.Protocol) {
	eth := data.Ethernet2()
	field, _ := eth.Source()
	ethernet2.source = NewField(field)
	field, _ = eth.Destination()
	ethernet2.destination = NewField(field)
	field, _ = eth.EthernetType()
	ethernet2.ethernetType = NewField(field)
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
	}
	return nil
}

func NewLayer(capnpLayer schemas.Protocol) Layer {
	which := capnpLayer.Which()
	var layer Layer
	if which == schemas.Protocol_Which_ethernet2 {
		layer = &Ethernet2{}
	} else {
		panic("Unknown layer")
	}
	layer.FromCapnp(capnpLayer)
	return layer
}
