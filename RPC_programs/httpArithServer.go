/**
* ArithServer
 */

package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Arguments struct {
	A, B int
}

type QuotientRemainder struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Arguments, reply *int) error {
	fmt.Println("Inside Multiply")
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Arguments, quo *QuotientRemainder) error {
	fmt.Println("Inside Divide")
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {

	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe("127.0.0.1:1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
