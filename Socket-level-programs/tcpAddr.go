package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage:	%s	hostname\n", os.Args[0])
		fmt.Println("Usage:	", os.Args[0], "hostname")
		os.Exit(1)
	}
	networkType := os.Args[1]
	addr := os.Args[2]
	//addr = [2001:4860:4860::8888]:80 or www.google.com:80 or 8.8.8.8:80
	tcpAddr, err := net.ResolveTCPAddr(networkType, addr)
	if err != nil {
		fmt.Println("Resolution	error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved	tcp address	is	", tcpAddr.String()) //172.217.163.36 > if name = "www.google.com"
	fmt.Println("Resolved	network	is	", tcpAddr.Network())    // ip > if name = "www.google.com"
	fmt.Println("Address IP is	", tcpAddr.IP)                 //172.217.163.36 > if name = "www.google.com"
	fmt.Println("Address Zone is	", tcpAddr.Zone)             // No Output
	fmt.Println("Port is	", tcpAddr.Port)                     // 80
	os.Exit(0)
}
