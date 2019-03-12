package main

import "fmt"

func main() {
	var messages chan string
	fmt.Println("Before close") // Before close
	close(messages)             // panic: close of nil channel
	fmt.Println("Before close")
}
