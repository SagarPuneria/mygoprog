package ipvalidation

import (
	"log"
	"net"
	"strconv"
)

//CheckIP validates an IP adderss
func CheckIP(ipaddr string) bool {
	check := net.ParseIP(ipaddr)
	if check.To4() == nil && check.To16() == nil {
		return false
	}
	return true
}

//CheckPort validates a Port address
func CheckPort(portadd string) bool {
	num, err := strconv.Atoi(portadd)
	if err != nil {
		log.Fatal("Error in Port input, Error Info:", err)
	} else if num > 0 && num < 65535 {
		return true
	}
	return false
}
