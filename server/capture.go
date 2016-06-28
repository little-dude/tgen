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

type Capture struct {
	state    IOState
	stop     bool
	stats    Stats
	mtx      sync.RWMutex
	Packets  chan *RawPacket
	pktCount uint32
	port     string
}

type RawPacket struct {
	data []byte
	ci   gopacket.CaptureInfo
}

type Buffer []*RawPacket

func (c *Capture) Stop() {
	c.mtx.Lock()
	c.stop = true
	c.mtx.Unlock()
}

func (c *Capture) ShouldStop() bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.stop
}

func (c *Capture) State() IOState {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.state
}

func (c *Capture) SetState(newState IOState) {
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

func (c *Capture) Join(timeout uint32) error {
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

func NewCapture(port string) *Capture {
	c := Capture{
		Packets: make(chan *RawPacket, 1000),
		state:   NotStarted,
		port:    port,
	}
	return &c
}

func (c *Capture) Start(handle *pcap.Handle, pktCount uint32, bufferize bool) {
	c.pktCount = pktCount
	c.SetState(Started)
	if bufferize {
		buf := newRingBug(10000)
		go c.consumeChunks(buf)
		go c.captureChunks(buf, handle)
	} else {
		go c.capture(handle)
	}
}

func (c *Capture) WriteCapture(f *os.File) error {
	defer f.Close()

	defer func() {
		Info.Println("Finished writing capture file")
	}()

	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(65536, layers.LinkTypeEthernet)
	for pkt := range c.Packets {
		w.WritePacket(pkt.ci, pkt.data)
	}
	return nil
}

func (c *Capture) capture(handle *pcap.Handle) {
	defer c.SetState(Done)
	defer close(c.Packets)
	defer handle.Close()
	count := uint32(0)

	for {
		data, ci, e := handle.ReadPacketData()
		if e == nil {
			count++
			select {
			case c.Packets <- &RawPacket{data: data, ci: ci}:
			default:
			}
		}
		if !(count < c.pktCount || c.pktCount == 0) || c.ShouldStop() {
			c.Stop()
			break
		}
	}
	c.setStats(handle)
}

func (c *Capture) setStats(handle *pcap.Handle) error {
	stats, e := handle.Stats()
	if e != nil {
		return NewError("Failed to get capture stats:", e.Error())
	}

	c.mtx.Lock()
	c.stats = Stats{
		Received: uint32(stats.PacketsReceived),
		KDropped: uint32(stats.PacketsDropped),
		IDropped: uint32(stats.PacketsIfDropped),
	}
	c.mtx.Unlock()
	return nil
}

// kind of ring buffer implementation
type ring struct {
	In       chan Buffer
	Out      chan Buffer
	head     int
	tail     int
	len      int
	capacity int
	buff     []Buffer
}

func (r *ring) run() {
	defer close(r.Out)
loop:
	// give priority to writes, and only read from the ring buffer if there is
	// nothing to write
	for {
		select {
		case buf := <-r.In:
			if buf == nil {
				break loop
			}
			r.set(buf)
		default:
			select {
			case r.Out <- r.get():
			case <-time.After(10 * time.Millisecond):
			}
		}
	}
	for r.head >= r.tail {
		r.Out <- r.get()
	}
}

func (r *ring) set(v Buffer) {
	r.head = r.head + 1
	r.buff[r.head%r.len] = v
	if r.head-r.tail == r.len {
		r.resize(r.len * 4)
	}
}

func (r *ring) get() (v Buffer) {
	if r.head < r.tail {
		return nil
	}
	v = r.buff[r.tail%r.len]
	r.tail = r.tail + 1
	if r.len > r.capacity && r.head-r.tail <= r.len/8 {
		r.resize(r.len / 2)
	}
	return v
}

func (r *ring) resize(size int) {
	newb := make([]Buffer, size)
	for i := range newb {
		newb[i] = nil
	}
	r.buff = append(r.buff, newb...)
	r.len = size
	r.head = r.head % r.len
	r.tail = r.tail % r.len
}

func newRingBug(capacity int) *ring {
	r := ring{
		buff:     make([]Buffer, capacity),
		capacity: capacity,
		len:      capacity,
		head:     -1,
		tail:     0,
		In:       make(chan Buffer, 1000),
		Out:      make(chan Buffer, 1000),
	}
	go r.run()
	return &r
}

func (c *Capture) consumeChunks(ring *ring) {
	defer close(c.Packets)
	defer c.SetState(Done)
	for buf := range ring.Out {
		for i := 0; i < len(buf); i++ {
			c.Packets <- buf[i]
		}
	}
}

func (c *Capture) captureChunks(ring *ring, handle *pcap.Handle) {
	defer close(ring.In)
	defer handle.Close()

	count := uint32(0)

main:
	for c.pktCount == 0 || count < c.pktCount {
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
						case ring.In <- buf[:last+1]:
							count += uint32(last + 1)
							continue main
						default:
						}
					}
				}
			}
		}

		select {
		case ring.In <- buf[:last+1]:
			if c.ShouldStop() {
				break main
			}
		default:
			Error.Println("could not push buffer into ring... lost a buffer")
		}
		count += uint32(last + 1)
	}
	c.setStats(handle)
}
