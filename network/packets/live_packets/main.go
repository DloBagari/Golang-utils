package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

// capture packets in live mode
// this program needs to run as root in order to allow it to do syscall

var (
	device            = "ens33"
	snapshotLen int32 = 1024
	promiscuous       = true // for wireless make this true
	err         error
	timeout     = 30 * time.Second
	handle      *pcap.Handle
)

func main() {
	// open device
	handle, err = pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	// set filter on handler if you nee that
	filter := "tcp and port 443"
	if err := handle.SetBPFFilter(filter); err != nil {
		log.Fatal(err)
	}
	// use handle as packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		//process the packet
		fmt.Println("\n------------------------")
		fmt.Println(packet)
	}
}
