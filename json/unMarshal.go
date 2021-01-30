package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

func (a *Animal) UnmarshalJSON(b []byte) error {
	fmt.Println(">>>>UnmarshalJSON:", string(b))
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}

	return nil
}

/*func (a *Animal) UnmarshalText(text []byte) error {
	fmt.Println(">>>>UnmarshalText:", string(text))
	switch strings.ToLower(string(text)) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}
	return nil
}*/

func main() {
	blob := `["gopher", "armadillo", "zebra", "unknown", "gopher", "bee", "gopher", "zebra"]`
	var zoo []Animal
	fmt.Println(">>>Before UnmarshalJSON zoo:", zoo)
	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
		log.Fatal(err)
	}
	fmt.Println(">>>After UnmarshalJSON zoo:", zoo)
	census := make(map[Animal]int)
	for _, animal := range zoo {
		census[animal] += 1
	}

	fmt.Printf("Zoo Census:\n* Gophers: %d\n* Zebras:  %d\n* Unknown: %d\n",
		census[Gopher], census[Zebra], census[Unknown])

}
