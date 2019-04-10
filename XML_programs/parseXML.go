/* Parse XML
 */

package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func main() {
	/*xmlString :=
			` <email type="work">
	   j.newmarch@boxhill.edu.au
	 </email>
	</person>`*/
	xmlString := `<person>
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
	fmt.Println("xmlString:", xmlString)
	r := strings.NewReader(xmlString)

	parser := xml.NewDecoder(r)
	depth := 0
	for {
		token, err := parser.Token()
		//fmt.Println("debug 1 token:", token)
		if err != nil {
			break
		}
		switch t := token.(type) {
		case xml.StartElement:
			elmt := xml.StartElement(t)
			/*fmt.Println("debug 2 StartElement:", elmt)
			fmt.Println("debug 2 StartElement.Name.Local:" + elmt.Name.Local + ";")
			fmt.Println("debug 2 StartElement.Name.Space:" + elmt.Name.Space + ";")
			fmt.Println("debug 2 StartElement.Attr:", elmt.Attr)
			for _, a := range elmt.Attr {
				fmt.Println("debug 2 StartElement.Attr a.Name.Local:" + a.Name.Local + ";")
				fmt.Println("debug 2 StartElement.Attr a.Name.Space:" + a.Name.Space + ";")
				fmt.Println("debug 2 StartElement.Attr a.Value:" + a.Value + ";")
			}*/
			name := elmt.Name.Local
			printElmt(name, depth)
			depth++
		case xml.EndElement:
			depth--
			elmt := xml.EndElement(t)
			//fmt.Println("debug 2 EndElement:", elmt)
			name := elmt.Name.Local
			printElmt(name, depth)
		case xml.CharData:
			bytes := xml.CharData(t)
			//fmt.Println("debug 2 CharData:" + string(bytes) + ";")
			printElmt("\""+string([]byte(bytes))+"\"", depth)
		case xml.Comment:
			fmt.Println("debug 2 xml.Comment:")
			printElmt("Comment", depth)
		case xml.ProcInst:
			fmt.Println("debug 2 xml.ProcInst:")
			printElmt("ProcInst", depth)
		case xml.Directive:
			fmt.Println("debug 2 xml.Directive:")
			printElmt("Directive", depth)
		default:
			fmt.Println("Unknown")
		}
	}
}

func printElmt(s string, depth int) {
	for n := 0; n < depth; n++ {
		fmt.Print("  ")
	}
	fmt.Println(s)
}
