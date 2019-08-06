package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	user := "NA"
	passwd := "NA"
	url := "https://api.vistara.io/auth/oauth/token"
	response := PostHttpResponseToGetAuthToken(url, user, passwd, 1)
	fmt.Println(">>FIRST GetAuthToken response:", string(response))
	url = "https://api.vistara.io/api/v2/activate"
	response = PostHttpResponseToActivateDevice(url, user, passwd, 1)
	fmt.Println(">>SECOND ActivateDevice response:", string(response))
	url = "https://api.vistara.io/api/v2/connectionNode/9bfb20cc-3681-4d7f-8049-30640c900e1c"
	response = GetHttpResponseToGetConnectionNode(url, user, passwd, 1)
	fmt.Println(">>THIRD GetConnectionNode response:", string(response))
}

func PostHttpResponseToActivateDevice(url, user, passwd string, index int) []byte {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception Occurred and Recovered in PostHttpResponseToActivateDevice(), Error Info: ", errD)
		}
	}()
	JsonData := `{ "dnsName": "HYDLPT052","hostName": "HYDLPT052","providerType": "","os": "Ubuntu 14.04.5 LTS","requestorType": "AGENT", "tenant": "", "deviceType": "Server","loggedUser": "","currentTimeZone": "330","serialNumber": "97LBZ32","systemUID": "4C4C4544-0037-4C10-8042-B9C04F5A3332","agentPlatform" : "linux","agentVersion" : "4.5.3-6","providerUID": "","cloudMetaData":{"awsAccountId": ""},"pluginName": "","assetTag": "","resourceNetworkInterface":[{ "ipAddressType": "",  "macAddress": "44:a8:42:ef:db:11",  "default": 0,  "ipAddress": ""},{ "ipAddressType": "",  "macAddress": "94:65:9c:2e:46:0e",  "default": 1,  "ipAddress": "172.24.102.102"},{ "ipAddressType": "",  "macAddress": "16:0b:01:21:26:96",  "default": 0,  "ipAddress": "192.168.122.1"},{ "ipAddressType": "",  "macAddress": "02:42:77:c9:96:72",  "default": 0,  "ipAddress": "172.19.0.1"},{ "ipAddressType": "",  "macAddress": "02:42:9e:9c:9b:59",  "default": 0,  "ipAddress": "172.18.0.1"},{ "ipAddressType": "",  "macAddress": "02:42:5b:4f:59:72",  "default": 0,  "ipAddress": "172.17.0.1"}]}`
	req, err1 := http.NewRequest("POST", url, strings.NewReader(JsonData))
	if err1 != nil {
		fmt.Println(": http.NewRequest ERROR: ", err1)
		return nil
	}
	req.Header.Set("Authorization", "Bearer 0ccaf7a9-76d3-40b8-9132-79bf9cff2388")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if user != "NA" && passwd != "NA" {
		req.SetBasicAuth(user, passwd)
	}
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true, //(OR)req.Header.Set("Connection", "close") //(OR) req.Close = true
	}
	client := &http.Client{Transport: tr}
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

func PostHttpResponseToGetAuthToken(url, user, passwd string, index int) []byte {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception Occurred and Recovered in PostHttpResponseToGetAuthToken(), Error Info: ", errD)
		}
	}()
	req, err1 := http.NewRequest("POST", url, strings.NewReader("client_secret=jVpj7XHHsvUUjSxBAv6BdhuquW66TR8Ujg4va2c2TUWjWhyfquSCsPkeQjmskKfZ&grant_type=client_credentials&client_id=hbV8Jf4RrCDUW9Sb4PA5CFWh5r3yhXkK"))
	if err1 != nil {
		fmt.Println(": http.NewRequest ERROR: ", err1)
		return nil
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Close = true //(OR) req.Header.Set("Connection", "close") (OR) tr := &http.Transport{DisableKeepAlives: true}
	if user != "NA" && passwd != "NA" {
		req.SetBasicAuth(user, passwd)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
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

func GetHttpResponseToGetConnectionNode(url, user, passwd string, index int) []byte {
	defer func() {
		if errD := recover(); errD != nil {
			fmt.Println("Exception Occurred and Recovered in GetHttpResponseToGetConnectionNode(), Error Info: ", errD)
		}
	}()
	req, err1 := http.NewRequest("GET", url, nil)
	if err1 != nil {
		fmt.Println(": http.NewRequest ERROR: ", err1)
		return nil
	}
	req.Header.Set("Authorization", "Bearer 0ccaf7a9-76d3-40b8-9132-79bf9cff2388")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Connection", "close") //(OR) req.Close = true (OR) tr := &http.Transport{DisableKeepAlives: true}
	if user != "NA" && passwd != "NA" {
		req.SetBasicAuth(user, passwd)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr,
		Timeout: time.Duration(240 * time.Second)}
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
