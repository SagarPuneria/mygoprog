/* JSON EchoServer
 */
package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

func main() {

	service := "127.0.0.1:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError("ResolveTCPAddr", err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError("ListenTCP", err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		encoder := json.NewEncoder(conn)
		decoder := json.NewDecoder(conn)

		for n := 0; n < 10; n++ {
			var person Person
			err = decoder.Decode(&person)
			checkError("Decode:", err)
			fmt.Println(person.String())
			err = encoder.Encode(person)
			checkError("Encode:", err)
		}
		conn.Close() // we're finished
	}
}

func checkError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
