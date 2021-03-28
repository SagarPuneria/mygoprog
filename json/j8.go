package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Dummy struct {
	Name      string `json:"name"`
	Number    int64
	Pointer   *string `json:"-"`  // we want to ignore JSON for this one
	Nopointer string  `json:"-,"` // NoPointer appears in JSON as key "-".
}

func main() {
	data := []byte(`
            {
                "name": "Mr Dummy",
                "NuMbeR": 4,
				"pointer": "yes",
				"Nopointer": "no"
            }
        `)

	var dummy Dummy
	err := json.Unmarshal(data, &dummy)
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
		os.Exit(1)
	}

	// we want to print the field names as well
	fmt.Printf("struct %+v\n---\n", dummy)

	b, err := json.Marshal(dummy)
	if err != nil {
		fmt.Printf("An error occured: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("json:", string(b))
}

/*
go run j8.go
struct {Name:Mr Dummy Number:4 Pointer:<nil> Nopointer:}
---
json: {"name":"Mr Dummy","Number":4,"-":""}
*/
