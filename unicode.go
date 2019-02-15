package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	var msg [512]byte
	msg[0] = 41
	msg[1] = 40
	msg[2] = 50
	msg[3] = 0
	msg[4] = 43
	msg[5] = 13
	msg[6] = 42
	msg[7] = 37
	//len1 := 8
	v := string(msg[:8])
	fmt.Println(">>>msg:", v)
	fmt.Println("len(v)", len(v))
	fmt.Println(v)
	fmt.Println("================")
	str := "golang百度一下，你就知道" //"golang.org"
	fmt.Println("str:", str)
	for i, v := range []byte(str) {
		fmt.Println(i, "v:", v, ",string(v):", string(v))
	}
	fmt.Println("----------")
	for i, r := range []rune(str) {
		fmt.Println(i, "rune:", r, ",", string(r), len(string(r)))
	}
	println("String length", len([]rune(str)))
	println("Byte length", len(str))
	fmt.Println("================")

	uint16s := utf16.Encode([]rune(str))
	fmt.Println("uint16s:", uint16s, ",len(uint16s):", len(uint16s))
	runes := utf16.Decode(uint16s)

	fmt.Println("string(runes):", string(runes), ",runes:", runes, ",len(runes)", len(runes))
}

/*
	3%18/3*3 = (((3%18)/3)*3) = 3
	3+18/2*3 = 3 + (18/2)*3 = 30
*/
