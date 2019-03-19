package main

import "fmt"

func main() {
	var arr []int
	fmt.Println(arr, arr == nil) // [] true
	arr1 := new([]int)           // new inbuilt function allocates nil slice to *arr1
	//OR var arr1 *[]int = new([]int)
	fmt.Println(*arr1, *arr1 == nil) // [] true
	fmt.Println(arr1, arr1 == nil)   // &[] false
	fmt.Println()
	arr2 := &[]int{} // composite literal's does not allocates nil slice to *arr2
	//OR var arr2 *[]int = &[]int{}
	fmt.Println(*arr2, *arr2 == nil) // [] false
	fmt.Println(arr2, arr2 == nil)   // &[] false
	fmt.Println()
}

/*
 [] true
 [] true
 &[] false

 [] false
 &[] false
*/
