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

	values := reflect.Indirect(reflect.ValueOf(T{}))
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := values.Field(i)
		fmt.Println(i, ">>field:", field)
		fmt.Println(i, ">>value:", value)
	}
}

/*OUTPUT:
IN-M-6ZQJG5J:json sagar.puneria$ go run j5.go
one:`1`
, false
----------
one:"1"
1, true
0 >>field: {f1 main string one:`1` 0 [0] false}
0 >>value:
1 >>field: {f2 main string one:"1" 16 [1] false}
1 >>value:
*/
