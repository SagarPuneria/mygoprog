package main

import "fmt"

var isoToUnicodeMap = map[uint8]rune{ // OR map[uint8]int32{
	0xc7: 0x12e,
	0xc8: 0x10c,
	0xca: 0x118,
	// and more
}

func isoBytesToUnicode(bytes []byte) string {
	codePoints := make([]int32, len(bytes))
	for n, v := range bytes {
		unicode, ok := isoToUnicodeMap[v]
		if !ok {
			unicode = int32(v)
		}
		codePoints[n] = unicode
	}
	return string(codePoints)
}

func main() {
	str := "golang"
	fmt.Println("ISO code point:", []byte(str)) //[]uint8(str)
	str = isoBytesToUnicode([]byte(str))        //golang百度一下，你就知道
	fmt.Println("Unicode code point:", []byte(str), str)
}
