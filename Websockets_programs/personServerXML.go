/* PersonServerXML
 */
package main

import (
	"fmt"
	"net/http"
	"os"

	"XmlCodeC" // Custom package PATH can be imported as case-insensitive

	"golang.org/x/net/websocket"
)

type Person struct {
	Name   string
	Emails []string
}

func ReceivePerson(ws *websocket.Conn) {
	var person Person
	err := xmlCodec.XMLCodec.Receive(ws, &person)
	if err != nil {
		fmt.Println("Can't receive")
	} else {

		fmt.Println("Name: " + person.Name)
		for _, e := range person.Emails {
			fmt.Println("An email: " + e)
		}
	}
}

func main() {

	http.Handle("/kgb", websocket.Handler(ReceivePerson))
	err := http.ListenAndServe("127.0.0.1:12345", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
