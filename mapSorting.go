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
	// go.12.1 version output: map[1:1 2:2 3:3 a:a b:b c:c]
	// go.12 version output: map[3:3 a:a 2:2 c:c 1:1 b:b]

	fmt.Println(map[string]interface{}{
		"3": 3,
		"a": "a",
		"2": 2,
		"c": "c",
		"1": 1,
		"b": "b",
	})

	fmt.Println(map[int]string{
		3: "a",
		2: "2",
		1: "1",
	})

	fmt.Println(map[string]int{
		"a": 3,
		"2": 2,
		"1": 1,
	})

}

/*
>go run mapSorting.go
map[1:1 2:2 3:3 a:a b:b c:c]
map[1:1 2:2 3:3 a:a b:b c:c]
map[1:1 2:2 3:a]
map[1:1 2:2 a:3]
*/
