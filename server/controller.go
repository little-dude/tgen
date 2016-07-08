package main

import (
	"github.com/little-dude/tgen/server/errors"
	"github.com/little-dude/tgen/server/log"
	"github.com/little-dude/tgen/server/schemas"
	"github.com/little-dude/tgen/server/stateless"
	"net"
	"strconv"
	"zombiezen.com/go/capnproto2"
)

// Controller represent the controller running on the host.
type Controller struct {
	ports   map[string]*Port
	streams map[uint16]*stateless.Stream
}

func NewController() *Controller {
	return &Controller{
		ports:   make(map[string]*Port),
		streams: make(map[uint16]*stateless.Stream),
	}
}

// GetPorts is a local implementation of the capnproto capability.
func (c *Controller) GetPorts(call schemas.Controller_getPorts) error {
	// retrieve the ports
	e := c.refreshPorts()
	if e != nil {
		return e
	}

	// initialize a list of capnp interfaces
	portsList, e := call.Results.NewPorts(int32(len(c.ports)))
	if e != nil {
		return e
	}

	// populate the list
	seg := call.Results.Segment()
	i := 0
	for portName, _ := range c.ports {
		port := schemas.Port_ServerToClient(c.ports[portName])
		ptr := capnp.NewInterface(seg, seg.Message().AddCap(port.Client)).ToPtr()
		portsList.SetPtr(i, ptr)
		i++
	}
	return nil
}

func (c *Controller) refreshPorts() error {

	// list the local ports
	log.Info.Println("Listing local ports")
	itfs, e := net.Interfaces()
	if e != nil {
		log.Error.Println(e.Error())
		return errors.New("Failed to list the interfaces:", e.Error())
	}

	// add new ports
	for _, itf := range itfs {
		if _, ok := c.ports[itf.Name]; !ok {
			c.ports[itf.Name] = NewPort(itf.Name, c)
		}
	}

	// remove ports that no longer exist
	garbagePorts := make([]string, 0)
outer:
	for portName, _ := range c.ports {
		for _, itf := range itfs {
			if itf.Name == portName {
				continue outer
			}
		}
		port := c.ports[portName]
		if port.tx.State.Active() || port.rx.State.Active() {
			log.Error.Println("Port", portName, "not found but sending and/or capturing. Not removing it for now.")
		} else {
			garbagePorts = append(garbagePorts, portName)
		}
	}
	for _, portName := range garbagePorts {
		delete(c.ports, portName)
	}

	return nil
}

// ListStreams is the implementation of the capnproto ListStreams method
func (c *Controller) ListStreams(call schemas.Controller_listStreams) error {
	streamIDs, e := call.Results.NewIds(int32(len(c.streams)))
	if e != nil {
		return e
	}

	i := 0
	for ID, _ := range c.streams {
		streamIDs.Set(i, ID)
		i++
	}

	return nil
}

func (c *Controller) newStreamID() (uint16, error) {
	ID := uint16(1)
	for {
		if _, ok := c.streams[ID]; ok {
			ID++
		} else {
			return ID, nil
		}
	}
	return 0, errors.New("Failed to create new stream ID")
}

func (c *Controller) FetchStream(call schemas.Controller_fetchStream) error {
	ID := call.Params.Id()
	log.Info.Println("Fetching stream with ID", strconv.Itoa(int(ID)))

	if _, ok := c.streams[ID]; !ok {
		return errors.New("Stream ID not found: ", strconv.Itoa(int(ID)))
	}
	log.Info.Println("Stream ID", strconv.Itoa(int(ID)), "found")

	// Create a new capnp stream
	capnpStream, e := call.Results.NewStream()
	if e != nil {
		return e
	}

	// Populate it
	e = c.streams[ID].ToCapnp(&capnpStream)
	if e != nil {
		return e
	}

	return nil
}

func (c *Controller) DeleteStream(call schemas.Controller_deleteStream) error {
	ID := call.Params.Id()

	if _, ok := c.streams[ID]; !ok {
		return errors.New("Cannot delete stream: stream ID", strconv.Itoa(int(ID)), "not found")
	}

	delete(c.streams, ID)
	return nil
}

func saveStream(stream *stateless.Stream, capnpStream *schemas.Stream) error {
	e := stream.FromCapnp(capnpStream)
	if e != nil {
		return e
	}
	log.Info.Println("Preparing stream...", stream)
	return stream.ToBytes()
}

func (c *Controller) SaveStream(call schemas.Controller_saveStream) error {
	if !call.Params.HasStream() {
		return errors.New("No stream provided")
	}
	if !call.Params.HasStream() {
		return errors.New("Missing stream to save")
	}

	capnpStream, e := call.Params.Stream()
	if e != nil {
		return e
	}

	ID := capnpStream.Id()

	if ID == 0 {
		log.Info.Println("Creating a new stream")
		stream := stateless.Stream{}
		ID, e = c.newStreamID()
		if e != nil {
			return e
		}
		stream.ID = ID
		e = saveStream(&stream, &capnpStream)
		if e != nil {
			return e
		}
		c.streams[ID] = &stream
	} else {
		log.Info.Println("Update stream with ID", strconv.Itoa(int(ID)))
		if _, ok := c.streams[ID]; !ok {
			return errors.New("No stream with ID", strconv.Itoa(int(ID)), "found")
		}
		e = saveStream(c.streams[ID], &capnpStream)
		if e != nil {
			return e
		}
	}
	call.Results.SetId(ID)
	return nil
}
