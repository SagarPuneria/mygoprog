package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	ip "chatserver/ipvalidation"
	st "chatserver/structure"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter IP address to strat the Server:")
	ipString := readInputValue(scanner)
	if ok := ip.CheckIP(ipString); !ok {
		log.Fatal("Invalid IP address")
	}
	fmt.Println("Enter Port address to start the Server:")
	portString := readInputValue(scanner)
	if ok := ip.CheckPort(portString); !ok {
		log.Fatal("Invalid Port address")
	}
	if strings.ContainsAny(ipString, ":") { //for IPv6 address
		ipString = "[" + ipString + "]"
	}
	ln, err := net.Listen("tcp", ipString+":"+portString)
	if err != nil {
		log.Fatal("Error in net.Listen and Error Info:", err)
	}
	defer ln.Close()
	fmt.Println("Server Started...")

	var serverDataStructure []st.ServerStruct
	go pingClient(&serverDataStructure)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error in ln.Accept and Error Info:", err)
			continue
		}
		fmt.Println("Connection Accepted")
		go func() {
			defer conn.Close()
			tempServerDataStructure := getClientDetails(conn)
			if tempServerDataStructure == (st.ServerStruct{}) {
				return
			}
			message := checkDuplicateClient(tempServerDataStructure, serverDataStructure)
			_, err := conn.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Println("Error in conn.Write and Error Info:", err)
				return
			} else if message != "client registered" {
				return
			}
			serverDataStructure = append(serverDataStructure, tempServerDataStructure)
			for {
				tempServerDataStructure = getClientDetails(conn)
				if tempServerDataStructure == (st.ServerStruct{}) {
					return
				} else if tempServerDataStructure.Data == "exit" {
					for i, clientStruct := range serverDataStructure {
						if clientStruct.ClientID == tempServerDataStructure.ClientID && clientStruct.GroupID == tempServerDataStructure.GroupID {
							serverDataStructure = append(serverDataStructure[:i], serverDataStructure[i+1:]...)
							return
						}
					}
				}
				sendReverseStringToRemainingClients(tempServerDataStructure, serverDataStructure)
			}
		}()
	}
}

// Based on following checks like client count and duplicate client within same
// group, checkDuplicateClient return whether client is registered or not
func checkDuplicateClient(tempServerDataStructure st.ServerStruct, serverDataStructure []st.ServerStruct) string {
	if serverDataStructure != nil {
		clientCount := 1
		for _, clientStruct := range serverDataStructure {
			if clientStruct.GroupID == tempServerDataStructure.GroupID {
				clientCount++
				if clientCount > 3 {
					return "client exceeds limit count, try with another group"
				} else if clientStruct.ClientID == tempServerDataStructure.ClientID {
					return "duplicate client id, try with different client id"
				}
			}
		}
	}
	return "client registered"
}

// getClientDetails retrieve client details
func getClientDetails(cn net.Conn) st.ServerStruct {
	// retriving data from buffer and storing in a decoder object
	decoderObject := gob.NewDecoder(cn)

	var tempServerDataStructure st.ServerStruct
	// decodes buffer and unmarshals it into Server Data struct
	err := decoderObject.Decode(&tempServerDataStructure)
	if err != nil {
		fmt.Println("decode error:", err)
		return st.ServerStruct{}
	}
	tempServerDataStructure.ClientAddr = cn.RemoteAddr()
	return tempServerDataStructure
}

// reverseString returns its argument string reversed rune-wise left to right.
func reverseString(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

// sendReverseStringToRemainingClients reverse the received message from one
// particular client and send reverse message to diferent clients of same group.
func sendReverseStringToRemainingClients(tempServerDataStructure st.ServerStruct, serverDataStructure []st.ServerStruct) {
	for _, clientStruct := range serverDataStructure {
		if clientStruct.GroupID == tempServerDataStructure.GroupID {
			if clientStruct.ClientID != tempServerDataStructure.ClientID {
				revStr := reverseString(tempServerDataStructure.Data)
				fmt.Println("server sending data to client", clientStruct.ClientID, "which is in same group and revStr:", revStr)
				conn, err := net.Dial(clientStruct.ClientAddr.Network(), clientStruct.ClientAddr.String())
				if err != nil {
					fmt.Println("Error in net.Dial and Error Info:", err)
					continue
				}
				go func() {
					defer conn.Close()
					_, err := conn.Write([]byte(revStr + "\n"))
					if err != nil {
						fmt.Println("Error in write and Error Info:", err)
					}
				}()
			}
		}
	}
}

// readInputValue which Scans and read a line from Stdin(Console) - return Console input value.
func readInputValue(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

//pingClient will lookinto registered cllient and will ping them in 10 sec interval
func pingClient(serverDataStructure *[]st.ServerStruct) {
	client := time.Tick(60 * time.Second)
	for {
		select {
		case <-client:
			if serverDataStructure != nil {
				var deleteServerStruct []st.ServerStruct
				for _, clientStruct := range *serverDataStructure {
					conn, err := net.Dial(clientStruct.ClientAddr.Network(), clientStruct.ClientAddr.String())
					if err != nil {
						fmt.Println("Error in net.Dial and Error Info:", err)
						fmt.Println("Client", clientStruct.ClientAddr.String(), "address with", clientStruct.ClientAddr.Network(), " network is inactive.\nTherefore server is discarding this client address from database.")
						deleteServerStruct = append(deleteServerStruct, clientStruct)
						continue
					}
					go func() {
						defer conn.Close()
						_, err := conn.Write([]byte("ping" + "\n"))
						if err != nil {
							fmt.Println("Error in write and Error Info:", err)
						}
					}()
				}
				if deleteServerStruct != nil {
					var tmpServerDataStructure []st.ServerStruct
					for _, delClientStruct := range deleteServerStruct {
						for _, clientStruct := range *serverDataStructure {
							if delClientStruct != clientStruct {
								tmpServerDataStructure = append(tmpServerDataStructure, clientStruct)
							} else if len(*serverDataStructure) == 1 && delClientStruct == clientStruct {
								serverDataStructure = nil
								break
							}
						}
					}
					if tmpServerDataStructure != nil {
						*serverDataStructure = tmpServerDataStructure
					}
				}
			}
		}
	}
}
