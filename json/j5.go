package main

import (
	"fmt"
	"reflect"
)

type T struct {
	f1 string "one:`1`"

	f2 string `one:"1"`
	//OR
	//f2 string "one:\"1\""
}

// Even if tag is any string literal (interpreted or raw) then Lookup and Get methods
// will find value for key only if value is enclosed between double quotes
func main() {
	t := reflect.TypeOf(T{})
	f1, _ := t.FieldByName("f1")
	fmt.Println(f1.Tag) // one:`1`
	v, ok := f1.Tag.Lookup("one")
	fmt.Printf("%s, %t\n", v, ok) // , false

	fmt.Println("----------")

	f2, _ := t.FieldByName("f2")
	fmt.Println(f2.Tag) // one:"1"
	v, ok = f2.Tag.Lookup("one")
	fmt.Printf("%s, %t\n", v, ok) // 1, true
}
