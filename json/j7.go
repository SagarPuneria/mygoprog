package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	/*type T struct {
		F1 int `orm:"f_1"`
		F2 int `orm:"f_2,omitempty"`
		F3 int `orm:"f_3,omitempty" view:"-"`
		F4 int `orm:"-"`
	}
	//json.Marshal: {"F1":1,"F2":0,"F3":2,"F4":3}
	*/
	type T struct {
		F1 int `json:"f_1"`
		F2 int `json:"f_2,omitempty"`
		F3 int `view:"-" json:"f_3,omitempty"`
		F4 int `json:"-"`
	}
	t := T{F1: 1, F3: 2, F4: 3}
	// OR
	// t := T{1, 0, 2, 3}

	fmt.Println(t) // {1 0 2 3}
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json.Marshal: %s\n", b) // json.Marshal: {"f_1":1,"f_3":2}
	fmt.Println("----------")

	rt := reflect.TypeOf(T{})
	f1, _ := rt.FieldByName("F1")
	fmt.Println("f1.Tag:", f1.Tag) // f1.Tag: json:"f_1"
	v, ok := f1.Tag.Lookup("json")
	fmt.Printf("F1 json value:%s; bool:%t\n", v, ok) // F1 json value:f_1; bool:true
	fmt.Println("----------")

	rv := reflect.Indirect(reflect.ValueOf(t))
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i)
		fmt.Println(i, ">>field:", field)
		fmt.Println(i, ">>value:", value)
	}
	/*
		0 >>field: {F1  int json:"f_1" 0 [0] false}
		0 >>value: 1
		1 >>field: {F2  int json:"f_2,omitempty" 8 [1] false}
		1 >>value: 0
		2 >>field: {F3  int view:"-" json:"f_3,omitempty" 16 [2] false}
		2 >>value: 2
		3 >>field: {F4  int json:"-" 24 [3] false}
		3 >>value: 3
	*/

	f2, _ := rt.FieldByName("F2")
	fmt.Println("f2.Tag:", f2.Tag) // f2.Tag: json:"f_2,omitempty"
	v, ok = f2.Tag.Lookup("json")
	fmt.Printf("F2 json value:%s; bool:%t\n", v, ok) // F2 json value:f_2,omitempty; bool:true
	fmt.Println("----------")

	f3, _ := rt.FieldByName("F3")
	fmt.Println("f3.Tag:", f3.Tag) // f3.Tag: view:"-" json:"f_3,omitempty"
	v, ok = f3.Tag.Lookup("json")
	fmt.Printf("F3 json value:%s; bool:%t\n", v, ok) // F3 json value:f_3,omitempty; bool:true
	v, ok = f3.Tag.Lookup("view")
	fmt.Printf("F3 view value:%s; bool:%t\n", v, ok) // F3 view value:-; bool:true
	fmt.Println("----------")

	f4, _ := rt.FieldByName("F4")
	fmt.Println("f4.Tag:", f4.Tag) // f4.Tag: json:"-"
	v, ok = f4.Tag.Lookup("json")
	fmt.Printf("F4 json value:%s; bool:%t\n", v, ok) // F4 json value:-; bool:true
}

/*
Examples of struct field tags and their meanings:

// Field appears in JSON as key "myName".
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`

// Field is ignored by this package.
Field int `json:"-"`

// Field appears in JSON as key "-".
Field int `json:"-,"`
*/

/*
OUTPUT:
IN-M-6ZQJG5J:json sagar.puneria$ go run j7.go
{1 0 2 3}
json.Marshal: {"f_1":1,"f_3":2}
----------
f1.Tag: json:"f_1"
F1 json value:f_1; bool:true
----------
0 >>field: {F1  int json:"f_1" 0 [0] false}
0 >>value: 1
1 >>field: {F2  int json:"f_2,omitempty" 8 [1] false}
1 >>value: 0
2 >>field: {F3  int view:"-" json:"f_3,omitempty" 16 [2] false}
2 >>value: 2
3 >>field: {F4  int json:"-" 24 [3] false}
3 >>value: 3
f2.Tag: json:"f_2,omitempty"
F2 json value:f_2,omitempty; bool:true
----------
f3.Tag: view:"-" json:"f_3,omitempty"
F3 json value:f_3,omitempty; bool:true
F3 view value:-; bool:true
----------
f4.Tag: json:"-"
F4 json value:-; bool:true
IN-M-6ZQJG5J:json sagar.puneria$
*/
