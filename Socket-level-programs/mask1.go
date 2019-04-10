package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("192.0.2.1")
	mask := ip.DefaultMask()
	fmt.Println("DefaultMask:", mask)
	fmt.Println("Network starting address:", ip.Mask(mask))

	ipv4Mask := net.CIDRMask(24, 32)
	fmt.Println("CIDR ipv4Mask:", ipv4Mask)
	fmt.Println("CIDR Network starting address:", ip.Mask(ipv4Mask))
}
