/* Head
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	url := os.Args[1]

	response, err := http.Head(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	defer response.Body.Close()

	fmt.Println(response.Status)
	for k, v := range response.Header {
		fmt.Println(k+":", v)
	}
	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println("##################ioutil.ReadAll ERROR:", err2)
		os.Exit(0)
	}
	fmt.Print("$$$$$$$$$$$$$body:", string(body))
	os.Exit(0)
}
