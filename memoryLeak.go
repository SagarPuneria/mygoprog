package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	memoryLeak()
	//memoryNoLeak()

}

func memoryNoLeak() {
	timeOut := int64(3)
	var wg sync.WaitGroup
	wg.Add(1)
	timer := time.AfterFunc(time.Second*time.Duration(timeOut), func() {
		defer wg.Done()
		fmt.Println("func start")
	})
	fmt.Println("Before wait")
	//time.Sleep(5 * time.Second)
	wg.Wait()
	res := timer.Stop()
	fmt.Println("After wait, res", res)
}

func memoryLeak() {
	timeOut := int64(3000)
	timer := time.NewTimer(time.Millisecond * time.Duration(timeOut))
	var wg sync.WaitGroup
	wg.Add(1)
	//You are leaking a goroutine. The goroutine created in line 64 will never exit as timer.C will
	//never be closed and thus the goroutine holds onto some variables forever.
	go func(timer *time.Timer) {
		defer wg.Done()
		fmt.Println("Go routine start")
		for i := range timer.C { //range loop over a channel runs until the channel is closed (or break is used)
			fmt.Println("i:", i)
		}
		fmt.Println("Go routine exit")
	}(timer)
	fmt.Println("Before wait")
	//time.Sleep(3 * time.Second)
	wg.Wait()
	fmt.Println("After wait")
}
