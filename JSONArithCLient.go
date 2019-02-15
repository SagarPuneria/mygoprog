/**
* TCPArithClient
 */

package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Arguments struct {
	A, B int
}

type QuotientRemainder struct {
	Quo, Rem int
}

func main() {
	service := "127.0.0.1:1234"

	client, err := jsonrpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := Arguments{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot QuotientRemainder
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
