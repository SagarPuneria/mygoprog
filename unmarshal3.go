/* Unmarshal
 */

package main

import (
	"encoding/xml"
	"fmt"
	"os"
	//"strings"
)

func main() {
	str := `<?xml version="1.0" encoding="utf-8"?>
	 <cm>
	 <id>Employee ID</id>
	 <reqid>0</reqid>
	 <params>
		 <ps>
			 <p>
				 <name>employee1</name>
				 <name>employee2</name>
				 <name>employee3</name>
			 </p>
			 <list>0</list>
		 </ps>
	 </params>
 </cm>`

	type EmployeeD struct {
		XMLName  xml.Name `xml:"cm"`
		Identity string   `xml:"id"`
		Reqid    string   `xml:"reqid"`
		List     string   `xml:"params>ps>list"`
		Name     []string `xml:"params>ps>p>name"`
	}
	var employeeDetails EmployeeD

	err := xml.Unmarshal([]byte(str), &employeeDetails)
	checkError(err)

	// now use the employee structure e.g.
	fmt.Println("employeeDetails.Identity: \"" + employeeDetails.Identity + "\"")
	fmt.Println("employeeDetails.Reqid: \"" + employeeDetails.Reqid + "\"")
	fmt.Println("employeeDetails.List: \"" + employeeDetails.List + "\"")
	fmt.Println("employeeDetails.Name:", employeeDetails.Name)

	byteXml, err := xml.Marshal(employeeDetails)
	checkError(err)
	fmt.Println("string(byteXml):", string(byteXml))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

/*OUTPUT:
<?xml version="1.0"?>
<cm>
    <id>Employee ID</id>
    <reqid>0</reqid>
    <params>
        <ps>
            <list>0</list>
            <p>
                <name>employee1</name>
                <name>employee2</name>
                <name>employee3</name>
            </p>
        </ps>
    </params>
</cm>
*/
