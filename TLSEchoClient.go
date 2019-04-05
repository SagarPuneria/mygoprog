/* TLSEchoClient
 */
package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	certPEMFile, err := os.Open("jan.newmarch.name.pem")
	checkError("os.Open", err)
	rootPEM := make([]byte, 1000) // bigger than the file
	count, err := certPEMFile.Read(rootPEM)
	checkError("certPEMFile.Read", err)
	certPEMFile.Close()
	fmt.Println("rootPEM:", string(rootPEM[:count]))
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(rootPEM[:count])
	if !ok {
		fmt.Println("Failed to parse root certificate")
	}
	config := tls.Config{RootCAs: roots}
	//config := tls.Config{RootCAs: roots, InsecureSkipVerify: true}
	/* InsecureSkipVerify: true >> Used in most of production but NOT secure.
	   In this mode, TLS is susceptible to man-in-the-middle attacks.
	   This should be used only for testing.
	   -------------------------------------
	   If InsecureSkipVerify is not set to TRUE and jan.newmarch.name.pem certificate
	   doesn't have valid IPAddress or doesn't contain any IPAddress then
	   tls.Dial error => x509: cannot validate certificate for
	   127.0.0.1 OR any ipAddress because it doesn't contain any IP SANs*/
	config.PreferServerCipherSuites = true
	config.MinVersion = tls.VersionTLS11
	config.MaxVersion = tls.VersionTLS11

	conn, err := tls.Dial("tcp", service, &config)
	checkError("Dial", err)

	for n := 0; n < 10; n++ {
		fmt.Println("Writing...")
		conn.Write([]byte("Hello " + string(n+48)))

		var buf [512]byte
		n, err := conn.Read(buf[0:])
		checkError("Read", err)

		fmt.Println(string(buf[0:n]))
	}
	os.Exit(0)
}

func checkError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
