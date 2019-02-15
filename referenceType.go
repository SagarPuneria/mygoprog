package main

import "fmt"

type A struct {
	a int
}

func (d A) GobEncode() ([]byte, error) {
	d.a = 10
	fmt.Println("(d A) GobEncode...")
	return nil, nil
}
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

	myMap := make(map[string]int) // Map Type is reference type
	myMap["s"] = 1
	fmt.Println("Before myMapfun:", myMap)
	myMapfun(myMap)
	fmt.Println("After myMapfun:", myMap)

	var sl = []int{1, 2, 3}
	fmt.Println("Before fun1, sl:", sl)
	fun1(&sl)
	fmt.Println("After fun1, sl:", sl)

	fmt.Println("=========")
	arr1 := new([]int)               // new imbuilt function allocates nil slice to *arr1
	fmt.Println(*arr1, *arr1 == nil) // [] true
	fmt.Println(arr1, arr1 == nil)   // &[] false
	fmt.Println()
	arr2 := &[]int{}                 // composite literal's does not allocates nil slice to *arr2
	fmt.Println(*arr2, *arr2 == nil) // [] false
	fmt.Println(arr2, arr2 == nil)   // &[] false
	fmt.Println()
}

//func fun(b [3]string) { // Any Array Type is value type
func fun(b []string) { // Any Slice Type is reference type, but append() will not be reflected in main
	b[0] = "z"
	//b = append(b, "f")
}

func myMapfun(m map[string]int) { // Map Type is reference type
	m["z"] = 2
	m["s"] = 3
}

func fun1(ls *[]int) { // Any Slice pointer Type is reference and append() will be reflected in main
	/*
		v := *ls
		v[0] = 9
		*ls = v // OPTIONAL
	*/
	(*ls)[0] = 9
	*ls = append(*ls, 5, 6)
}
