package main

import (
	"fmt"
	"log"
	"net"
)

// if we need to find out the default gateway of a machine, use this command: netstat -nr

func main() {
	//read all exist network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range interfaces {
		fmt.Println(i.Name)
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			log.Fatal(err)
		}
		//interface capabilities
		fmt.Printf("interface Name: %v\n", i.Name)
		fmt.Printf("interface Flags: %v\n", i.MTU)
		fmt.Printf("interface hardware address: %v\n", i.HardwareAddr)
		fmt.Println("##addresses:")
		ddresses, err := byName.Addrs()
		for k, v := range ddresses {
			fmt.Printf("interface address #%v: %v\n", k, v.String())
		}
		fmt.Println("-------------")
	}
}
