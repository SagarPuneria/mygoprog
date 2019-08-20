package main

import (
	"fmt"
	"reflect"
)

type T struct {
	f1 string ``
	f2 string
}

// Setting up empty tag has the same effect as not using tag at all
func main() {
	t := reflect.TypeOf(T{})
	f1, _ := t.FieldByName("f1")
	fmt.Printf("%q\n", f1.Tag) // ""
	f2, _ := t.FieldByName("f2")
	fmt.Printf("%q\n", f2.Tag) // ""
}
