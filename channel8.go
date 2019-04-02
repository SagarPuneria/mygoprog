/*
How to send signal to all go routines to exit now.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	message := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			/*
				Do some logic
			*/
			fmt.Print(i)
			<-message
			fmt.Println(",Now i can exit")
		}()
	}
	close(message)
	fmt.Println("Before Wait")
	wg.Wait() // Here control holds execution of current program until all goroutines are finish.
	fmt.Println("After Wait")
}

/*
e:\workspace\TDD\gorilla_mux\src\howtogo>go run channel8.go
Before Wait
8,Now i can exit
8,Now i can exit
8,Now i can exit
8,Now i can exit
8,Now i can exit
8,Now i can exit
8,Now i can exit
8,Now i can exit
After Wait
*/
