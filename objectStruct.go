package main

import (
	"fmt"
)

type interfaceB interface {
	object()
}

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
	a := A{} // var a A
	// (OR)
	//a := &A{}
	fmt.Println("a:", a)
	a.name = "sagar"
	fmt.Println("a.name:", a.name)
	a.object()
	fmt.Println("a.name:", a.name)

	var b interfaceB
	b = &a // In case a := A{} // var a A
	// b = a // In case a := &A{}
	// In interface type, If the method is receiver pointer, then object must also be pointer.
	b.object()
}
