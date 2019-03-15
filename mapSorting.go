package main

import "fmt"

func main() {
	fmt.Println(map[interface{}]string{
		3:   "3",
		"a": "a",
		2:   "2",
		"c": "c",
		1:   "1",
		"b": "b",
	})
}

// go.12.1 version output: map[1:1 2:2 3:3 a:a b:b c:c]
// go.12 version output: map[3:3 a:a 2:2 c:c 1:1 b:b]
