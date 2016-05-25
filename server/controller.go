package server

import (
	schema "github.com/little-dude/tgen/capnp"
	"net"
	"zombiezen.com/go/capnproto2"
)

// Controller represent the controller running on the host.
type Controller struct {
	ports []Port
}

// GetPorts is a capability returning the list of the host ports
func (controller *Controller) GetPorts(call schema.Controller_getPorts) error {
	Trace.Println("GetPorts called on ", controller)
	itfs, e := net.Interfaces()
	if e != nil {
		return e
	}
	controller.updatePorts(itfs)
	capnpPorts, e := call.Results.NewPorts(int32(len(controller.ports)))
	Trace.Println(controller.ports)
	if e != nil {
		return e
	}
	resultsSeg := call.Results.Segment()
	for i, _ := range controller.ports {
		capnpPort := schema.Port_ServerToClient(&controller.ports[i])
		ptr := capnp.NewInterface(resultsSeg, resultsSeg.Message().AddCap(capnpPort.Client)).ToPtr()
		capnpPorts.SetPtr(i, ptr)
	}
	return nil
}

func (controller *Controller) updatePorts(interfaces []net.Interface) {
	var portFound bool
	for _, itf := range interfaces {
		portFound = false
		for _, p := range controller.ports {
			if p.name == itf.Name {
				Info.Println("Found existing port ", p.name)
				portFound = true
				break
			}
		}
		if portFound == false {
			Info.Println("Found new port ", itf.Name)
			controller.ports = append(controller.ports, Port{name: itf.Name})
		}
	}
}
