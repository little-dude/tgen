package server

import (
	// "github.com/google/gopacket/pfring" FIXME: pf_ring does seem to work :(
	"github.com/google/gopacket/pcap"
	"github.com/little-dude/tgen/schemas"
	"strconv"
	"time"
	// "zombiezen.com/go/capnproto2"
)

type empty struct{}

type Port struct {
	name         string
	controller   *Controller
	isSending    bool
	sendDone     chan empty // semaphore used by transmitting goroutine to tell the main goroutine that it finished transmitting.
	sendStop     chan empty // semaphore used to tell transmitting goroutine to stop.
	sendError    chan error // channel used to send the first error a transmitting goroutine encounters.
	isCapturing  bool
	captureDone  chan empty
	captureStop  chan empty
	captureError chan error
}

func NewPort(name string, controller *Controller) *Port {
	return &Port{
		name:         name,
		controller:   controller,
		isSending:    false,
		sendDone:     make(chan empty),
		sendStop:     make(chan empty),
		sendError:    make(chan error),
		isCapturing:  false,
		captureDone:  make(chan empty),
		captureStop:  make(chan empty),
		captureError: make(chan error),
	}
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

func (p *Port) WaitSend(call schemas.Port_waitSend) error {
	timeout := call.Params.Timeout()
	p.waitSend(timeout)
	select {
	case e := <-p.sendError:
		call.Results.SetError(e.Error())
	default:
		call.Results.SetError("")
	}
	call.Results.SetDone(!p.isSending)
	return nil
}

func (p *Port) waitSend(timeout uint32) {
	if timeout == 0 {
		timeout = 1
	}
	select {
	case <-p.sendDone:
		p.isSending = false
	case <-time.After(time.Millisecond * time.Duration(timeout)):
	}
}

func (p *Port) StartSend(call schemas.Port_startSend) error {

	if p.isSending {
		return NewError(p.name, " is already transmitting")
	}

	streamIDs, e := call.Params.Ids()
	if streamIDs.Len() == 0 {
		return NewError("No stream ID given")
	}

	streams := make([]*Stream, 0)
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

	Trace.Println("Creating pcap handle on", p)
	handle, e := createPcapHandle(p.name)
	if e != nil {
		Error.Println("Failed to create the pcap handle:", e.Error())
		return NewError("Failed to create the pcap handle: ", e.Error())
	}

	p.isSending = true

	go func() {

		defer func() {
			p.sendDone <- empty{}
			Info.Println("Done sending on port", p)
		}()

	outer:
		for _, stream := range streams {
			Info.Println("Starting to send stream", stream.ID)

			for _, pkt := range stream.Packets {
				select {
				case <-p.sendStop:
					break outer
				default:
					e = handle.WritePacketData(pkt)
					if e != nil {
						select {
						case p.sendError <- e:
						default:
							Error.Println("Failed to write packet: ", e.Error())
						}
					}
				}
			}
		}
	}()

	return nil
}
