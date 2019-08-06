/* ClientGet
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "http://host:port/page")
		os.Exit(1)
	}
	fmt.Println("os.Args[1]:", os.Args[1])
	url, err := url.Parse(os.Args[1])
	fmt.Println("url:", url)
	checkError(err)

	client := &http.Client{}

	request, err := http.NewRequest("GET", url.String(), nil)
	checkError(err)
	// only accept UTF-8
	request.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
	request.Header.Set("Connection", "close")

	response, err := client.Do(request)
	checkError(err)
	defer response.Body.Close()
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	chSet := getCharset(response)
	fmt.Printf("got charset %s\n", chSet)
	if chSet != "UTF-8" {
		fmt.Println("Cannot handle", chSet)
		os.Exit(4)
	}

	/*var buf [512]byte
	reader := response.Body
	fmt.Println("got body")
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			os.Exit(0)
		}
		fmt.Print(string(buf[0:n]))
	}*/
	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println("ioutil.ReadAll ERROR:", err2)
		os.Exit(0)
	}
	fmt.Print("body:", string(body))

	os.Exit(0)
}

func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	fmt.Println("contentType:", contentType)
	if contentType == "" {
		// guess
		return "UTF-8"
	}
	idx := strings.Index(contentType, "charset")
	if idx == -1 {
		// guess
		return "UTF-8"
	}
	return strings.Trim(contentType[idx:], " ")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
