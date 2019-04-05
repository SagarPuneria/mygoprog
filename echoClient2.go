/* EchoClient
 */
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/url"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "wss://host:port/kbg2")
		os.Exit(1)
	}
	service := os.Args[1]
	//echoClient2.exe "wss://172.16.0.246:4321/kbg2"

	urlOrigin, err := url.Parse("http://localhost")
	checkError(err)
	fmt.Println("urlOrigin:", urlOrigin)
	urlLocation, err := url.Parse(service)
	checkError(err)

	certPEMFile, err := os.Open("jan.newmarch.name.pem")
	checkError(err)
	rootPEM := make([]byte, 1000) // bigger than the file
	count, err := certPEMFile.Read(rootPEM)
	checkError(err)
	certPEMFile.Close()
	fmt.Println("rootPEM:", string(rootPEM[:count]))
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(rootPEM[:count])
	if !ok {
		fmt.Println("Failed to parse root certificate")
	}
	request := &websocket.Config{Location: urlLocation, Origin: urlOrigin, Version: 13, TlsConfig: &tls.Config{RootCAs: roots}}
	//Version(WebSocket protocol version) must be equal to value of ProtocolVersionHybi13(Protocol Version for Hypertext-Bidirectional is 13)
	request.TlsConfig.MinVersion = tls.VersionTLS11
	request.TlsConfig.MaxVersion = tls.VersionTLS11
	conn, err := websocket.DialConfig(request)
	//conn, err := websocket.Dial(service, "", "http://localhost") // WithOut TlsConfig
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

//https://studygolang.com/resources/1552
//http://andrewwdeane.blogspot.com/2013/01/gobing-down-secure-websockets.html
