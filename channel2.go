package main

import "fmt"

func main() {
	messages := make(chan string)

	// messages <- "ping" // fatal error: all goroutines are asleep - deadlock!
	//OR
	fmt.Println(<-messages) // fatal error: all goroutines are asleep - deadlock!
}
