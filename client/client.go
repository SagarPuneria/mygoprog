package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	ip "chatserver/ipvalidation"
	st "chatserver/structure"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Server IP address you want to connect:")
	ipString := readInputValue(scanner)
	if ok := ip.CheckIP(ipString); !ok {
		log.Fatal("Invalid IP address")
	}
	fmt.Println("Enter Server Port address you want to connect:")
	portString := readInputValue(scanner)
	if ok := ip.CheckPort(portString); !ok {
		log.Fatal("Invalid Port address")
	}
	if strings.ContainsAny(ipString, ":") { //for IPv6 address
		ipString = "[" + ipString + "]"
	}
	conn, err := net.Dial("tcp", ipString+":"+portString)
	if err == nil {
		fmt.Println("Connected to Server.")
	} else {
		log.Fatal("Error in net.Dial and Error Info:", err)
	}
	defer conn.Close()
	cID, gID := getID(scanner)
	msg := st.ID{ClientID: cID, GroupID: gID}
	sendClientDetails(conn, msg)

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error in conn.Read and Error Info:", err)
		return
	} else if strings.TrimSpace(message) != "client registered" {
		fmt.Println("client is not register and message from server:", message)
		return
	}
	fmt.Println("Message from server: " + message)

	lAddr := conn.LocalAddr()
	go listenClientConnection(lAddr)

	greeting := "Enter your choice\n\t1. Send message\n\t2. Exit"
	for {
		fmt.Println(greeting)
		choiceInteger, err := strconv.Atoi(readInputValue(scanner))
		if err != nil {
			fmt.Println("Error in choice input and Error Info:", err)
			fmt.Println("Enter correct integer number")
			continue
		}
		switch choiceInteger {
		case 1:
			fmt.Println("Enter message to be send to server")
			msg2 := st.DataID{msg, readInputValue(scanner)}
			sendClientDetails(conn, msg2)
		case 2:
			msg2 := st.DataID{msg, "exit"}
			sendClientDetails(conn, msg2)
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
			continue
		}
	}
}

// getID read console input and return clientID and groupID
func getID(scanner *bufio.Scanner) (clientIDInteger, groupIDInteger int) {
	fmt.Println("Enter ClientID")
	clientIDInteger, cIDerr := strconv.Atoi(readInputValue(scanner))
	if cIDerr != nil {
		log.Fatal("Error in ClientID input, Error Info:", cIDerr)
	}
	fmt.Println("Enter GroupID")
	groupIDInteger, gIDerr := strconv.Atoi(readInputValue(scanner))
	if gIDerr != nil {
		log.Fatal("Error in GroupID input, Error Info:", gIDerr)
	}
	return clientIDInteger, groupIDInteger
}

// sendClientDetails send client details which is encoded into bytes to server
func sendClientDetails(conn io.Writer, msg interface{}) {
	// create a encoder object
	encoderObject := gob.NewEncoder(conn)
	switch msg.(type) {
	case st.ID:
		// encode buffer and marshal it into a encoderObject
		err := encoderObject.Encode(msg.(st.ID))
		if err != nil {
			log.Fatal("Encode error in st.ID, Error Info:", err)
		}
	case st.DataID:
		// encode buffer and marshal it into a encoderObject
		err := encoderObject.Encode(msg.(st.DataID))
		if err != nil {
			log.Fatal("Encode error in st.DataID, Error Info:", err)
		}
	}
}

// listenClientConnection listen on its own client address.
// After server connection accepted it is ready to receive message from server.
func listenClientConnection(lAddr net.Addr) {
	ln, err := net.Listen(lAddr.Network(), lAddr.String())
	if err != nil {
		log.Fatal("Error in net.Listen and Error Info:", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error in ln.Accept and Error Info:", err)
		}
		go func() {
			defer conn.Close()
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Fatal("Error in conn.Read and Error Info:", err)
			}
			if strings.TrimSpace(message) != "ping" {
				fmt.Println("Reverse message from server: " + message)
			}
		}()
	}
}

// readInputValue which Scans and read a line from Stdin(Console) - return Console input value.
func readInputValue(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
