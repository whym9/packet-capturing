package main

import (
	"flag"
	"fmt"
	"io"

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

	count := 0
	for {
		_, _, err := handle.ZeroCopyReadPacketData()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		count++
	}
	fmt.Println(count)
}
