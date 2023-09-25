package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("%s: usage: <host>", os.Args[0])
	}
	host := os.Args[1]
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatalf("lookup ip: %s: %v", host, err)
	}
	if len(ips) == 0 {
		log.Fatalf("no IPs found for %s", host)
	}

	var ipv4Found bool
	var ipv6Found bool

	for _, ip := range ips {
		if ip.To4() != nil {
			fmt.Println(ip)
			ipv4Found = true
		} else {
			fmt.Println(ip)
			ipv6Found = true
		}
	}

	if !ipv4Found {
		fmt.Println("No IPv4 addresses found.")
	}

	if !ipv6Found {
		fmt.Println("No IPv6 addresses found.")
	}
}
