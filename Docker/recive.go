package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type rawpacket struct {
	Type string `json:"type"`
}

type PacketTCP struct {
	Type    string `json:"Type"`
	SrcMac  string `json:"SrcMac"`
	DstMac  string `json:"DstMac"`
	SrcIP   string `json:"SrcIP"`
	DstIP   string `json:"DstIP"`
	SrcPort int    `json:"SrcPort"`
	DstPort int    `json:"DstPort"`
	TTL     int    `json:"TTL"`
	SeqNum  int    `json:"SeqNum"`
	AckNum  int    `json:"AckNum"`
	Flags   string `json:"Flags"`
	WinSize int    `json:"WinSize"`
	Payload string `json:"Payload"`
}

type PacketUdp struct {
	SrcIP   string `json:"srcIP"`
	DstIP   string `json:"dstIP"`
	SrcPort int    `json:"srcPort"`
	DstPort int    `json:"dstPort"`
	SrcMAC  string `json:"srcMAC"`
	DstMAC  string `json:"dstMAC"`
	Payload string `json:"payload"`
}

type PacketICMPv4 struct {
	ICMPType int    `json:"icmpType"`
	ICMPCode int    `json:"icmpCode"`
	SrcIP    string `json:"srcIP"`
	DstIP    string `json:"dstIP"`
	SrcMAC   string `json:"srcMAC"`
	DstMAC   string `json:"dstMAC"`
}

type PacketIP struct {
	SrcMac  string `json:"SrcMac"`
	DstMac  string `json:"DstMac"`
	SrcIp   string `json:"SrcIp"`
	DstIp   string `json:"DstIp"`
	SrcPort int    `json:"SrcPort"`
	DstPort int    `json:"DstPort"`
	Payload string `json:"Payload"`
	TTL     int    `json:"TTL"`
}

func main() {
	// Listen on port 8484 for incoming connections
	listener, err := net.Listen("tcp", ":8484")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// Wait for a client to connect
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	// Parse the JSON data sent by the client
	var rawpacket PacketTCP
	err = json.NewDecoder(conn).Decode(&rawpacket)
	if err != nil {
		panic(err)
	}

	fmt.Println(rawpacket.Type)

	// // Call the appropriate handler function based on the packet type
	// switch packetType {
	// case "tcp":
	// 	handle_tcp(packet)
	// case "udp":
	// 	handle_udp(packet)
	// case "icmp":
	// 	handle_icmp(packet)
	// case "ip":
	// 	handle_ip(packet)
	// default:
	// 	log.Fatalf("Error: unknown packet type '%s'", packetType)
	// }
}
