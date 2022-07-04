package main

import (
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	statistics map[string]int = map[string]int{"TCP": 0, "UDP": 0, "IPv4": 0, "IPv6": 0}
)

func main() {

	fileName := os.Args[1]
	handle, err := pcap.OpenOffline(fileName)
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	packets := gopacket.NewPacketSource(
		handle, handle.LinkType()).Packets()

	for pkt := range packets {
		packetStatistics(pkt)
	}
	print(statistics)
}
func packetStatistics(packet gopacket.Packet) {

	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		statistics["TCP"]++
	}

	if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
		statistics["UDP"]++
	}

	if ipv4Layer := packet.Layer(layers.LayerTypeIPv4); ipv4Layer != nil {
		statistics["IPv4"]++
	}

	if ipv6Layer := packet.Layer(layers.LayerTypeIPv6); ipv6Layer != nil {
		statistics["IPv6"]++
	}
}

func print(arg map[string]int) {
	fmt.Println("Amounts of TCP, UDP, IPv4 and IPv6 packets:")
	for layer, amount := range arg {
		fmt.Printf("%v: %v\n", layer, amount)
	}

}
