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
	//close(message)
	fmt.Println("Before Wait")
	wg.Wait() // fatal error: all goroutines are asleep - deadlock! [Go routines never exit, because it is waiting at message channel to receive string.]
	fmt.Println("After Wait")
}

/*
e:\workspace\TDD\gorilla_mux\src\howtogo>go run channel9.go
Before Wait
88888888fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc000056078)
        C:/Go/src/runtime/sema.go:56 +0x40
sync.(*WaitGroup).Wait(0xc000056070)
        C:/Go/src/sync/waitgroup.go:130 +0x6c
main.main()
        e:/workspace/TDD/gorilla_mux/src/howtogo/channel9.go:29 +0x146

goroutine 19 [chan receive]:
*/
