package main

import (
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
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

	count := 0
	for range packets {
		count++
	}

	fmt.Println(count)
}
