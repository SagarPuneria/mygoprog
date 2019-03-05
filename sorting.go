package main

import "fmt"

func main() {
	b := []int{2, 1, 4, 9, 6, 7}
	fmt.Println(b)
	sort(b)
	fmt.Println(b)
}

func sort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j <= len(a)-1; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
