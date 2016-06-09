package server

import (
	// "github.com/google/gopacket/pfring" FIXME: pf_ring does seem to work :(
	"github.com/google/gopacket/pcap"
	"github.com/little-dude/tgen/schemas"
	// "zombiezen.com/go/capnproto2"
)

type Port struct {
	name string
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
	return nil
}

func createPcapHandle(portName string) (*pcap.Handle, error) {
	inactiveHandle, err := pcap.NewInactiveHandle(portName)
	defer inactiveHandle.CleanUp()
	if err != nil {
		return nil, err
	}
	inactiveHandle.SetPromisc(false)
	return inactiveHandle.Activate()
}

func (p *Port) StartSend(call schemas.Port_startSend) error {
	// handle, e := createPcapHandle(p.name)
	// if e != nil {
	// 	return NewError(INTERNAL_ERROR, "Failed to create the pcap handle: ", e.Error()))
	// }
	// for _, stream := range port.streams {
	// 	for _, pkt := range stream.Packets {
	// 		e = handle.WritePacketData(pkt)
	// 		if e != nil {
	// 			return NewError(INTERNAL_ERROR, "Failed to write packet: ", e.Error()))
	// 		}
	// 	}
	// }
	return nil
}
