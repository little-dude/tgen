package rxtx

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/little-dude/tgen/server/errors"
	"github.com/little-dude/tgen/server/log"
	"time"
)

type PcapStats struct {
	Received uint32
	KDropped uint32
	IDropped uint32
}

type Rx struct {
	State    *RxTxState
	stats    PcapStats
	port     string
	Packets  chan *RawPacket
	pktCount uint32
}

type RawPacket struct {
	Data []byte
	Ci   gopacket.CaptureInfo
}

func (rx *Rx) Stats() (PcapStats, error) {
	if rx.State.Done() {
		// no need to lock to access the stats here since when capture is done,
		// no goroutine should access the stats anymore
		return rx.stats, nil
	}
	if rx.State.Inactive() {
		return PcapStats{}, errors.New("capture did not start yet")
	}
	if rx.State.Active() {
		return PcapStats{}, errors.New("still capturing")
	}
	panic("Unknown state")
}

func NewRx(port string) *Rx {
	rx := Rx{port: port, State: NewRxTxState()}
	return &rx
}

func (rx *Rx) getPcapHandle(direction pcap.Direction, bpf string) (handle *pcap.Handle, e error) {
	handle, e = pcap.OpenLive(
		rx.port,
		65635, // capture max packet size by default
		true,  // set promiscuous mode
		time.Millisecond*10)
	if e != nil {
		return
	}

	e = handle.SetDirection(direction)
	if e != nil {
		handle.Close()
		return
	}

	if bpf != "" {
		e = handle.SetBPFFilter(bpf)
		if e != nil {
			handle.Close()
			return
		}
	}
	return
}

func (rx *Rx) Capture(packets chan<- *RawPacket, pktCount uint32, direction pcap.Direction, bpf string) error {
	handle, e := rx.getPcapHandle(direction, bpf)
	if e != nil {
		return e
	}

	rx.pktCount = pktCount
	rx.State.SetRun()
	go rx.capture(handle, packets)
	return nil
}

func (rx *Rx) capture(handle *pcap.Handle, packets chan<- *RawPacket) {
	defer func(handle *pcap.Handle, packets chan<- *RawPacket) {
		log.Info.Println("done capturing")
		close(packets)
		handle.Close()
		rx.State.SetDone()
	}(handle, packets)

	count := uint32(0)
	for {
		data, ci, e := handle.ReadPacketData()
		if e == nil {
			count++
			// WARNING: this can block forever is nothing is consuming the incoming packets
			packets <- &RawPacket{Data: data, Ci: ci}
		}
		if !(count < rx.pktCount || rx.pktCount == 0) || rx.State.Stopping() {
			break
		}
	}
	rx.setStats(handle)
}

func (rx *Rx) CaptureChunks(chunks chan<- []*RawPacket, pktCount uint32, direction pcap.Direction, bpf string) error {
	handle, e := rx.getPcapHandle(direction, bpf)
	if e != nil {
		return e
	}
	rx.pktCount = pktCount
	rx.State.SetRun()
	go rx.captureChunks(handle, chunks)
	return nil
}

func (rx *Rx) captureChunks(handle *pcap.Handle, chunks chan<- []*RawPacket) {
	defer rx.State.SetDone()
	defer close(chunks)
	defer handle.Close()

	count := uint32(0)

main:
	for rx.pktCount == 0 || count < rx.pktCount {
		buf := make([]*RawPacket, 1000)
		last := -1
		for i := 0; i < 1000; i++ {
			for {
				data, ci, e := handle.ReadPacketData()
				if e == nil {
					buf[i] = &RawPacket{Data: data, Ci: ci}
					last = i
					break
				} else if e == pcap.NextErrorTimeoutExpired {
					if rx.State.Stopping() {
						break main
					}
					if last >= 0 {
						chunks <- buf[:last+1]
						count += uint32(last + 1)
						continue main
					}
				}
			}
		}
		if rx.State.Stopping() {
			break main
		}
		chunks <- buf[:last+1]
		count += uint32(last + 1)
	}
	rx.setStats(handle)
}

func (rx *Rx) setStats(handle *pcap.Handle) error {
	stats, e := handle.Stats()
	if e != nil {
		return errors.New("Failed to get capture stats:", e.Error())
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

func (rx *Rx) Stop() {
	rx.State.SetStop()
	rx.State.WaitDone(0)
}
