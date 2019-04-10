package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	addr := net.ParseIP("192.0.2.1")
	if addr == nil {
		fmt.Println("Invalid	address")
		os.Exit(1)
	}
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println("Address	is	", addr.String(),
		"\nDefault	mask	length	is	", bits, "\nLeading	ones	count	is	", ones,
		"\nMask	is	(hex)	", mask.String(), "\nNetwork	is	", network.String())
	os.Exit(0)
}
