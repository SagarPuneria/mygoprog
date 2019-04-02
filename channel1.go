package main

import (
	"fmt"
)

func main() {
	message := make(chan string) // Sending or Receiving any msg on nil channel will be > fatal error: all goroutines are asleep - deadlock! To avoid fatal error we use make function.
	go func() {
		fmt.Println("Inside go func(), message received:", <-message)
		fmt.Println("Inside go func(), len(message):", len(message), ",cap(message):", cap(message)) // len(message): 0 ,cap(message):0

		message <- "ping" // Here control holds execution until variable message send any string. If unable send any string it will NOT be deadlock only in GO routine function.
	}()
	message <- "ping" // Here control holds execution until variable message send any string. If won't send any string then it will be deadlock.

	fmt.Println("Inside main func(), len(message):", len(message), ",cap(message):", cap(message)) // len(message): 0 ,cap(message):0

	fmt.Println("Inside main func(), message received:", <-message) // Here control holds execution until variable message received any string. If won't received any string then it will be deadlock.
}

/*
e:\workspace\TDD\gorilla_mux\src\howtogo>go run channel1.go
Inside go func(), message received: ping
Inside go func(), len(message): 0 ,cap(message): 0
Inside main func(), len(message): 0 ,cap(message): 0
Inside main func(), message received: ping
*/
