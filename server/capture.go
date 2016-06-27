package server

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"os"
	"sync"
	"time"
)

type RawPacket struct {
	data []byte
	ci   gopacket.CaptureInfo
}

type Buffer []*RawPacket

type CaptureState struct {
	Started bool
	Done    bool
}

var NotStarted = CaptureState{}
var Started = CaptureState{Started: true}
var Done = CaptureState{Done: true}

type Capture struct {
	state    CaptureState
	stop     bool
	stats    Stats
	mtx      sync.RWMutex
	cfast    chan Buffer
	cslow    chan Buffer
	pktCount uint32
}

type Stats struct {
	Received uint32
	KDropped uint32
	IDropped uint32
}

func (c *Capture) SetStop() {
	c.mtx.Lock()
	c.stop = true
	c.mtx.Unlock()
}

func (c *Capture) ShouldStop() bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.stop
}

func (c *Capture) State() CaptureState {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.state
}

func (c *Capture) SetState(newState CaptureState) {
	defer c.mtx.Unlock()
	c.mtx.Lock()
	c.state = newState
}

func (c *Capture) Stats() (Stats, error) {
	state := c.State()
	if state == Done {
		c.mtx.RLock()
		defer c.mtx.RUnlock()
		return c.stats, nil
	}
	if state == NotStarted {
		return Stats{}, NewError("no capture occured")
	}
	if state == Started {
		return Stats{}, NewError("still capturing")
	}
	panic("Unknown state")
}

func (c *Capture) Wait(timeout uint32) error {
	if c.State() == NotStarted {
		return NewError("no capture to wait for")
	}
	start := time.Now()
	t := time.Millisecond * time.Duration(timeout)
	for time.Now().Sub(start) < t || timeout == 0 {
		if c.State() == Done {
			return nil
		}
		time.Sleep(time.Millisecond * 50)
	}
	return NewError("capture did not finish")
}

func NewCapture(captureFile string, port string, pktCount uint32) (*Capture, error) {

	f, e := os.Create(captureFile)
	if e != nil {
		return &Capture{}, NewError("Could not create capture file:", e.Error())
	}

	handle, e := pcap.OpenLive(port, 65635, true, time.Millisecond*10)
	if e != nil {
		return &Capture{}, NewError("Could not create pcap handle:", e.Error())
	}

	c := Capture{
		cfast:    make(chan Buffer, 1000),
		cslow:    make(chan Buffer, 1),
		pktCount: pktCount,
		state:    NotStarted,
	}

	c.SetState(Started)
	go c.WriteCapture(f)
	go c.BufferCapture()
	go c.Capture(handle)

	return &c, nil
}

func (c *Capture) WriteCapture(f *os.File) {

	defer func() {
		f.Close()
		Info.Println("Finished writing capture file")
		c.SetState(Done)
	}()

	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(65536, layers.LinkTypeEthernet)
	for {
		buf, ok := <-c.cslow
		if !ok {
			return
		}
		for i := 0; i < len(buf); i++ {
			w.WritePacket(buf[i].ci, buf[i].data)
		}
	}
}

func (c *Capture) BufferCapture() {
	defer close(c.cslow)

	count := uint32(0)
	buffers := make([]Buffer, 0, 1000)

main:
	for {
		select {
		case buf := <-c.cfast:
			if buf == nil {
				break main
			}
			count += uint32(len(buf))
			if count >= c.pktCount && c.pktCount > 0 {
				Info.Println("received", count, "packets: stopping capture")
				buffers = append(buffers, buf[:c.pktCount-(count-uint32(len(buf)))])
				c.SetStop()
				break main
			}
			buffers = append(buffers, buf)
			if cap(buffers) == 0 {
				more := make([]Buffer, len(buffers), 2*len(buffers))
				copy(more, buffers)
				buffers = more
			}
		default:
			if len(buffers) > 0 {
				select {
				case c.cslow <- buffers[0]:
					buffers = buffers[1:]
				case <-time.After(time.Millisecond * 20):
				}
			}
		}

	}

	Info.Println("waiting for capture file to be written")
	for i := 0; i < len(buffers); i++ {
		c.cslow <- buffers[i]
	}
}

func (c *Capture) Capture(handle *pcap.Handle) {
	defer close(c.cfast)
	defer handle.Close()

main:
	for {
		buf := make(Buffer, 1000)
		last := -1
		for i := 0; i < 1000; i++ {
			for {
				data, ci, e := handle.ReadPacketData()
				if e == nil {
					buf[i] = &RawPacket{data: data, ci: ci}
					last = i
					break
				} else if e == pcap.NextErrorTimeoutExpired {
					if c.ShouldStop() {
						break main
					}
					if last >= 0 {
						select {
						case c.cfast <- buf[:last+1]:
						default:
						}
					}
				}
			}
		}

		select {
		case c.cfast <- buf[:last+1]:
			if c.ShouldStop() {
				break main
			}
		default:
		}
	}

	stats, e := handle.Stats()
	if e != nil {
		Error.Println("Failed to get capture stats:", e.Error())
		return
	}

	c.mtx.Lock()
	c.stats = Stats{
		Received: uint32(stats.PacketsReceived),
		KDropped: uint32(stats.PacketsDropped),
		IDropped: uint32(stats.PacketsIfDropped),
	}
	c.mtx.Unlock()
}
