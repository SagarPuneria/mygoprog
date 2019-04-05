/* ASN.1 DaytimeClient
 */
package main

import (
	"encoding/asn1"
	"fmt"
	"io/ioutil"
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

	conn, err := net.Dial("tcp", service)
	checkError(err)

	defer conn.Close()

	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println("result:", result, string(result))

	var newtime time.Time
	fmt.Println("Before unmarshall newtime:", newtime)
	_, err1 := asn1.Unmarshal(result, &newtime)
	checkError(err1)
	fmt.Println("After unmarshall newtime:", newtime)

	fmt.Println("After unmarshal: ", newtime.String())

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
