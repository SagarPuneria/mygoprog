package main

import "fmt"

func main() {
	messages := make(chan string)
	close(messages)

	//messages <- "ping" // panic: send on closed channel
	// OR
	msg := <-messages // Here control doesn't holds execution, since channel is closed.
	fmt.Println(msg)  // Empty string
}
