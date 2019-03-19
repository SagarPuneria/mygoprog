package main

import "fmt"

func main() {
	var messages chan string
	go func() {
		messages <- "ping" // fatal error: all goroutines are asleep - deadlock!, Since sending a msg on nil chan
	}()
	msg := <-messages // fatal error: all goroutines are asleep - deadlock!, Since receiving a msg on nil chan
	fmt.Println(msg)
}
