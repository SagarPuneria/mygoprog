package main

import "fmt"

func main() {
	/* local variable definition */
	var a = [3]int{1, 2, 3} // Any Array Type is value type
	//var a = []int{1, 2, 3} // Any Slice Type is reference type
	b := a
	fmt.Println(a, b)
	a[0] = 10
	fmt.Println(a, b)
	fmt.Println("===============")
	//var x = [3]string{"a", "b", "c"} // Any Array Type is value type
	var x = []string{"a", "b", "c"} // Any Slice Type is reference type
	fmt.Println("Before fun:", x)
	fun(x)
	fmt.Println("Afore fun:", x)
}

//func fun(b [3]string) { // Any Array Type is value type
func fun(b []string) { // Any Slice Type is reference type
	b[0] = "z"
}