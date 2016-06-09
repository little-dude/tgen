package server

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/little-dude/tgen/schemas"
)

type Stream struct {
	ID            uint16
	Count         uint32
	PacketsPerSec uint32
	Layers        []Layer
	Packets       [][]byte
}

func (s *Stream) String() string {
	return fmt.Sprint(
		"Count: ", s.Count,
		", PacketsPerSec: ", s.PacketsPerSec,
		", ID: ", s.ID,
		", Layers: ", s.Layers)
}

func (s *Stream) ToCapnp(capnpStream *schemas.Stream) error {
	capnpStream.SetCount(s.Count)
	capnpStream.SetPacketsPerSec(s.PacketsPerSec)
	capnpStream.SetId(s.ID)

	// initialize a list of layers
	capnpLayers, e := capnpStream.NewLayers(int32(len(s.Layers)))
	if e != nil {
		return e
	}
	seg := capnpLayers.Segment()

	// populate the list
	for i, l := range s.Layers {

		// create a new capnp layer struct
		capnpLayer, e := schemas.NewProtocol(seg)
		if e != nil {
			return e
		}

		// populate it
		l.ToCapnp(&capnpLayer)

		// add it to the list
		e = capnpLayers.Set(i, capnpLayer)
		if e != nil {
			return e
		}
	}

	return nil
}

func (s *Stream) FromCapnp(capnpStream *schemas.Stream) error {
	s.PacketsPerSec = capnpStream.PacketsPerSec()
	s.Count = capnpStream.Count()

	// populate the layers
	capnpLayers, e := capnpStream.Layers()
	if e != nil {
		return e
	}

	s.Layers = make([]Layer, int32(capnpLayers.Len()))
	var capnpLayer schemas.Protocol
	for i := 0; i < capnpLayers.Len(); i++ {
		capnpLayer = capnpLayers.At(i)
		s.Layers[i], e = NewLayer(&capnpLayer)
		if e != nil {
			return e
		}
	}
	return nil
}

func (s *Stream) ToBytes() error {
	if len(s.Layers) == 0 {
		return NewError("No layers configured for stream", s.String())
	}
	layers := make([][]gopacket.SerializableLayer, len(s.Layers))
	for i, l := range s.Layers {
		layers[i] = l.ToPackets()
	}
	opts := gopacket.SerializeOptions{}
	var buf gopacket.SerializeBuffer
	if s.Count == 0 {
		return NewError("packet count is 0: cannot create stream data")
	}
	s.Packets = make([][]byte, s.Count)
	for i := uint32(0); i < s.Count; i++ {
		buf = gopacket.NewSerializeBuffer()
		for j := len(layers) - 1; j >= 0; j-- {
			layers[j][i%uint32(len(layers[j]))].SerializeTo(buf, opts)
		}
		s.Packets[i] = buf.Bytes()
	}
	return nil
}
