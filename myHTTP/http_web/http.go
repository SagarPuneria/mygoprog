package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	user := "NA"
	passwd := "NA"
	url := "http://127.0.0.1:8082/view/TestPage"
	response := GetHttpResponseToGetConnectionNode(url, user, passwd, 1)
	fmt.Println(">>GetConnectionNode response:", string(response))
}

func GetHttpResponseToGetConnectionNode(url, user, passwd string, index int) []byte {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception Occurred and Recovered in GetHttpResponseToGetConnectionNode(), Error Info: ", errD)
		}
	}()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr,
		Timeout: time.Duration(240 * time.Second)}
	req, err1 := http.NewRequest("GET", url, nil)
	if err1 != nil {
		fmt.Println(": http.NewRequest ERROR: ", err1)
		return nil
	}
	req.Header.Set("Authorization", "Bearer 0ccaf7a9-76d3-40b8-9132-79bf9cff2388")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Connection", "close") //(OR) req.Close = true
	if user != "NA" && passwd != "NA" {
		req.SetBasicAuth(user, passwd)
	}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(": client.Do ERROR: ", err)
		return nil
	}
	defer response.Body.Close()
	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println(": ioutil.ReadAll ERROR:", err2)
		return nil
	}
	return body
}
