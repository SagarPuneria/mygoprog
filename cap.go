package main

import "fmt"

func main() {
	//var x = make([]int, 1, 2)
	var x = make([]int, 2, 3)
	x[0] = 31
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	x = append(x, 1)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	x = append(x, 2)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	x = append(x, 3)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	x = append(x, 4)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	x = append(x, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)

	x = append(x, 6)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	x = append(x, 7)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	x = append(x, 8)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	x = append(x, 9)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
	x = append(x, 10)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
