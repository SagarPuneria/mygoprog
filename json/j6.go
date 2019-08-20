package main

import (
	"fmt"
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
}
