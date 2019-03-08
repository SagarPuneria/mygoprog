package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages // Here control holds execution until variable messages received any string. If won't received any string it will be deadlock.
	fmt.Println(msg)  // ping
}
