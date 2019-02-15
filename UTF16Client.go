/* UTF16 Client
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
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError(err)

	shorts := readShorts(conn)
	fmt.Println("shorts:", shorts)
	runes := utf16.Decode(shorts)
	str := string(runes)

	fmt.Println(str)

	os.Exit(0)
}

func readShorts(conn net.Conn) []uint16 {
	var buf [512]byte

	// read everything into the buffer
	n, err := conn.Read(buf[0:2])
	fmt.Println("buf[0:2]:", buf[0:2], ",string(buf[0:2]):", string(buf[0:2]))
	for true {
		m, err := conn.Read(buf[n:])
		if m == 0 || err != nil {
			break
		}
		n += m
		fmt.Println("buf[:n]:", buf[:n], ",string(buf[:n]):", string(buf[:n]))
	}

	checkError(err)
	var shorts []uint16
	shorts = make([]uint16, n/2)

	if buf[0] == 0xff && buf[1] == 0xfe {
		// big endian
		for i := 2; i < n; i += 2 {
			shorts[i/2-1] = uint16(buf[i])<<8 + uint16(buf[i+1])
		}
		fmt.Println("big endian shorts[:]:", shorts[:])
	} else if buf[1] == 0xfe && buf[0] == 0xff {
		// little endian
		for i := 2; i < n; i += 2 {
			shorts[i/2-1] = uint16(buf[i+1])<<8 + uint16(buf[i])
		}
		fmt.Println("little endian shorts[:]:", shorts[:])
	} else {
		// unknown byte order
		fmt.Println("Unknown order")
	}
	return shorts

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
