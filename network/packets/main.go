package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"log"
)

//return list of network devices
func getAllDevices() ([]pcap.Interface, error) {
	//find all network devices
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func describeNetworkInterface(itf pcap.Interface) {
	fmt.Println("---------------------")
	fmt.Println("Name: ", itf.Name)
	fmt.Println("Description: ", itf.Description)
	for _, address := range itf.Addresses {
		fmt.Println("-- Ip Address: ", address.IP)
		fmt.Println("-- Subnet mask: ", address.Netmask)
	}
}

func main() {
	networkDevices, err := getAllDevices()
	if err != nil {
		log.Fatal(err)
	}
	for _, device := range networkDevices {
		describeNetworkInterface(device)
	}
}
