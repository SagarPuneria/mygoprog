/* UDPDaytimeServer
 */
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	service := "127.0.0.1:1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError("ResolveUDPAddr", err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError("ListenUDP", err)
	defer conn.Close()
	fmt.Println("ListenUDP")

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {

	var buf [512]byte

	n, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		fmt.Println("ReadFromUDP, error:", err)
		os.Exit(1)
	}
	fmt.Println("ReadFromUDP is buf:", string(buf[:n]))

	daytime := time.Now().String()

	_, err = conn.WriteToUDP([]byte(daytime), addr)
	if err != nil {
		fmt.Println("WriteToUDP, error:", err)
		os.Exit(1)
	}
	fmt.Println("WriteToUDP")
}

func checkError(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error at "+msg+": %s", err.Error())
		os.Exit(1)
	}
}
