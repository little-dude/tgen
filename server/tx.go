package server

import "github.com/google/gopacket/pcap"

type Tx struct {
	state *RxTxState
}

func (tx *Tx) Start(handle *pcap.Handle, streams []*Stream) {
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
}

func NewTx() *Tx {
	tx := Tx{
		state: NewRxTxState(),
	}
	return &tx
}
