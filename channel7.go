/*
How to send signal to all go routines to exit now.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)
	for i := 0; i < 8; i++ {
		go func() {
			/*
				Do some logic
			*/
			fmt.Print(i)
			<-message
			fmt.Println(",Now i can exit")
		}()
	}
	//close(message)
	fmt.Println("Before Wait")
	time.Sleep(2 * time.Second)
	fmt.Println("After Wait")
}

/*
e:\workspace\TDD\gorilla_mux\src\howtogo>go run channel7.go
8Before Wait
5888888After Wait

e:\workspace\TDD\gorilla_mux\src\howtogo>go run channel7.go
Before Wait
88888888After Wait
*/
