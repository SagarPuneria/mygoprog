package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type A struct{}

func (d A) GobEncode() ([]byte, error) {
	fmt.Println("(d A) GobEncode...")
	return nil, nil
}

func (d *A) GobDecode(data []byte) error {
	fmt.Println("(d *A) GobDecode...")
	return nil
}

type B struct{}

func (d *B) GobEncode() ([]byte, error) {
	fmt.Println("(d *B) GobEncode...")
	return nil, nil
}

func (d *B) GobDecode(data []byte) error {
	fmt.Println("(d *B) GobDecode...")
	return nil
}

func main() {
	AWithPointerData()
	AWithoutPointerData()
	BWithPointerData()
	BWithoutPointerData()
}

func AWithPointerData() {
	buf := &bytes.Buffer{}

	enc := gob.NewEncoder(buf)
	fmt.Println("Before AWithPointerData Encode")
	enc.Encode(A{})
	fmt.Println("After AWithPointerData Encode")

	dec := gob.NewDecoder(buf)
	fmt.Println("Before AWithPointerData Decode")
	dec.Decode(&A{}) // dec.Decode(A{}) > wrong [(d *A)GobDecode method will not invoke because api dec.Decode(A{}) return error]
	fmt.Println("After AWithPointerData Decode")
	fmt.Println("------------------")
}

func AWithoutPointerData() {
	buf := &bytes.Buffer{}

	enc := gob.NewEncoder(buf)
	fmt.Println("Before AWithoutPointerData Encode")
	enc.Encode(A{})
	fmt.Println("After AWithoutPointerData Encode")

	dec := gob.NewDecoder(buf)
	fmt.Println("Before AWithoutPointerData Decode")
	dec.Decode(A{}) //dec.Decode(&A{}) > correct
	fmt.Println("After AWithoutPointerData Decode")
	fmt.Println("------------------")
}

func BWithPointerData() {
	buf := &bytes.Buffer{}

	enc := gob.NewEncoder(buf)
	fmt.Println("Before BWithPointerData Encode")
	enc.Encode(&B{})
	fmt.Println("After BWithPointerData Encode")

	dec := gob.NewDecoder(buf)
	fmt.Println("Before BWithPointerData Decode")
	dec.Decode(&B{}) // dec.Decode(B{}) > wrong [(d *B)GobDecode method will not invoke because api dec.Decode(B{}) return error]
	fmt.Println("After BWithPointerData Decode")
	fmt.Println("------------------")
}

func BWithoutPointerData() {
	buf := &bytes.Buffer{}

	enc := gob.NewEncoder(buf)
	fmt.Println("Before BWithoutPointerData Encode")
	enc.Encode(B{})
	// enc.Encode(B{}) > wrong [(d *B)GobEncode method will not invoke because api dec.Encode(B{}) return error] and
	// api dec.Decode(&B{}) return error since buf is emtpty because enc.Encode(B{}) also return error
	fmt.Println("After BWithoutPointerData Encode")

	dec := gob.NewDecoder(buf)
	fmt.Println("Before BWithoutPointerData Decode")
	dec.Decode(&B{})
	fmt.Println("After BWithoutPointerData Decode")
	fmt.Println("------------------")
}
