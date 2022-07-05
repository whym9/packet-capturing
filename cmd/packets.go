package main

import (
	"flag"
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func main() {
	fileName := flag.String("input", "lo.cap", "input file")
	flag.Parse()
	handle, err := pcap.OpenOffline(*fileName)
	if err != nil {
		panic(err)
	}
	defer handle.Close()

	packets := gopacket.NewPacketSource(
		handle, handle.LinkType()).Packets()

	count := 0
	for range packets {
		count++
	}

	fmt.Println(count)
}
