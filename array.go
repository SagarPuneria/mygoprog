package main

import "fmt"

func main() {
	/* local variable definition */
	var a = [3]int{1, 2, 3} // Any Array Type is value type
	//var a = []int{1, 2, 3} // Any Slice Type is reference type
	b := a
	fmt.Println(a, b) // [1 2 3] [1 2 3]
	a[0] = 10
	fmt.Println(a, b) // [10 2 3] [1 2 3]
	fmt.Println("===============")
	var x = [3]string{"a", "b", "c"} // Any Array Type is value type
	//var x = []string{"a", "b", "c"} // Any Slice Type is reference type
	fmt.Println("Before fun:", x) // Before fun: [a b c]
	fun(x)
	fmt.Println("After fun:", x) // After fun: [a b c]
}

func fun(b [3]string) { // Any Array Type is value type
	//func fun(b []string) { // Any Slice Type is reference type
	b[0] = "z"
	fmt.Println("Inside fun, b:", b)
}

/*
 [1 2 3] [1 2 3]
 [10 2 3] [1 2 3]
 ===============
 Before fun: [a b c]
 Inside fun, b: [z b c]
 After fun: [a b c]
*/
