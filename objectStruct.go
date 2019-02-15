package main

import (
	"fmt"
)

type A struct {
	name string
}

//https://golang.org/doc/faq#different_method_sets

func (d *A) object() { // Correct way
	//func (d A) object() { // Wrong way
	fmt.Println("(d A) GobEncode...")
	d.name = "sagar 2"
}

func main() {
	//a := A{} // (OR) var a A
	a := &A{}
	fmt.Println("a:", a)
	a.name = "sagar"
	fmt.Println("a.name:", a.name)
	a.object()
	fmt.Println("a.name:", a.name)
}
