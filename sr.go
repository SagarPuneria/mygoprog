package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	Arg string
}
type Reply struct {
	Rly string
}
type A int

func (i *A) StringModify(a *Args, r *Reply) error {
	fmt.Println("Inside StringModify")
	a.Arg = "go version is go1.11.4 is latest"
	r.Rly = a.Arg
	return nil
}

func main() {

	a := new(A)
	rpc.Register(a)
	rpc.HandleHTTP()

	err := http.ListenAndServe("127.0.0.1:1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
