package server

import (
	schemas "github.com/little-dude/tgen/capnp"
)

type Stream struct {
	name          string
	loop          bool
	repeat        uint32
	packetsPerSec uint32
	// layers        []Layer
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
