package server

import (
	// "github.com/google/gopacket/pfring" FIXME: pf_ring does seem to work :(
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
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
	// lans       []*LAN
}

func NewPort(name string, controller *Controller) *Port {
	return &Port{
		name:       name,
		controller: controller,
		rx:         NewRx(name, pcap.DirectionInOut, true),
		tx:         NewTx(),
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

	p.rx = NewRx(p.name, pcap.DirectionInOut, true)

	f, e := os.Create(path)
	if e != nil {
		return NewError("Could create capture file:", e.Error())
	}

	go func(f *os.File, packets <-chan *RawPacket) {
		defer f.Close()
		defer func() { Info.Println("Finished writing capture file") }()
		w := pcapgo.NewWriter(f)
		w.WriteFileHeader(65536, layers.LinkTypeEthernet)
		for pkt := range packets {
			w.WritePacket(pkt.ci, pkt.data)
		}
	}(f, p.rx.Packets)

	e = p.rx.Start(pktCount)
	if e != nil {
		// at this point, the goroutine started by rx.Save() is reading on
		// p.rx.Packets to force it to stop, we close the channel

		// FIXME: maybe the "Save()" function should not be a method of `Rx`,
		// it would make it clearer to have a consumer of p.rx.Packets.
		close(p.rx.Packets)
		return NewError("Failed to start capture:", e.Error())
	}

	Info.Println("capture started on", p.name)
	return nil
}
