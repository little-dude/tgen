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
		sendDone:     make(chan empty, 1),
		sendStop:     make(chan empty, 1),
		sendError:    make(chan error, 1),
		isCapturing:  false,
		captureDone:  make(chan empty, 1),
		captureStop:  make(chan empty, 1),
		captureError: make(chan error, 1),
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
	if !p.isCapturing {
		call.Results.SetDone(true)
		call.Results.SetError("")
		return nil
	}
	p.waitCapture(timeout)
	select {
	case e := <-p.captureError:
		call.Results.SetError(e.Error())
	default:
		call.Results.SetError("")
	}
	call.Results.SetDone(!p.isCapturing)
	return nil
}

func (p *Port) StopCapture(call schemas.Port_stopCapture) error {
	if !p.isCapturing {
		return NewError(p.name, "is not capturing")
	}
	select {
	case p.captureStop <- empty{}:
		Trace.Println("Sending signal to stop capture")
		return nil
	default:
		return NewError("Could not send signal to stop capture")
	}
	// case <-time.After(time.Millisecond * time.Duration(100)):
	// }
}

func (p *Port) newCaptureError(msg string, e error) {
	select {
	case p.captureError <- e:
	default:
	}
	Error.Println(msg, ":", e.Error())
}

func (p *Port) StartCapture(call schemas.Port_startCapture) error {
	if p.isCapturing {
		return NewError(p.name, " is already capturing")
	}

	path, e := call.Params.SavePath()
	if e != nil {
		return e
	}
	snapshotLen := call.Params.SnapshotLength()
	// FIXME: check for overflow
	timeout := time.Duration(time.Second * time.Duration(call.Params.Timeout()))
	packetCount := call.Params.PacketCount()
	promiscuous := call.Params.Promiscuous()

	p.captureStop = make(chan empty, 1)
	p.captureError = make(chan error, 1)
	started := make(chan empty, 1)
	go p.capture(path, snapshotLen, timeout, packetCount, promiscuous, started)

	select {
	case <-started:
		p.isCapturing = true
		Info.Println("capture started on", p.name)
		return nil
	case <-time.After(3 * time.Second):
		return NewError("Capture did not start after 3 seconds")
	}
	return NewError("Unknown error")
}

func (p *Port) capture(path string, snapshotLen uint32, timeout time.Duration, packetCount uint32, promiscuous bool, started chan empty) {

	defer func() {
		select {
		case p.captureDone <- empty{}:
		default:
		}
	}()

	Trace.Println("Creating capture file", path)
	f, e := os.Create(path)
	if e != nil {
		p.newCaptureError("Could not create capture file", e)
		return
	}

	defer func() {
		e := f.Close()
		if e != nil {
			p.newCaptureError("Could not close capture file properly", e)
		}
	}()

	writer := pcapgo.NewWriter(f)
	writer.WriteFileHeader(snapshotLen, layers.LinkTypeEthernet)

	// FIXME: snapshotLen should be int32 or uint32?
	Trace.Println("Creating handle for capture")
	handle, e := pcap.OpenLive(p.name, int32(snapshotLen), promiscuous, timeout)
	if e != nil {
		p.newCaptureError("Could not create capture file", e)
		return
	}
	defer func() {
		// FIXME: this sometime hangs. As a result we cannot close the handle
		// properly for now.
		// handle.Close()
	}()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	count := uint32(0)
	var packet gopacket.Packet
	started <- empty{}

	for {
		select {
		case <-p.captureStop:
			Trace.Println("Received signal to stop capture")
			return
		default:
			select {
			case packet = <-packetSource.Packets():
				count++
				e = writer.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
				if e != nil {
					p.newCaptureError("Failed to write captured packet", e)
				}
				if packetCount > uint32(0) && packetCount == count {
					Info.Println("Received enough packets, sotpping capture")
					return
				}
			case <-time.After(time.Millisecond * time.Duration(100)):
				continue
			}
		}
	}
}
