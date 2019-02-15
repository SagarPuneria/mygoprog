package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	cName, err := net.LookupCNAME("www.unity.com")
	if err != nil {
		fmt.Println("Error in LookupCNAME:	", err.Error())
		os.Exit(2)
	}
	fmt.Println("cName:", cName)
	addrs, err := net.LookupHost("LMIPL-234")
	if err != nil {
		fmt.Println("Error in LookupHost:	", err.Error())
		os.Exit(2)
	}
	for _, s := range addrs {
		fmt.Println(">>", s)
	}
	os.Exit(0)
}
