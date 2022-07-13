package main

import (
	"flag"
	"fmt"
	"io"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	statistics map[string]int = map[string]int{"TCP": 0, "UDP": 0, "IPv4": 0, "IPv6": 0}

	eth layers.Ethernet
	ip4 layers.IPv4
	ip6 layers.IPv6
	tcp layers.TCP
	udp layers.UDP
	dns layers.DNS
)

func main() {
	parser := gopacket.NewDecodingLayerParser(
		layers.LayerTypeEthernet,
		&eth,
		&ip4,
		&ip6,
		&tcp,
		&udp,
		&dns,
	)

	fileName := flag.String("file", "lo.pcap", "The path to the input file")
	flag.Parse()
	handle, err := pcap.OpenOffline(*fileName)
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	decoded := make([]gopacket.LayerType, 0, 10)
	for {
		data, _, err := handle.ZeroCopyReadPacketData()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		parser.DecodeLayers(data, &decoded)

		for _, layer := range decoded {
			switch layer {
			case layers.LayerTypeIPv4:
				statistics["IPv4"]++
			case layers.LayerTypeIPv6:
				statistics["IPv6"]++
			case layers.LayerTypeTCP:
				statistics["TCP"]++
			case layers.LayerTypeUDP:
				statistics["UDP"]++
			}
		}
	}

	print(statistics)
}

func print(arg map[string]int) {
	fmt.Println("Amounts of TCP, UDP, IPv4 and IPv6 packets:")
	for layer, amount := range arg {
		fmt.Printf("%v: %v\n", layer, amount)
	}

}
