package main

import (
	"fmt"
	"log"
	"net"
)

// find the name servers of a domain, these ns are stored in the NS records of the domain
// also we can lookup for mail servers of the domain, the MX records, by using net.lookupMX

func main() {
	ns, err := net.LookupNS("google.com")
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range ns {
		fmt.Println(i.Host)
	}
}
