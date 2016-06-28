package server

import (
	// "github.com/google/gopacket/pfring" FIXME: pf_ring does seem to work :(
	"github.com/google/gopacket/pcap"
	"github.com/little-dude/tgen/schemas"
	"os"
	"strconv"
	"time"
)

type empty struct{}

type Port struct {
	name       string
	controller *Controller
	isSending  bool
	sendDone   chan empty
	sendStop   chan empty
	sendError  chan error
	capture    *Capture
	interfaces []*Interface
}

func NewPort(name string, controller *Controller) *Port {
	capture := &Capture{}
	capture.SetState(NotStarted)
	return &Port{
		name:       name,
		controller: controller,

		isSending: false,
		sendDone:  make(chan empty, 1),
		sendStop:  make(chan empty, 1),
		sendError: make(chan error, 1),

		capture: capture,
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
		<-p.sendDone
		p.isSending = false
		return
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
	var ID uint16
outer:
	for i := 0; i < streamIDs.Len(); i++ {
		ID = streamIDs.At(i)
		for _, stream := range p.controller.streams {
			if stream.ID == ID {
				streams = append(streams, stream)
				continue outer
			}
		}
		return NewError("No stream found with ID", strconv.Itoa(int(ID)))
	}

	handle, e := pcap.OpenLive(p.name, 9999, true, -time.Millisecond*10)
	if e != nil {
		Error.Println("Failed to create the pcap handle:", e.Error())
		return NewError("Failed to create the pcap handle: ", e.Error())
	}

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

	p.isSending = true

	// wait a little bit to make sure the transmitting goroutine is running,
	// before returning.
	// FIXME: we could do this properly with a semaphore or something
	time.Sleep(time.Millisecond * 100)

	return nil
}

func (p *Port) WaitCapture(call schemas.Port_waitCapture) error {
	timeout := call.Params.Timeout()
	p.capture.Join(timeout)

	if p.capture.State() == Done {
		call.Results.SetDone(true)
	} else {
		call.Results.SetDone(false)
	}

	stats, _ := p.capture.Stats()
	call.Results.SetReceived(stats.Received)
	call.Results.SetDropped(stats.KDropped)

	return nil
}

func (p *Port) StopCapture(call schemas.Port_stopCapture) error {
	if p.capture.State() == Started {
		p.capture.Stop()
		p.capture.Join(0)
		return nil
	} else {
		return NewError(p.name, "is not capturing")
	}
}

func (p *Port) StartCapture(call schemas.Port_startCapture) error {
	if p.capture.State() == Started {
		return NewError(p.name, " is already capturing")
	}
	pktCount := call.Params.PacketCount()

	path, e := call.Params.File()
	if e != nil {
		return NewError(e.Error())
	}

	f, e := os.Create(path)
	if e != nil {
		return NewError("Could create capture file:", e.Error())
	}

	handle, e := pcap.OpenLive(p.name, 65635, true, time.Millisecond*10)
	if e != nil {
		return NewError("Could not create pcap handle:", e.Error())
	}

	p.capture = NewCapture(p.name)
	go p.capture.WriteCapture(f)
	p.capture.Start(handle, pktCount, true)

	Info.Println("capture started on", p.name)
	return nil
}
