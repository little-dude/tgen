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
	rx         *Rx
	tx         *Tx
	interfaces []*Interface
}

func NewPort(name string, controller *Controller) *Port {
	return &Port{
		name:       name,
		controller: controller,
		tx:         NewTx(),
		rx:         NewRx(),
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
	e := p.tx.state.WaitDone(timeout)
	if e == nil {
		call.Results.SetDone(true)
	} else {
		call.Results.SetDone(false)
	}
	return nil
}

func (p *Port) StartSend(call schemas.Port_startSend) error {
	streamIDs, e := call.Params.Ids()
	if streamIDs.Len() == 0 {
		return NewError("No stream ID given")
	}

	streams := make([]*Stream, streamIDs.Len())
	for i := 0; i < streamIDs.Len(); i++ {
		if stream, ok := p.controller.streams[streamIDs.At(i)]; ok {
			streams[i] = stream
		} else {
			return NewError("Stream ID not found: ", strconv.Itoa(int(streamIDs.At(i))))
		}
	}

	if p.tx.state.Active() {
		return NewError("already transmitting")
	}

	p.tx = NewTx()

	handle, e := pcap.OpenLive(p.name, 65635, true, time.Millisecond*10)
	if e != nil {
		Error.Println("Failed to create the pcap handle:", e.Error())
		return NewError("Failed to create the pcap handle: ", e.Error())
	}
	p.tx.state.SetRun()
	go p.tx.Start(handle, streams)
	return nil
}

func (p *Port) WaitCapture(call schemas.Port_waitCapture) error {
	timeout := call.Params.Timeout()
	e := p.rx.state.WaitDone(timeout)
	if e == nil {
		call.Results.SetDone(true)
	} else {
		call.Results.SetDone(false)
	}

	stats, _ := p.rx.Stats()
	call.Results.SetReceived(stats.Received)
	call.Results.SetDropped(stats.KDropped)

	return nil
}

func (p *Port) StopCapture(call schemas.Port_stopCapture) error {
	if p.rx.state.Active() {
		p.rx.state.SetStop()
		return p.rx.state.WaitDone(0)
	} else {
		return NewError(p.name, "is not capturing")
	}
}

func (p *Port) StartCapture(call schemas.Port_startCapture) error {
	if p.rx.state.Active() {
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

	p.rx = NewRx()
	go p.rx.Save(f)
	p.rx.Start(handle, pktCount, true)

	Info.Println("capture started on", p.name)
	return nil
}
