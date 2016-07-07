package server

import (
	// "github.com/google/gopacket/pfring" FIXME: pf_ring does seem to work :(
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"github.com/little-dude/tgen/schemas"
	"os"
	"strconv"
	"sync"
	"time"
	capnp "zombiezen.com/go/capnproto2"
)

type empty struct{}

type Port struct {
	name          string
	controller    *Controller
	rx            *Rx
	tx            *Tx
	lans          []*LAN
	capturingLock sync.RWMutex
	capturing     bool
}

func NewPort(name string, controller *Controller) *Port {
	return &Port{
		name:       name,
		controller: controller,
		rx:         NewRx(name),
		tx:         NewTx(name),
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
	if e != nil {
		return NewError(e.Error())
	}
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

	p.tx = NewTx(p.name)
	p.tx.SendStreams(streams)
	return nil
}

func (p *Port) isCapturing() bool {
	p.capturingLock.RLock()
	defer p.capturingLock.RUnlock()
	return p.capturing
}

func (p *Port) waitCapture(timeout uint32) bool {
	start := time.Now()
	e := p.rx.state.WaitDone(timeout)
	if e != nil {
		return false
	} else {
		t := time.Millisecond * time.Duration(timeout)
		for time.Now().Sub(start) < t || timeout == 0 {
			if !p.isCapturing() {
				return true
			} else {
				time.Sleep(time.Millisecond * 50)
			}
		}
		return false
	}
}

func (p *Port) WaitCapture(call schemas.Port_waitCapture) error {
	if p.waitCapture(call.Params.Timeout()) {
		call.Results.SetDone(true)
		stats, _ := p.rx.Stats()
		call.Results.SetReceived(stats.Received)
		call.Results.SetDropped(stats.KDropped)
	} else {
		call.Results.SetDone(false)
	}
	return nil
}

func (p *Port) StopCapture(call schemas.Port_stopCapture) error {
	if p.isCapturing() {
		if p.rx.state.Active() {
			p.rx.state.SetStop()
		}
		p.waitCapture(0)
		return nil
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

	p.rx = NewRx(p.name)

	f, e := os.Create(path)
	if e != nil {
		return NewError("Could create capture file:", e.Error())
	}

	p.capturingLock.Lock()
	p.capturing = true
	p.capturingLock.Unlock()

	captureBuf := NewRingBuf(1000)

	go func(f *os.File, chunks <-chan []*RawPacket) {
		defer func() {
			f.Close()
			Info.Println("Finished writing capture file")
			p.capturingLock.Lock()
			p.capturing = false
			p.capturingLock.Unlock()
		}()
		w := pcapgo.NewWriter(f)
		w.WriteFileHeader(65536, layers.LinkTypeEthernet)
		for chunk := range chunks {
			for i := 0; i < len(chunk); i++ {
				w.WritePacket(chunk[i].ci, chunk[i].data)
			}
		}
	}(f, captureBuf.Out)

	e = p.rx.CaptureChunks(captureBuf.In, pktCount, pcap.DirectionInOut, "")
	if e != nil {
		captureBuf.Close()
		return NewError("Failed to start capture:", e.Error())
	}

	Info.Println("capture started on", p.name)
	return nil
}

func (p *Port) AddLan(call schemas.Port_addLan) error {
	cidr, e := call.Params.Cidr()
	if e != nil {
		return e
	}
	lan, e := NewLAN(p.name, cidr, []uint32{})
	if e != nil {
		return e
	}
	p.lans = append(p.lans, lan)
	return call.Results.SetLan(schemas.Lan_ServerToClient(lan))
}

func (p *Port) GetLans(call schemas.Port_getLans) error {
	// initialize a list of capnp interfaces
	lans, e := call.Results.NewLans(int32(len(p.lans)))
	if e != nil {
		return e
	}

	// populate the list
	seg := call.Results.Segment()
	i := 0
	for _, lan := range p.lans {
		// MAGIC!
		e := lans.SetPtr(i, capnp.NewInterface(seg, seg.Message().AddCap(schemas.Lan_ServerToClient(lan).Client)).ToPtr())
		if e != nil {
			return e
		}
		i++
	}
	return nil
}

func (p *Port) DeleteLan(call schemas.Port_deleteLan) error {
	return NewError("Not implemented")
}
