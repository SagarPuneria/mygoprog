package main

import "fmt"

func main() {
	/* local variable definition */
	var a = []int{1, 2, 3} // Any Slice Type is reference type
	b := a
	fmt.Println(a, b) // [1 2 3] [1 2 3]
	a = append(a, 5)
	a[0] = 10 // OR b[0] = 10
	fmt.Println(a, b) // [10 2 3 5] [1 2 3]

	var sl = []int{1, 2, 3}
	fmt.Println("sl:", sl) // sl: [1 2 3]
	fun(&sl)
	fmt.Println("sl:", sl) // sl: [9 2 3 5 6]
}
func fun(b *[]int) { // Any Slice Type is reference type
	(*b)[0] = 9
	*b = append(*b, 5, 6)
}
/*
[1 2 3] [1 2 3]
[10 2 3 5] [1 2 3]
sl: [1 2 3]
sl: [9 2 3 5 6]
*/