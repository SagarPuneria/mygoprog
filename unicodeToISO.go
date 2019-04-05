package main

import (
	"fmt"
)

var unicodeToISOMap = map[int32]uint8{ // OR map[rune]uint8
	0x12e: 0xc7,
	0x10c: 0xc8,
	0x118: 0xca,
	// plus more
}

/* Turn a UTF-8 string into an ISO 8859 encoded byte array
 */
func unicodeStrToISO(str string) []byte {
	// get the unicode code points
	codePoints := []rune(str)

	// create a byte array of the same length
	bytes := make([]byte, len(codePoints))

	for n, v := range codePoints {
		// see if the point is in the exception map
		iso, ok := unicodeToISOMap[v]
		if !ok {
			// just use the value
			iso = uint8(v)
		}
		bytes[n] = iso
	}
	return bytes
}

func main() {
	str := "golang"
	fmt.Println("Unicode code point:", []rune(str))
	bytes := unicodeStrToISO(str) //golang百度一下，你就知道
	fmt.Println("ISO code point:", bytes, string(bytes))
}
