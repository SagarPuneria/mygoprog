/**
* TCPArithClient
 */

package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Arguments struct {
	A, B int
}

type QuotientRemainder struct {
	Quo, Rem int
}

func main() {
	service := "127.0.0.1:1234"

	client, err := rpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := Arguments{17, 8}
	var reply int
	// Synchronous client.call(means it will wait until Arith.Multiply method get executed)
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith.Multiply: %d*%d=%d\n", args.A, args.B, reply)

	var quot QuotientRemainder
	// Synchronous client.call(means it will wait until Arith.Divide method get executed)
	/*err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)*/

	// Asynchronous client.Go(means it will not wait until Arith.Divide method get executed)
	divCall := client.Go("Arith.Divide", args, &quot, nil)
	replyCall := <-divCall.Done // it will wait untill Arith.Divide method get executed
	//fmt.Printf("After Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem) // OR
	fmt.Printf("After Arith: %d/%d=%d remainder %d\n", replyCall.Args.(Arguments).A, replyCall.Args.(Arguments).B, replyCall.Reply.(*QuotientRemainder).Quo, replyCall.Reply.(*QuotientRemainder).Rem)
}
