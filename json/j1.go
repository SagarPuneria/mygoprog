package main

import "fmt"

type T1 struct {
	f1 string
}
type T2 struct {
	T1
	f2     int64
	f3, f4 float64
}

// A field declaration may be followed by an optional string literal (tag)
func main() {
	t := T2{T1{"foo"}, 1, 2, 3}
	fmt.Println(t.f1)    // foo
	fmt.Println(t.T1.f1) // foo
	fmt.Println(t.f2)    // 1
}
