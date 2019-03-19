package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)
	go func() {
		message <- "ping"                                                          // Here control holds execution until variable message send any string. If won't send any string it will be deadlock.
		fmt.Println("len(message):", len(message), ",cap(message):", cap(message)) // len(message): 0 ,cap(message):0
	}()
	time.Sleep(2 * time.Second)
	msg := <-message // Here control holds execution until variable message received any string. If won't received any string it will be deadlock.
	fmt.Println(msg) // ping
	time.Sleep(1 * time.Second)
}

/*
e:\workspace\TDD\gorilla_mux\src\golang-network-programming>go run channel1.go
len(message):0
ping
[OR]
e:\workspace\TDD\gorilla_mux\src\golang-network-programming>go run channel1.go
ping
len(message): 0
*/
