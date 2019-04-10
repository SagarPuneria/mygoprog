/* ASN.1
 */

package main

import (
	"encoding/asn1"
	"fmt"
	"os"
)

func main() {
	mdata, err := asn1.Marshal(299)
	checkError(err)
	fmt.Println("After marshal mdata: ", mdata, string(mdata[:]))

	var n int
	rest, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)

	fmt.Println("After unmarshal n: ", n)
	fmt.Println("After unmarshal rest: ", rest, string(rest[:]))

	s := "hello"
	mdata, _ = asn1.Marshal(s)

	var newstr string
	asn1.Unmarshal(mdata, &newstr)
	fmt.Println("newstr:", newstr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
