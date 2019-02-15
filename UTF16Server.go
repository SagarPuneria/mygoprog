/* UTF16 Server
 */
package main

import (
	"fmt"
	"net"
	"os"
	"unicode/utf16"
)

const BOM = '\ufffe'

func main() {

	fmt.Println("BOM:", BOM)
	service := "127.0.0.1:1210"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		str := "golang百度一下，你就知道" //"j'ai arrêté"
		fmt.Println("str:", str)
		shorts := utf16.Encode([]rune(str))
		fmt.Println("[]rune(str):", []rune(str))
		fmt.Println("shorts:", shorts)
		writeShorts(conn, shorts)

		conn.Close() // we're finished
	}
}

func writeShorts(conn net.Conn, shorts []uint16) {
	var bytes [2]byte

	// send the BOM as first two bytes in big  endian format
	bytes[0] = BOM >> 0x08 // BOM >> 8
	bytes[1] = BOM & 0xff  // BOM & 255
	fmt.Println("> bytes[0:]:", bytes[0:], ",string(bytes[0:]):", string(bytes[0:]))
	_, err := conn.Write(bytes[0:])
	if err != nil {
		return
	}

	for _, v := range shorts {
		bytes[0] = byte(v >> 8)
		bytes[1] = byte(v & 255)
		fmt.Println("bytes[0:]:", bytes[0:], ",string(bytes[0:]):", string(bytes[0:]))

		_, err = conn.Write(bytes[0:])
		if err != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
