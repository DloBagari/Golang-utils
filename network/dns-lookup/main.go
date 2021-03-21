package main

import (
	"fmt"
	"log"
	"net"
)

//DNS: Domain Name System. translate ip address to a name
func lookIp(address string) ([]string, error) {
	hosts, err := net.LookupAddr(address)
	if err != nil {
		return nil, err
	}
	return hosts, err
}

func lookHostName(hostName string) ([]string, error) {
	ips, err := net.LookupHost(hostName)
	if err != nil {
		return nil, err
	}
	return ips, err
}

func main() {
	// if the ip address is local, it will read the host name from /etc/hosts
	ips, err := lookIp("127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ips)
	hosts, err := lookHostName("www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hosts)
}
