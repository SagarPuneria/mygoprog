/* MD5Hash
 */

package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	hash := md5.New()
	bytes := []byte("hi")
	m, err := hash.Write(bytes)
	fmt.Println("m:", m, ",err:", err)
	hashValue := hash.Sum(nil)
	fmt.Println("hashValue:", hashValue, string(hashValue))
	hashSize := hash.Size() // hashSize is always 16 bytes
	fmt.Println("hashSize:", hashSize)
	for n := 0; n < hashSize; n += 4 {
		var val uint32
		//fmt.Println(hashValue[n], ",hashValue[n])<<24:", uint32(hashValue[n])<<24, uint32(hashValue[n]))
		val = uint32(hashValue[n])<<24 +
			uint32(hashValue[n+1])<<16 +
			uint32(hashValue[n+2])<<8 +
			uint32(hashValue[n+3])
		fmt.Printf("%x, %v\n", val, val)
	}
	fmt.Println()
}
