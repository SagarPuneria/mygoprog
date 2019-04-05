/* EchoClient
 */
package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "ws://host:port")
		os.Exit(1)
	}
	service := os.Args[1]
	conn, err := websocket.Dial(service, "", "http://localhost")
	checkError(err)
	for {
		var msg string
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			fmt.Println("websocket.Message.Receive, err:", err)
			if err == io.EOF {
				fmt.Println("graceful shutdown by server")
				break
			}
			fmt.Println("Couldn't receive msg " + err.Error())
			break
		}
		fmt.Println("Received from server: " + msg)
		// return the msg
		err = websocket.Message.Send(conn, msg)
		if err != nil {
			fmt.Println("Couldn't return msg")
			break
		}
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
