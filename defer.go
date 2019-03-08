package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Inside defer function")
		if errD := recover(); errD != nil {
			fmt.Println("Exception occurred and recovered in main function, Error Info: ", errD)
		}
	}()

	for i := 0; i < 3; i++ {
		defer fmt.Println("defer, i:", i)
	}
	for i := 5; i < 8; i++ {
		defer func() { fmt.Println("defer function, i:", i) }()
	}
	fmt.Println("Before panic")
	panic("Got panic")
	fmt.Println("After panic")
}

/*
e:\workspace\TDD\gorilla_mux\src\golang-network-programming>go run defer.go
END 1
defer function, i: 8
defer function, i: 8
defer function, i: 8
defer, i: 2
defer, i: 1
defer, i: 0
Inside defer function
Exception occurred and recovered in main function, Error Info:  Got panic

e:\workspace\TDD\gorilla_mux\src\golang-network-programming>
*/
