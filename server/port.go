package server

import (
	schemas "github.com/little-dude/tgen/capnp"
	"zombiezen.com/go/capnproto2"
)

type Port struct {
	name    string
	streams []Stream
}

func (port *Port) GetConfig(call schemas.Port_getConfig) error {
	Trace.Println("GetConfig called on ", port)
	seg := call.Results.Segment()
	capnpPort, _ := schemas.NewPort_Config(seg)
	capnpPort.SetName(port.name)
	return call.Results.SetConfig(capnpPort)
}

func (port *Port) SetConfig(call schemas.Port_setConfig) error {
	Trace.Println("SetConfig called on ", port)
	// capnpPort, _ := call.Params.Config()
	return nil
}

func (port *Port) GetStreams(call schemas.Port_getStreams) error {
	Trace.Println("GetStreams called on ", port)
	seg := call.Results.Segment()
	capnpStreams, _ := call.Results.NewStreams(int32(len(port.streams)))
	for i, _ := range port.streams {
		capnpStream := schemas.Stream_ServerToClient(&port.streams[i])
		ptr := capnp.NewInterface(seg, seg.Message().AddCap(capnpStream.Client)).ToPtr()
		capnpStreams.SetPtr(i, ptr)
	}
	return nil
}

func (port *Port) NewStream(call schemas.Port_newStream) error {
	Trace.Println("NewStream called on ", port)
	stream := Stream{name: "new_stream"}
	port.streams = append(port.streams, stream)
	// Create a new locally implemented Stream capability.
	capnpStream := schemas.Stream_ServerToClient(&stream)
	// Notice that methods can return other interfaces.
	Trace.Println(port.streams)
	return call.Results.SetStream(capnpStream)
}

func (port *Port) DelStream(call schemas.Port_delStream) error {
	Trace.Println("DelStream called on ", port)
	name, _ := call.Params.Name()
	for i, stream := range port.streams {
		if stream.name == name {
			port.streams = append(port.streams[:i], port.streams[i+1:]...)
			break
		}
	}
	return nil
}
