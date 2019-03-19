package main

import (
	"fmt"
	"time"
)

func main() {
	/* local variable definition */
	var a = []int{1, 2, 3} // Any Slice Type is reference type
	b := a
	fmt.Println(a, b) // [1 2 3] [1 2 3]
	a[0] = 10         // OR b[0] = 10
	a = append(a, 5)
	fmt.Println(a, b) // [10 2 3 5] [10 2 3]
	fmt.Println("===============")
	var x = []string{"a", "b", "c"} //Any Slice Type is reference type
	fmt.Println("Before fun:", x)   // Before fun: [a b c]
	go fun(x)
	x[1] = "y"
	fmt.Println("x:", x) // x: [a y c]
	time.Sleep(2 * time.Second)
	fmt.Println("After fun:", x) // After fun: [z y c]
}
func fun(b []string) { // Any Slice Type is reference type
	time.Sleep(time.Second)
	fmt.Println("b:", b) // b: [a y c]
	b[0] = "z"
	b = append(b, "t")
	fmt.Println("Inside fun b:", b) // Inside fun b: [z y c t]
}

/*
 [1 2 3] [1 2 3]
 [10 2 3 5] [10 2 3]
 ===============
 Before fun: [a b c]
 x: [a y c]
 b: [a y c]
 Inside fun b: [z y c t]
 After fun: [z y c]
*/
