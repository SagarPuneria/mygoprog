package main

import (
	"fmt"
)

func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

func MergeSort(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	n := len(s) / 2
	l := MergeSort(s[:n])
	r := MergeSort(s[n:])
	return Merge(l, r)
}

func main() {
	s := []int{9, 4, 3, 6, 1, 2, 10, 5, 7, 8}
	fmt.Printf("%v\n%v\n", s, MergeSort(s))
}

/*package main

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
}*/
