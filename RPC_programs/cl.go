package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	Arg string
}
type Reply struct {
	Rly string
}

func main() {

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	a := Args{"go version is go1.11.3 is not latest"}
	var r Reply
	// Synchronous call(means it will wait until A.StringModify method get executed)
	err = client.Call("A.StringModify", &a, &r)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("arg:\"%s\"\nreply:\"%s\"\n", a.Arg, r.Rly)
}
