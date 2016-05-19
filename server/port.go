package server

import (
	schema "github.com/little-dude/tgen/capnp"
)

// Port represent one of the host ports
type Port struct {
	name string
}

// Getname is a capability that returns the name of the port
func (port Port) GetName(call schema.Port_getName) error {
	Trace.Println("GetName called on ", port)
	return call.Results.SetName(port.name)
}
