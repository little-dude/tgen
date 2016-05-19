package server

import (
	schema "github.com/little-dude/tgen/capnp"
	"net"
	"zombiezen.com/go/capnproto2"
)

// Controller represent the controller running on the host.
type Controller struct{}

// GetPorts is a capability returning the list of the host ports
func (controller Controller) GetPorts(call schema.Controller_getPorts) error {
	Trace.Println("GetPorts called on ", controller)
	itfs, e := net.Interfaces()
	Info.Println("Found interface ", itfs)
	if e != nil {
		return e
	}
	ports, e := call.Results.NewPorts(int32(len(itfs)))
	if e != nil {
		return e
	}
	resultsSeg := call.Results.Segment()
	for index, itf := range itfs {
		port := schema.Port_ServerToClient(Port{name: itf.Name})
		ptr := capnp.NewInterface(resultsSeg, resultsSeg.Message().AddCap(port.Client)).ToPtr()
		ports.SetPtr(index, ptr)
	}
	return nil
}
