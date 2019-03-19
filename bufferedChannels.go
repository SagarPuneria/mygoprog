package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	fmt.Println("debug1 len(ch):", len(ch), ",cap(ch):", cap(ch)) // debug1 len(ch): 0 ,cap(ch): 2
	ch <- 9
	fmt.Println("debug2 after sending ch <- 9, len(ch):", len(ch)) // debug2 after sending ch <- 9, len(ch): 1
	ch <- 8
	fmt.Println("debug3 after sending ch <- 8, len(ch):", len(ch))                    // debug3 after sending ch <- 8, len(ch): 2
	fmt.Println("<-ch:", <-ch)                                                        // <-ch: 9
	fmt.Println("debug4 after reading <-ch, len(ch):", len(ch))                       // debug4 after reading <-ch, len(ch): 1
	fmt.Println("<-ch:", <-ch)                                                        // <-ch: 8
	fmt.Println("debug5 after reading <-ch, len(ch):", len(ch), ",cap(ch):", cap(ch)) // debug5 after reading <-ch, len(ch): 0 ,cap(ch): 2
}
