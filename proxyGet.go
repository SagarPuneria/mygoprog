/* ProxyGet
 */

package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

//proxyGet.exe "http://172.16.0.14:3128" "http://google.com"
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ", os.Args[0], "http://proxy-host:port http://host:port/page")
		os.Exit(1)
	}
	proxyString := os.Args[1]
	proxyURL, err := url.Parse(proxyString)
	fmt.Println("proxyURL:", proxyURL)
	checkError(err)
	rawURL := os.Args[2]
	url, err := url.Parse(rawURL)
	fmt.Println("rawURL:", url, url.Host)
	checkError(err)

	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	client := &http.Client{Transport: transport}

	request, err := http.NewRequest("GET", url.String(), nil)
	user := "user1"
	passwd := "change"
	strAuth := user + ":" + passwd
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(strAuth))
	proxyAuth := "Basic " + encodedAuth
	request.Header.Set("Proxy-Authorization", proxyAuth)
	request.Header.Set("Connection", "close")

	response, err := client.Do(request)

	checkError(err)
	defer response.Body.Close()
	fmt.Println("Read ok")

	if response.Status != "200 OK" {
		fmt.Println("response.Status:", response.Status)
		os.Exit(2)
	}
	fmt.Println("Reponse ok")

	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println("ioutil.ReadAll ERROR:", err2)
		os.Exit(0)
	}
	fmt.Print("body:", string(body))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
