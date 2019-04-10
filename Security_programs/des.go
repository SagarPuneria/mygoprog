package main

import (
	//"bytes"
	"crypto/des"
	"fmt"
)

func main() {
	key := "12345678" // 8 bytes! this is the DES block size in bytes
	cipher, err := des.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%d bytes NewCipher key with block size of %d bytes\n", len(key), cipher.BlockSize())
	src := []byte("helloooo")
	enc := make([]byte, len(src))
	//var enc [512]byte
	cipher.Encrypt(enc[:], src)
	fmt.Println("Encrpted message:", string(enc[:]), ";")

	var decrypt []byte = make([]byte, len(src))
	//var decrypt [8]byte
	cipher.Decrypt(decrypt[:], enc[:])
	fmt.Println("Decrypted message:", string(decrypt[:]), ";")
	/*OR
	result := bytes.NewBuffer(nil)
	result.Write(decrypt[0:8])
	fmt.Println("result:", string(result.Bytes()), ";")*/
}
