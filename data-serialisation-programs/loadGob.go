/* LoadJSON
 */

package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

func main() {
	var person Person
	loadJSON("person.gob", &person)

	fmt.Println("Person", person.String())
}

//func loadJSON(fileName string, key interface{}) { //OR
func loadJSON(fileName string, key *Person) {
	inFile, err := os.Open(fileName)
	checkError("Open", err)
	decoder := gob.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError("Decode:", err)
	inFile.Close()
}

func checkError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
