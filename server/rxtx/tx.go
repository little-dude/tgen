package rxtx

import (
	"github.com/google/gopacket/pcap"
	"github.com/little-dude/tgen/server/errors"
	"github.com/little-dude/tgen/server/log"
	"time"
)

type Tx struct {
	State *RxTxState
	port  string
	Out   chan []byte
}

func (tx *Tx) openHandle() (*pcap.Handle, error) {
	handle, e := pcap.OpenLive(tx.port, 65635, true, time.Millisecond*10)
	if e != nil {
		return handle, errors.New("Could not create pcap handle:", e.Error())
	}

	e = handle.SetDirection(pcap.DirectionOut)
	if e != nil {
		return handle, errors.New("Could not set pcap handle direction:", e.Error())
	}
	return handle, nil
}

func (tx *Tx) Start() error {
	handle, e := tx.openHandle()
	if e != nil {
		return e
	}
	tx.Out = make(chan []byte, 1000)
	tx.State.SetRun()
	go tx.send(handle)
	return nil
}

func (tx *Tx) send(handle *pcap.Handle) {
	defer handle.Close()
	defer tx.State.SetDone()
	for data := range tx.Out {
		e := handle.WritePacketData(data)
		if e != nil {
			log.Error.Println("Failed to send packet:", e.Error())
		}
	}
}

func NewTx(port string) *Tx {
	tx := Tx{
		State: NewRxTxState(),
		port:  port,
	}
	return &tx
}
