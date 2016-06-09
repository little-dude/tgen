package server

import (
	// "github.com/google/gopacket/pfring" FIXME: pf_ring does seem to work :(
	"github.com/google/gopacket/pcap"
	"github.com/little-dude/tgen/schemas"
	"strconv"
	// "zombiezen.com/go/capnproto2"
)

type Port struct {
	name       string
	controller *Controller
}

func (p *Port) GetConfig(call schemas.Port_getConfig) error {
	capnpConfig, e := call.Results.NewConfig()
	if e != nil {
		return e
	}
	capnpConfig.SetName(p.name)
	return nil
}

func (p *Port) SetConfig(call schemas.Port_setConfig) error {
	return NewError("Not yet implemented ")
}

func createPcapHandle(portName string) (*pcap.Handle, error) {
	inactiveHandle, e := pcap.NewInactiveHandle(portName)
	defer inactiveHandle.CleanUp()
	if e != nil {
		return nil, e
	}
	inactiveHandle.SetPromisc(false)
	return inactiveHandle.Activate()
}

func (p *Port) StartSend(call schemas.Port_startSend) error {
	streamIDs, e := call.Params.Ids()
	if streamIDs.Len() == 0 {
		return NewError("No stream ID given")
	}

	streams := make([]Stream, 0)
	var streamFound bool
	var ID uint16
	for i := 0; i < streamIDs.Len(); i++ {
		ID = streamIDs.At(i)
		streamFound = false
		for _, stream := range p.controller.streams {
			if stream.ID == ID {
				streams = append(streams, stream)
				streamFound = true
				break
			}
		}
		if streamFound == false {
			return NewError("No stream found with ID", strconv.Itoa(int(ID)))
		}
	}

	handle, e := createPcapHandle(p.name)
	if e != nil {
		return NewError("Failed to create the pcap handle: ", e.Error())
	}

	for _, stream := range streams {
		for _, pkt := range stream.Packets {
			e = handle.WritePacketData(pkt)
			if e != nil {
				return NewError("Failed to write packet: ", e.Error())
			}
		}
	}
	return nil
}
