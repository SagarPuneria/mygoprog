package main

import "fmt"

func main() {
	b := []int{2, 1, 4, 9, 6, 7}
	fmt.Println(b)
	sort(b)
	fmt.Println(b)
}

// Implements Bubble Sort in Go: O(n^2)
func sort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j <= len(a)-1; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

// Implements Optimized Bubble Sort in Go: Best Case (O(n)))
func OptimizedBubbleSort(arr []int) {
	length := len(arr)
	var i, j int
	var swapped bool
	i = 0
	for i < length-1 {
		swapped = false
		j = i + 1
		for j < length {
			if arr[j] < arr[i] {
				temp := arr[j]
				arr[j] = arr[i]
				arr[i] = temp
				swapped = true
			}
			j += 1
		}
		i += 1
		if !swapped {
			return
		}
	}
	return
}

// Implements Optimized Bubble Sort: O(n^2)
func InsertionSort(arr []int) {
	length := len(arr)
	var i, j int
	i = 1
	for i < length {
		j = i - 1
		insertElement := arr[i]
		for j >= 0 {
			if insertElement < arr[j] {
				arr[j+1] = arr[j]
			} else {
				break
			}
			j -= 1
		}
		arr[j+1] = insertElement
		i += 1
	}
}

// Implements Selection Sort in Go: O(n^2)
func SelectionSort(arr []int) {
	length := len(arr)
	var i, j, minIndex int
	i = 0
	for i < length-1 {
		minIndex = i
		j = i + 1
		for j < length {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
			j += 1
		}
		if minIndex != i {
			temp := arr[i]
			arr[i] = arr[minIndex]
			arr[minIndex] = temp
		}
		i += 1
	}
}
