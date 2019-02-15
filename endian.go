package main

import (
	"fmt"
	"unsafe"
)

func main() {
	if getEndian() {
		fmt.Println("little endian")
	} else {
		fmt.Println("big endian")
	}
}

//true = little endian, false = big endian
func getEndian() (ret bool) {
	var i int32 = 0x01234567
	fmt.Printf("%#08x\n", i)
	u := unsafe.Pointer(&i)
	fmt.Println("unsafe.Pointer(&i):", u)
	pb := (*byte)(u)
	bs := *pb
	fmt.Printf("%#02x\n", bs)
	fmt.Println("bs:", bs, 0x67)
	if bs == 0x67 {
		return true
	} else {
		return false
	}

}
