package server

import (
	schemas "github.com/little-dude/tgen/capnp"
	"github.com/little-dude/tgen/server/log"
	"github.com/little-dude/tgen/server/protocols"
)

type Stream struct {
	name          string
	loop          bool
	repeat        uint32
	packetsPerSec uint32
	layers        []protocols.Layer
}

func (stream *Stream) GetConfig(call schemas.Stream_getConfig) error {
	log.Trace.Println("GetConfig called on ", stream)
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
	log.Trace.Println("SetConfig called on ", stream)
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
	log.Trace.Println("GetLayers called on ", stream)
	seg := call.Results.Segment()
	capnpLayers, _ := schemas.NewProtocol_List(seg, int32(len(stream.layers)))
	for idx, layer := range stream.layers {
		log.Trace.Println("layer ", layer)
		capnpLayers.Set(idx, layer.ToCapnp(seg))
	}
	return call.Results.SetLayers(capnpLayers)
	return nil
}

func (stream *Stream) SetLayers(call schemas.Stream_setLayers) error {
	log.Trace.Println("SetLayers called on ", stream)
	capnpLayers, _ := call.Params.Layers()
	stream.layers = make([]protocols.Layer, capnpLayers.Len())
	for i := 0; i < capnpLayers.Len(); i++ {
		capnpLayer := capnpLayers.At(i)
		stream.layers[i] = protocols.NewLayer(capnpLayer)
		log.Trace.Println("layer ", stream.layers[i])
	}
	return nil
}
