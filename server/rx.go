package server

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"os"
	"time"
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
		buf := newRingBug(10000)
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
			select {
			case rx.Packets <- &RawPacket{data: data, ci: ci}:
			default:
			}
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
			if rx.state.Stopping() {
				break main
			}
		default:
			Error.Println("could not push buffer into ring... lost a buffer")
		}
		count += uint32(last + 1)
	}
	rx.setStats(handle)
}
