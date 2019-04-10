/* ThreadedEchoServer
 */
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	service := "127.0.0.1:1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError("ResolveTCPAddr", err)

	//listener,	err	:=	net.Listen("tcp", service) //OR
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError("ListenTCP", err)
	defer listener.Close()

	for {
		//conn, err := listener.Accept() //OR
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("connection accept failed")
			continue
		}
		fmt.Println("connection accepted success")
		// run as a goroutine
		go handleClient(conn)
	}
}

//func handleClient(conn net.Conn) { //OR
func handleClient(conn *net.TCPConn) {
	// close connection on exit
	defer conn.Close()

	err := conn.SetKeepAlive(true)
	if err != nil {
		fmt.Println("SetKeepAlive error:", err)
		return
	}
	timeout := 5 * time.Second
	// SetDeadline sets the read and write deadlines associated
	// with the connection. It is equivalent to calling both
	// SetReadDeadline and SetWriteDeadline.
	// A deadline is an absolute time after which I/O operations fail with a timeout
	// (see type Error) irrespective of blocking/waiting at Read or Write operation.
	err = conn.SetDeadline(time.Now().Add(timeout))
	if err != nil {
		fmt.Println("SetDeadline error:", err)
		return
	}
	var buf [512]byte
	for {
		// read upto 512 bytes
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		fmt.Println("server read:", string(buf[0:]))

		// write the n bytes read
		_, err = conn.Write(buf[0:n])
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}
	}
}

func checkError(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error at "+msg+": %s", err.Error())
		os.Exit(1)
	}
}
