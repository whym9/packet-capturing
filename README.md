# packet-capturing 
This repository is made by me to learn and practice packet capturing with go.
## packets.go
This program has a command line argument '-input' that takes the directory to the pcap file as an argument or by default "lo.pcap".
Then it opens the file and prints out the amount of packets inside it.
It panics if the errors occured while opening the file.
It uses **flag** library to set command line arguments and default argument. And it uses **goPacket** library to open pcap files and work with packages.

## packet_statistics.go
This program has a command line argument '-file'. It takes the directory to the pcap file or by default "lo.pcap" as an argument.
It then opens the file and prints out the amount of packets with TCP, UDP, IPv4 and IPv6 protocols. 
It panics if the errors occured while opening the file.
It uses **flag** library to set command line arguments and default argument. And it uses **goPacket** library to open pcap files and work with packages.
