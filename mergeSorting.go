package main

import "fmt"

func main() {
	b := []int{2, 1, 4, 9, 6, 7}
	fmt.Println(b)
	mergerSort(b, 0, len(b)-1)
	fmt.Println(b)
}

func mergerSort(a []int, l, h int) {
	var m int
	if l < h {
		m = (l + h) / 2
		mergerSort(a, l, m)
		mergerSort(a, m+1, h)
		merge(a, l, m, h)
	}
}

func merge(a []int, l, m, h int) {
	var i, j, k, p int
	b := make([]int, len(a))
	i = l
	j = m + 1
	k = l
	for (i <= m) && (j <= h) {
		if a[i] <= a[j] {
			b[k] = a[i]
			k++
			i++
		} else {
			b[k] = a[j]
			k++
			j++
		}
	}
	for i <= m {
		b[k] = a[i]
		k++
		i++
	}
	for j <= h {
		b[k] = a[j]
		k++
		j++
	}
	for p = 0; p <= h; p++ {
		a[p] = b[p]
	}
}
