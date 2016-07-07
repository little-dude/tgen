package server

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/little-dude/tgen/schemas"
	"net"
	//	"strconv"
	//	"strings"
)

type LAN struct {
	ARP        *ARPInstance
	Interfaces *ARPMap
	Network    *net.IPNet
	Rx         *Rx
	Tx         *Tx
	VLANs      []uint32
	port       string
	packets    chan *RawPacket
}

func NewLAN(port string, cidr string, vlans []uint32) (*LAN, error) {
	_, network, e := net.ParseCIDR(cidr)
	if e != nil {
		return &LAN{}, NewError("Invalid network address:", e.Error())
	}
	l := LAN{
		Rx:         NewRx(port),
		Tx:         NewTx(port),
		Network:    network,
		VLANs:      vlans,
		port:       port,
		Interfaces: NewARPMap(),
	}
	l.ARP = NewARPInstance(l.Tx.Out)
	l.ARP.Interfaces = l.Interfaces
	return &l, nil
}

func (l *LAN) HandlePackets() {
	go func() {
		for rawPkt := range l.packets {
			pkt := gopacket.NewPacket(rawPkt.data, layers.LayerTypeEthernet, gopacket.Lazy)
			switch pkt.Layer(layers.LayerTypeEthernet).(*layers.Ethernet).NextLayerType() {
			case layers.LayerTypeARP:
				// ARP
				if l.ARP.IsEnabled() {
					l.ARP.Packets <- pkt.Layer(layers.LayerTypeARP).(*layers.ARP)
				}
			case layers.LayerTypeIPv4:
				// TODO
			default:
				// ignore the packet if not ARP or IPv4
			}
		}
	}()
}

func (l *LAN) BuildBPF() string {
	// build a filter that looks like this:
	//
	// vlan X && vlan Y && vlan Z && (
	//     arp ||
	//     (ip  && (
	//         dst net IP mask MASK || dst net 255.255.255.255))
	// )
	//
	//

	// "vlan X && vlan Y && vlan Z &&"
	//	filter := make([]string, 0)
	//
	//	for vlan := range l.VLANs {
	//		filter = append(filter, "vlan", strconv.Itoa(vlan), "&&")
	//	}
	//
	//	// open a parenthesis
	//	filter = append(filter, "(")
	//
	//	// "arp ||"
	//	filter = append(filter, "arp", "||")
	//
	//	// (ip && (dst net IP mask MASK || dst net 255.255.255.255))"
	//	filter = append(filter, "(", "ip", "&&", "(", "dst", "net", l.Network.IP.String(), "mask", l.Network.Mask.String(), "||", "dst", "net", "255.255.255.255", ")", ")")
	//
	//	// close the parenthesis
	//	filter = append(filter, ")")
	//
	//	f := strings.Join(filter, " ")
	//	Trace.Println("BPF filter:", f)
	//  return f
	return ""
}

func (l *LAN) GetConfig(call schemas.Lan_getConfig) error {
	config, e := call.Results.NewConfig()
	if e != nil {
		return e
	}
	config.SetCidr(l.Network.String())
	devices, e := config.NewDevices(l.Interfaces.Len())
	if e != nil {
		return e
	}
	i := 0

	// FIXME: implement a proper way to iterate over the interfaces, this is
	// not even thread safe afaik. see for example:
	// https://stackoverflow.com/questions/35810674/in-go-golang-is-it-possible-to-iterate-over-a-custom-type/35810778#35810778
	for ip, mac := range l.Interfaces.table {
		dev := devices.At(i)
		dev.SetIp([]byte(IPv4FromInt(ip)))
		dev.SetMac([]byte(mac))
		// devices.Set(i, dev) ?
		i++
	}
	return nil
}

func (l *LAN) SetConfig(call schemas.Lan_setConfig) error {
	config, e := call.Params.Config()
	if e != nil {
		return e
	}

	cidr, e := config.Cidr()
	if e != nil {
		return e
	}
	if cidr != "" && cidr != l.Network.String() {
		return NewError("The LAN CIDR cannot be changed")
	}

	devices, e := config.Devices()
	if e != nil {
		return e
	}
	for i := 0; i < devices.Len(); i++ {
		device := devices.At(i)
		ip, e := device.Ip()
		if e != nil {
			return e
		}
		mac, e := device.Mac()
		if e != nil {
			return e
		}
		l.Interfaces.Set(net.IP(ip), net.HardwareAddr(mac))
	}
	return nil
}

func (l *LAN) Start(call schemas.Lan_start) error {
	if l.Rx.state.Active() {
		return NewError("LAN is already active")
	}
	Info.Println("starting lan", l.Network)
	l.Tx.Start()

	// start the goroutine that handles ARP packets
	l.ARP.Start(l.Tx.Out)
	l.ARP.Enable()
	l.packets = make(chan *RawPacket, 1000)

	// start the goroutine that handle captured packets
	l.HandlePackets()

	// start the capture
	e := l.Rx.Capture(l.packets, 0, pcap.DirectionIn, l.BuildBPF())
	if e != nil {
		Error.Println(e.Error())
		return e
	}
	Info.Println("LAN started")
	return nil
}

func (l *LAN) Stop(call schemas.Lan_stop) error {
	l.Rx.Stop()
	l.ARP.Disable()
	l.ARP.Stop()
	return nil
}
