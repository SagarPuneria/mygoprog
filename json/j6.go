package main

import (
	"fmt"
	"reflect"
)

func main() {
	type T1 struct {
		f int `json:"foo"`
	}
	type T2 struct {
		f int `json:"bar"`
	}
	t1 := T1{10}
	fmt.Println(t1) // {10}
	var t2 T2
	t2 = T2(t1)
	fmt.Println(t2) // {10}

	t := reflect.TypeOf(t1)
	v := reflect.Indirect(reflect.ValueOf(t1))
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Println(i, ">>field:", field) // 0 >>field: {f main int json:"foo" 0 [0] false}
		fmt.Println(i, ">>value:", value) // 0 >>value: 10
	}
}
