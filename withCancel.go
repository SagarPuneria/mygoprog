package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	withCancel()
	f := func() <-chan int {
		dst := make(chan int)
		i := 1
		fmt.Println(">> Inside f")
		go func() {
			for {
				select {
				case dst <- i:
					i++
					fmt.Println("Inside go func", i)
				}
			}
		}()
		return dst
	}
	fmt.Println("f:", f())
	fmt.Println("f:", f())
	/*for j := range f() {
		fmt.Println("j:", j)
		if j == 5 {
			break
		}
	}*/
	time.Sleep(5 * time.Second)
}

func withCancel() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	fmt.Println("1 debug")
	gen := func(ctx context.Context) <-chan int {
		fmt.Println("###### 4 ctx 1 debug")
		dst := make(chan int)
		n := 1
		go func() {
			fmt.Println("@@@@@@@@ 4 ctx 2 debug")
			for {
				fmt.Println("4 ctx inside for loop debug")
				select {
				case <-ctx.Done():
					fmt.Println("7 debug")
					return // returning not to leak the goroutine
				case dst <- n:
					fmt.Println("4 ctx 3 debug")
					n++
				}
			}
			fmt.Println("After for loop debug")
		}()
		fmt.Println("After go routine debug")
		return dst
	}

	fmt.Println("2 debug")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	fmt.Println("3 debug")
	for n := range gen(ctx) {
		fmt.Println("4 debug")
		fmt.Println(n)
		if n == 5 {
			fmt.Println("5 debug")
			break
		}
	}
	fmt.Println("6 debug")
}
