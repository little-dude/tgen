package server

import (
	"github.com/little-dude/tgen/schemas"
	"net"
	"strconv"
	"zombiezen.com/go/capnproto2"
)

// Controller represent the controller running on the host.
type Controller struct {
	ports   []Port
	streams []Stream
}

// GetPorts is a local implementation of the capnproto capability.
func (c *Controller) GetPorts(call schemas.Controller_getPorts) error {
	// retrieve the ports
	e := c.getHostPorts()
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
	for i, _ := range c.ports {
		port := schemas.Port_ServerToClient(&c.ports[i])
		ptr := capnp.NewInterface(seg, seg.Message().AddCap(port.Client)).ToPtr()
		portsList.SetPtr(i, ptr)
	}
	return nil
}

func (c *Controller) getHostPorts() error {
	Info.Println("Listing local ports")
	itfs, e := net.Interfaces()
	if e != nil {
		Error.Println(e.Error())
		return NewError("Failed to list the interfaces:", e.Error())
	}

	c.ports = make([]Port, len(itfs))
	for i, itf := range itfs {
		c.ports[i] = Port{name: itf.Name, controller: c}
		Info.Println("Found port:", itf.Name)
	}
	return nil
}

// ListStreams is the implementation of the capnproto ListStreams method
func (c *Controller) ListStreams(call schemas.Controller_listStreams) error {
	streamIDs, e := call.Results.NewIds(int32(len(c.streams)))
	if e != nil {
		return e
	}

	for i, stream := range c.streams {
		streamIDs.Set(i, stream.ID)
	}
	return nil
}

func (c *Controller) newStreamID() (uint16, error) {
	id := uint16(1)
outer:
	for true {
		for _, stream := range c.streams {
			if stream.ID == id {
				id++
				continue outer
			}
		}
		return id, nil
	}
	return 0, NewError("Failed to create new stream ID")
}

func (c *Controller) FetchStream(call schemas.Controller_fetchStream) error {
	streamID := call.Params.Id()
	Info.Println("Fetching stream with ID", strconv.Itoa(int(streamID)))
	for _, stream := range c.streams {
		if stream.ID == streamID { // We found the stream we need to return.
			Info.Println("Stream ID", strconv.Itoa(int(streamID)), "found")
			// Create a new capnp stream
			capnpStream, e := call.Results.NewStream()
			if e != nil {
				return e
			}
			// Populate it
			e = stream.ToCapnp(&capnpStream)
			if e != nil {
				return e
			}
			return nil
		}
	}
	// We did not find the stream with the corresponding stream ID, return an error
	return NewError("Stream ID not found: ", strconv.Itoa(int(streamID)))
}

func (c *Controller) DeleteStream(call schemas.Controller_deleteStream) error {
	streamID := call.Params.Id()
	for i, stream := range c.streams {
		if stream.ID == streamID {
			c.streams = append(c.streams[:i], c.streams[i+1:]...)
			return nil
		}
	}
	return NewError("Cannot delete stream: stream ID", strconv.Itoa(int(streamID)), "not found")
}

func saveStream(stream *Stream, capnpStream *schemas.Stream) error {
	e := stream.FromCapnp(capnpStream)
	if e != nil {
		return e
	}
	Info.Println("Preparing stream...", stream)
	return stream.ToBytes()
}

func (c *Controller) getStream(ID uint16) (*Stream, error) {
	for _, stream := range c.streams {
		if stream.ID == ID {
			return &stream, nil
		}
	}
	return nil, NewError("No stream with ID", strconv.Itoa(int(ID)), "found")
}

func (c *Controller) SaveStream(call schemas.Controller_saveStream) error {
	if !call.Params.HasStream() {
		return NewError("No stream provided")
	}
	if !call.Params.HasStream() {
		return NewError("Missing stream to save")
	}

	capnpStream, e := call.Params.Stream()
	if e != nil {
		return e
	}

	streamID := capnpStream.Id()

	if streamID == 0 {
		stream := Stream{}
		streamID, e = c.newStreamID()
		if e != nil {
			return e
		}
		stream.ID = streamID
		e = saveStream(&stream, &capnpStream)
		if e != nil {
			return e
		}
		c.streams = append(c.streams, stream)
	} else {
		stream, e := c.getStream(streamID)
		e = saveStream(stream, &capnpStream)
		if e != nil {
			return e
		}
	}
	call.Results.SetId(streamID)
	return nil
}
