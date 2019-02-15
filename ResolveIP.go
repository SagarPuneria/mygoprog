package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:	%s	hostname\n", os.Args[0])
		fmt.Println("Usage:	", os.Args[0], "hostname")
		os.Exit(1)
	}
	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", name) // name = "LMIPL-234" or "www.google.com"
	if err != nil {
		fmt.Println("Resolution	error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved	address	is	", addr.String())  //172.217.163.36 > if name = "www.google.com"
	fmt.Println("Resolved	network	is	", addr.Network()) // ip > if name = "www.google.com"
	fmt.Println("Address IP is	", addr.IP)              //172.217.163.36 > if name = "www.google.com"
	fmt.Println("Address Zone is	", addr.Zone)          // No Output
	os.Exit(0)
}
