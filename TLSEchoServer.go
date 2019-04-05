/* TLSEchoServer
 */
package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	cert, err := tls.LoadX509KeyPair("jan.newmarch.name.pem", "private.pem")
	checkError("LoadX509KeyPair", err)
	config := tls.Config{Certificates: []tls.Certificate{cert}}

	now := time.Now()
	config.Time = func() time.Time { return now }
	config.Rand = rand.Reader

	listener, err := tls.Listen("tcp", service, &config)
	checkError("Listen", err)
	fmt.Println("Listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("Accepted")
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	for {
		fmt.Println("Trying to read")
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		_, err = conn.Write(buf[0:n])
		if err != nil {
			return
		}
	}
}

func checkError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
