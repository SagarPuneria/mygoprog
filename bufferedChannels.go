package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	fmt.Println("debug1 len(ch):", len(ch))
	ch <- 9
	fmt.Println("debug2 after sending ch <- 9, len(ch):", len(ch))
	ch <- 8
	fmt.Println("debug3 after sending ch <- 8, len(ch):", len(ch))
	fmt.Println("<-ch:", <-ch)
	fmt.Println("debug4 after reading <-ch, len(ch):", len(ch))
	fmt.Println("<-ch:", <-ch)
	fmt.Println("debug5 after reading <-ch, len(ch):", len(ch), ",cap(ch):", cap(ch))
}