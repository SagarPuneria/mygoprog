package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	go func() {
		messages <- "ping"         // Here control holds execution until variable messages send any string. If won't send any string it will be deadlock.
		fmt.Println(len(messages)) // 0
	}()
	time.Sleep(2 * time.Second)
	msg := <-messages // Here control holds execution until variable messages received any string. If won't received any string it will be deadlock.
	fmt.Println(msg)  // ping
}

/*
e:\workspace\TDD\gorilla_mux\src\golang-network-programming>go run channel1.go
0
ping
*/
