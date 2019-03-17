package main

import "fmt"

var reverseOrder string

func recursiveReverseFibo(lastElement, secondLastElement int) string {
	res := lastElement - secondLastElement
	reverseOrder += fmt.Sprintf("%v > ", lastElement)
	if res != 0 {
		recursiveReverseFibo(secondLastElement, res)
	} else {
		reverseOrder += fmt.Sprintf("%v > %v", secondLastElement, res)
	}
	return reverseOrder
}

func iterativeFibonacci(x int) {
	a, b := 0, 1
	for a <= x {
		fmt.Println(a)
		a, b = b, a+b
	}
}

// RecursiveAddLastTwoFibonacciNumber Fibonaacci O(2^n) Exponential runtime
func RecursiveAddLastTwoFibonacciNumber(num int) int {
	if num < 3 {
		return 1
	} else {
		return RecursiveAddLastTwoFibonacciNumber(num-1) + RecursiveAddLastTwoFibonacciNumber(num-2)
	}
}

// IterativeAddLastTwoFibonacciNumber Fibonaacci O(n) Linear runtime
func IterativeAddLastTwoFibonacciNumber(num int) int {
	if num < 3 {
		return 1
	}
	sum := 0
	first := 1
	second := 1
	count := 3
	for count <= num {
		sum = first + second
		first = second
		second = sum
		count += 1
	}
	return sum
}

func main() {
	iterativeFibonacci(15)
	series := recursiveReverseFibo(13, 8)
	fmt.Println("Reverse order fibonacci series", series)
	fmt.Println(RecursiveAddLastTwoFibonacciNumber(7)) // 5 + 8 = 13
	fmt.Println(IterativeAddLastTwoFibonacciNumber(7)) // 5 + 8 = 13
}
