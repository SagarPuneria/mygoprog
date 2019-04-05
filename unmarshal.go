/* Unmarshal
 */

package main

import (
	"encoding/xml"
	"fmt"
	"os"
	//"strings"
)

/*type Name struct {
	Family   string `xml:"family"`
	Personal string `xml:"personal"`
}*/

type Email struct {
	//XMLName xml.Name `xml:"email"`
	Type    string `xml:"type,attr"`
	Address string `xml:",chardata"`
}

type Person struct {
	XMLName xml.Name `xml:"person"`
	//Name       Name     `xml:"name"`
	FamilyName   string  `xml:"name>family"`
	PersonalName string  `xml:"name>personal"`
	Email        []Email `xml:"email"`
}

func main() {
	str := `<?xml version="1.0" encoding="utf-8"?>
 <person>
   <name>
	 <family> Newmarch </family>
	 <personal> Jan </personal>
   </name>
   <email type="personal">
	 jan@newmarch.name
   </email>
   <email type="work">
	 j.newmarch@boxhill.edu.au
   </email>
 </person>`

	var person Person

	err := xml.Unmarshal([]byte(str), &person)
	checkError(err)

	// now use the person structure e.g.
	fmt.Println("person.XMLName:", person.XMLName)
	fmt.Println("person.XMLName.Space:", person.XMLName.Space)
	fmt.Println("person.XMLName.Local:", person.XMLName.Local)
	fmt.Println("person.FamilyName: \"" + person.FamilyName + "\"")
	fmt.Println("person.PersonalName: \"" + person.PersonalName + "\"")
	//fmt.Println("person.Email[1].XMLName:", person.Email[1].XMLName)
	fmt.Println("person.Email[1].Address: \"" + person.Email[1].Address + "\"")
	fmt.Println("person.Email:", person.Email)
	//fmt.Println("Family name: \"" + person.Name.Family + "\"")

	byteXml, err := xml.Marshal(person)
	checkError(err)
	fmt.Println("string(byteXml):", string(byteXml))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
