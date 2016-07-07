package server

import (
	"github.com/google/gopacket/pcap"
	"time"
)

type Tx struct {
	state *RxTxState
	port  string
	Out   chan []byte
}

func (tx *Tx) SendStreams(streams []*Stream) {
	handle, e := tx.openHandle()
	if e != nil {
		Error.Println(e.Error())
		return
	}

	go func(handle *pcap.Handle, streams []*Stream) {
		tx.state.SetRun()
		defer func() {
			tx.state.SetDone()
			Info.Println("Done sending")
		}()
	outer:
		for _, stream := range streams {
			Info.Println("Starting to send stream", stream.ID)
			for i := 0; i < len(stream.Packets); i++ {
				e := handle.WritePacketData(stream.Packets[i])
				if e != nil {
					Error.Println(e.Error())
				}
				if tx.state.Stopping() {
					break outer
				}
			}
		}
	}(handle, streams)
}

func (tx *Tx) openHandle() (*pcap.Handle, error) {
	handle, e := pcap.OpenLive(tx.port, 65635, true, time.Millisecond*10)
	if e != nil {
		return handle, NewError("Could not create pcap handle:", e.Error())
	}

	e = handle.SetDirection(pcap.DirectionOut)
	if e != nil {
		return handle, NewError("Could not set pcap handle direction:", e.Error())
	}
	return handle, nil
}

func (tx *Tx) Start() error {
	handle, e := tx.openHandle()
	if e != nil {
		return e
	}
	tx.Out = make(chan []byte, 1000)
	go tx.send(handle)
	return nil
}

func (tx *Tx) send(handle *pcap.Handle) {
	defer handle.Close()
	for data := range tx.Out {
		e := handle.WritePacketData(data)
		if e != nil {
			Error.Println("Failed to send packet:", e.Error())
		}
	}
}

func NewTx(port string) *Tx {
	tx := Tx{
		state: NewRxTxState(),
		port:  port,
	}
	return &tx
}
