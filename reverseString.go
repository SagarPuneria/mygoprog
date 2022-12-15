package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "$cw&ev124@fer"
	fmt.Println("given str:", str)
	fmt.Println("reverseStringWithOutAffectingSpecialCharater:", reverseStringWithOutAffectingSpecialCharater(str))
	fmt.Println("reverse string:", reverse(str))
}

func isAlphabet(x rune) bool {

	if unicode.IsDigit(x) || unicode.IsLetter(x) {
		return true
	}
	return false
}

func reverseStringWithOutAffectingSpecialCharater(s string) string {
	runes := []rune(s)
	// Initialize left and right pointers
	var r int = len(runes) - 1
	l := 0

	// Traverse string from both ends until
	// 'l' and 'r'
	for l < r {
		// Ignore special characters
		if !isAlphabet(runes[l]) {
			l++
		} else if !isAlphabet(runes[r]) {
			r--

		} else {
			// Both str[l] and str[r] are not spacial
			runes[l], runes[r] = runes[r], runes[l]
			l++
			r--
		}
	}
	return string(runes)
}

func reverse(s string) string {
	// Reverse string using rune is a builtin type(efficient solution)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)

	/*
	// While it works, since strings are immutable(it's very inefficient solution)
	for _, v := range s {
		result = string(v) + result
	}
	*/
}
