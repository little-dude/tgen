package server

// THA = Target Hardware Address
// TIA = Target Internet Address
// SHA = Sender Hardware Address
// SIA = Sender Internet Address

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"net"
	"sync"
)

type ARPInstance struct {
	lock       sync.RWMutex
	stop       chan bool
	Cache      *ARPMap
	Interfaces *ARPMap
	enabled    bool
	Packets    chan *layers.ARP
}

func NewARPInstance(out chan []byte) *ARPInstance {
	return &ARPInstance{
		stop:       make(chan bool),
		Cache:      NewARPMap(),
		Interfaces: NewARPMap(),
		enabled:    false,
	}
}

func (proto *ARPInstance) Enable() {
	proto.lock.Lock()
	proto.enabled = true
	proto.lock.Unlock()
}

func (proto *ARPInstance) Disable() {
	proto.lock.Lock()
	proto.enabled = false
	proto.lock.Unlock()
}

func (proto *ARPInstance) IsEnabled() bool {
	proto.lock.Lock()
	defer proto.lock.Unlock()
	return proto.enabled
}

func (proto *ARPInstance) Stop() {
	close(proto.Packets)
}

func (proto *ARPInstance) Start(out chan<- []byte) {
	proto.Packets = make(chan *layers.ARP, 1000)
	Info.Println("starting ARP")
	go func() {
		defer func() {
			Info.Println("stopping ARP")
		}()
		for arp := range proto.Packets {
			// FIXME? we assume this is a valid ARP packet for ipv4 over ethernet but
			// we should check the Hardware Type and Protocol Type to make sure of it
			// also, we should check the hardware address length and protocol address
			// length
			switch arp.Operation {
			case layers.ARPRequest:
				proto.Cache.Set(net.IP(arp.SourceProtAddress), net.HardwareAddr(arp.SourceHwAddress))
				reply, e := proto.Reply(arp)
				if e != nil {
					Error.Println(e.Error())
				} else {
					out <- reply
				}
			case layers.ARPReply:
				proto.Cache.Set(net.IP(arp.SourceProtAddress), net.HardwareAddr(arp.SourceHwAddress))
			default:
				Error.Println("got an ARP packet that is not a request nor a reply")
			}
		}
	}()
}

func (proto *ARPInstance) Reply(request *layers.ARP) ([]byte, error) {
	tha, ok := proto.Interfaces.Get(net.IP(request.DstProtAddress))
	if !ok {
		return []byte{}, NewError(net.IP(request.DstProtAddress).String(), "not found in local ARP table")
	}
	eth := layers.Ethernet{
		SrcMAC:       tha,
		DstMAC:       net.HardwareAddr(request.SourceHwAddress),
		EthernetType: layers.EthernetTypeARP,
	}
	arp := layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     6,
		ProtAddressSize:   4,
		Operation:         layers.ARPReply,
		SourceHwAddress:   []byte(tha),
		SourceProtAddress: request.DstProtAddress,
		DstHwAddress:      request.SourceHwAddress,
		DstProtAddress:    request.SourceProtAddress,
	}
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{FixLengths: true, ComputeChecksums: true}
	e := gopacket.SerializeLayers(buf, opts, &eth, &arp)
	if e != nil {
		return []byte{}, NewError("Could not create ARP reply data:", e.Error())
	}
	return buf.Bytes(), nil
}

type ARPMap struct {
	lock  sync.RWMutex
	table map[uint32]net.HardwareAddr
}

func (m *ARPMap) Get(ip net.IP) (target net.HardwareAddr, ok bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	target, ok = m.table[IPv4ToInt(ip)]
	return
}

func (m *ARPMap) Set(ip net.IP, mac net.HardwareAddr) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.table[IPv4ToInt(ip)] = mac
}

func (m *ARPMap) Del(ip net.IP) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.table, IPv4ToInt(ip))
}

func (m *ARPMap) Has(ip net.IP) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	_, ok := m.table[IPv4ToInt(ip)]
	return ok
}

func (m *ARPMap) Len() int32 {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return int32(len(m.table))
}

func NewARPMap() *ARPMap {
	return &ARPMap{
		table: make(map[uint32]net.HardwareAddr),
	}
}
