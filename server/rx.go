package server

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"os"
)

type PcapStats struct {
	Received uint32
	KDropped uint32
	IDropped uint32
}

type Rx struct {
	state    *RxTxState
	stats    PcapStats
	Packets  chan *RawPacket
	pktCount uint32
}

type RawPacket struct {
	data []byte
	ci   gopacket.CaptureInfo
}

type Buffer []*RawPacket

func (rx *Rx) Stats() (PcapStats, error) {
	if rx.state.Done() {
		// no need to lock to access the stats here since when capture is done,
		// no goroutine should access the stats anymore
		return rx.stats, nil
	}
	if rx.state.Inactive() {
		return PcapStats{}, NewError("capture did not start yet")
	}
	if rx.state.Active() {
		return PcapStats{}, NewError("still capturing")
	}
	panic("Unknown state")
}

func NewRx() *Rx {
	c := Rx{
		Packets: make(chan *RawPacket, 1000),
		state:   NewRxTxState(),
	}
	return &c
}

func (rx *Rx) Start(handle *pcap.Handle, pktCount uint32, bufferize bool) {
	rx.pktCount = pktCount
	rx.state.SetRun()
	if bufferize {
		buf := newRingBuf(10000)
		go rx.consumeChunks(buf)
		go rx.captureChunks(buf, handle)
	} else {
		go rx.capture(handle)
	}
}

func (rx *Rx) Save(f *os.File) error {
	defer f.Close()

	defer func() {
		Info.Println("Finished writing capture file")
	}()

	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(65536, layers.LinkTypeEthernet)
	for pkt := range rx.Packets {
		w.WritePacket(pkt.ci, pkt.data)
	}
	return nil
}

func (rx *Rx) capture(handle *pcap.Handle) {
	defer rx.state.SetDone()
	defer close(rx.Packets)
	defer handle.Close()
	count := uint32(0)

	for {
		data, ci, e := handle.ReadPacketData()
		if e == nil {
			count++
			// WARNING: this can block forever is nothing is consuming the incoming packets
			rx.Packets <- &RawPacket{data: data, ci: ci}
		}
		if !(count < rx.pktCount || rx.pktCount == 0) || rx.state.Stopping() {
			break
		}
	}
	rx.setStats(handle)
}

func (rx *Rx) setStats(handle *pcap.Handle) error {
	stats, e := handle.Stats()
	if e != nil {
		return NewError("Failed to get capture stats:", e.Error())
	}

	// no need to lock here, since the main goroutine only reads the stats when
	// the state of the capture is "Done", which occurs when this goroutine
	// exits
	rx.stats = PcapStats{
		Received: uint32(stats.PacketsReceived),
		KDropped: uint32(stats.PacketsDropped),
		IDropped: uint32(stats.PacketsIfDropped),
	}
	return nil
}

// kind of ring buffer implementation
// inspired by https://github.com/zfjagann/golang-ring/blob/master/ring.go
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
	for {
		// give priority to writes, and only read from the ring buffer if there
		// is nothing to write
		select {
		case buf := <-r.In: // data incoming: add it to the buffer and continue
			if buf == nil { // r.In is closed, exit
				break loop
			}
			r.set(buf)
		default: // no data incoming, let see if there is something to read and if someone wants to read it
			if r.tail < r.head { // there is something to read
				select {
				case r.Out <- r.peek(): // a goroutine is reading from r.Out
					r.get()
				default: // nobody is reading from r.Out, do nothing and continue
				}
			} else { // nothing to read, the buffer is empty
				// wait for data to come, so that we do not consume CPU
				buf := <-r.In
				if buf == nil { // r.In is closed, exit
					break loop
				}
				r.set(buf)
			}
		}
	}
	for r.head >= r.tail {
		r.Out <- r.get()
	}
}

func (r *ring) peek() (v Buffer) {
	return r.buff[r.tail%r.len]
}

func (r *ring) set(v Buffer) {
	if r.head-r.tail == r.len-1 {
		r.resize(r.len * 4)
	}
	r.head = r.head + 1
	r.buff[r.head%r.len] = v
}

func (r *ring) get() (v Buffer) {
	v = r.buff[r.tail%r.len]
	r.tail = r.tail + 1
	// shrinking is expensive for big buffers, we don't do it too often
	if r.len > r.capacity && r.head-r.tail <= r.len/10 {
		r.resize(r.len / 5)
	}
	return v
}

func (r *ring) resize(size int) {
	newbuf := make([]Buffer, size)
	t := r.tail % r.len
	h := r.head % r.len
	// note: extend is normally called before t == h
	if t >= h {
		copy(newbuf, r.buff[t:])
		copy(newbuf[r.len-t:], r.buff[:h+1])
		r.head = r.len - t + h
	} else {
		copy(newbuf, r.buff[t:h+1])
		r.head = h - t
	}
	r.buff = newbuf
	r.len = size
	r.tail = 0
}

func newRingBuf(capacity int) *ring {
	r := ring{
		buff:     make([]Buffer, capacity),
		capacity: capacity,
		len:      capacity,
		head:     -1,
		tail:     0,
		In:       make(chan Buffer, capacity/10),
		Out:      make(chan Buffer, capacity/10),
	}
	go r.run()
	return &r
}

func (rx *Rx) consumeChunks(ring *ring) {
	defer close(rx.Packets)
	defer rx.state.SetDone()
	for buf := range ring.Out {
		for i := 0; i < len(buf); i++ {
			rx.Packets <- buf[i]
		}
	}
}

func (rx *Rx) captureChunks(ring *ring, handle *pcap.Handle) {
	defer close(ring.In)
	defer handle.Close()

	count := uint32(0)

main:
	for rx.pktCount == 0 || count < rx.pktCount {
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
					if rx.state.Stopping() {
						break main
					}
					if last >= 0 {
						// this should not block too long, since the ring
						// buffer prioritizes producer over consumer
						ring.In <- buf[:last+1]
						count += uint32(last + 1)
						continue main
					}
				}
			}
		}
		if rx.state.Stopping() {
			break main
		}
		// this should not block too long, since the ring buffer prioritizes
		// producer over consumer
		ring.In <- buf[:last+1]
		count += uint32(last + 1)
	}
	rx.setStats(handle)
}
