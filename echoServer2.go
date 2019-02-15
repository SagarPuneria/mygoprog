/* EchoServer
 */
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	fmt.Println("Echoing")
	fmt.Println(ws.Config)

	for n := 0; n < 10; n++ {
		msg := fmt.Sprintf("Hello %d", n+48)
		fmt.Println("Sending to client: " + msg)
		//err := websocket.Message.Send(ws, msg) // OR
		err := websocket.Message.Send(ws, []byte(msg))
		if err != nil {
			fmt.Println("Can't send")
			break
		}

		//var reply []byte // OR
		var reply string
		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	addr := os.Args[1]
	//echoServer2.exe "172.16.0.246:1234"

	http.Handle("/kbg2", websocket.Handler(Echo))
	err := http.ListenAndServeTLS(addr, "jan.newmarch.name.pem", "private.pem", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
