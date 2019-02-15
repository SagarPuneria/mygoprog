/* UDPDaytimeClient
 */
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError("ResolveUDPAddr", err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError("DialUDP", err)
	defer conn.Close()

	_, err = conn.Write([]byte("anything"))
	checkError("Write", err)

	time.Sleep(10 * time.Second)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError("Read", err)

	fmt.Println(string(buf[0:n]))

	os.Exit(0)
}

func checkError(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error at "+msg+": %s", err.Error())
		os.Exit(1)
	}
}
