package main

import "fmt"

func fibonacci(x int) {
	a, b := 0, 1
	for a <= x {
		fmt.Println(a)
		a, b = b, a+b
	}
}

func main() {
	fibonacci(144)
}
