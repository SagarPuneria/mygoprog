package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	message := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
	label:
		for {
			select {
			case msg := <-message:
				fmt.Println("go routine msg:", msg)
			default:
				fmt.Println("Exit from go routine")
				wg.Done()
				//goto label // Iterate for loop
				//break // break keyword only break select statement NOT for loop
				break label // Break for loop
			}
		}
	}()
	message <- "Hello World"
	fmt.Println("Before Wait")
	wg.Wait() // Here control holds execution of current program until all goroutines are finish.
	fmt.Println("After Wait")
	time.Sleep(1 * time.Second)
}

/*
> go run channel10.go
go routine msg: Hello World
Before Wait
Exit from go routine
After Wait
*/
