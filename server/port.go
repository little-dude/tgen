package server

import (
	// "github.com/google/gopacket/pfring" FIXME: pf_ring does seem to work :(
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"github.com/little-dude/tgen/schemas"
	"os"
	"strconv"
	"time"
	// "zombiezen.com/go/capnproto2"
)

type empty struct{}

type RawPacket struct {
	data []byte
	ci   gopacket.CaptureInfo
}

type Port struct {
	name        string
	controller  *Controller
	isSending   bool
	sendDone    chan empty
	sendStop    chan empty
	sendError   chan error
	isCapturing bool
	captureDone chan empty
	captureStop chan empty
}

func NewPort(name string, controller *Controller) *Port {
	return &Port{
		name:        name,
		controller:  controller,
		isSending:   false,
		sendDone:    make(chan empty, 1),
		sendStop:    make(chan empty, 1),
		sendError:   make(chan error, 1),
		isCapturing: false,
		captureDone: make(chan empty, 1),
		captureStop: make(chan empty, 1),
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
	// func OpenLive(device string, snaplen int32, promisc bool, timeout time.Duration) (handle *Handle, _ error)
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

func (p *Port) waitCapture(timeout uint32) {
	if !p.isCapturing {
		return
	}
	if timeout == 0 {
		<-p.captureDone
		p.isCapturing = false
		return
	}
	select {
	case <-p.captureDone:
		p.isCapturing = false
	case <-time.After(time.Millisecond * time.Duration(timeout)):
	}
}

func (p *Port) WaitCapture(call schemas.Port_waitCapture) error {
	timeout := call.Params.Timeout()
	p.waitCapture(timeout)
	call.Results.SetDone(!p.isCapturing)
	return nil
}

func (p *Port) StopCapture(call schemas.Port_stopCapture) error {
	if !p.isCapturing {
		return NewError(p.name, "is not capturing")
	}
	select {
	case p.captureStop <- empty{}:
		Trace.Println("Waiting for capture to finish")
		p.waitCapture(0)
		return nil
	default:
		return NewError("Could not send signal to stop capture")
	}
}

func (p *Port) StartCapture(call schemas.Port_startCapture) error {
	if p.isCapturing {
		return NewError(p.name, " is already capturing")
	}

	packetCount := call.Params.PacketCount()
	path, e := call.Params.FilePath()
	if e != nil {
		return NewError(e.Error())
	}

	handle, e := pcap.OpenLive(p.name, 65635, true, time.Millisecond*10)
	if e != nil {
		return NewError("Could not create pcap handle:", e.Error())
	}

	f, e := os.Create(path)
	if e != nil {
		return NewError("Could create capture file:", e.Error())
	}

	p.captureStop = make(chan empty, 1)
	c := make(chan []*RawPacket)

	Info.Println("Starting capture")
	p.isCapturing = true
	go p.capture(c, handle, packetCount)
	go p.saveCapture(c, f)
	return nil
}

func (p *Port) saveCapture(c chan []*RawPacket, f *os.File) {
	defer func() {
		p.captureDone <- empty{} // signal that the capture is finished
		f.Close()
		Info.Println("Finished writing capture file")
	}()
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(65536, layers.LinkTypeEthernet)
	for {
		buf, ok := <-c // block until the capturing goroutine sends a buffer to read
		if !ok {
			return
		}

		for i := 0; i < len(buf); i++ {
			w.WritePacket(buf[i].ci, buf[i].data)
		}
	}
}

func sendBuf(buffers [][]*RawPacket, c chan []*RawPacket) (remaining [][]*RawPacket) {
	// TODO: investigate why when writing on tmpfs we have better perf. One
	// idea it that `select` is slower when the chan cannot receive data, which
	// happens more often when writing is slow.
	select {
	case c <- buffers[0]:
		buffers = buffers[1:]
		return buffers
	default:
		return buffers
	}
}

func minPkt(count, pktCount uint32) uint32 {
	if pktCount == 0 {
		return 1000
	}
	if pktCount-count > 1000 {
		return 1000
	}
	return pktCount - count
}

func (p *Port) capture(c chan []*RawPacket, handle *pcap.Handle, pktCount uint32) {
	defer handle.Close()
	defer close(c)

	// We can buffer up to 1M packets. With an average packet size of 576kB
	// packet size, that's about 550M which seems big enough.
	buffers := make([][]*RawPacket, 0, 1000) // 1k buffers
	count := uint32(0)
main:
	for count < pktCount || pktCount == 0 {
		bufSize := minPkt(count, pktCount)
		buf := make([]*RawPacket, bufSize)
		for i := uint32(0); i < bufSize; i++ {
			for {
				data, ci, e := handle.ReadPacketData()
				if e == nil {
					buf[i] = &RawPacket{data: data, ci: ci}
					count++
					break
				}
				// check if we should stop the capture
				select {
				case <-p.captureStop:
					break main
				default:
				}
			}
		}

		buffers = append(buffers, buf)
		buffers = sendBuf(buffers, c)
		if cap(buffers) == 0 {
			// we filled all our buffers already. we need to allocate more
			// memory, and copy over all the buffers which is not great.
			more := make([][]*RawPacket, len(buffers), 2*len(buffers))
			copy(more, buffers)
			buffers = more
		}
	}

	// wait for all the buffers to be consumed. this can take a while
	Info.Println("stopping capturing", count, "packets. waiting capture file to be written")
	for i := 0; i < len(buffers); i++ {
		c <- buffers[i] // if we have a lot of buffers, this can take very long
	}
	Info.Println("done")
}
