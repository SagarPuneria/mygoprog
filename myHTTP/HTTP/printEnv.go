/* Print Env
 */

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// file handler for most files
	fmt.Println("http.Dir(/var/www):", http.Dir("e:\\workspace\\TDD\\gorilla_mux\\src\\go-network-programming"))
	fileServer := http.FileServer(http.Dir("e:\\workspace\\TDD\\gorilla_mux\\src\\go-network-programming"))
	http.Handle("/", fileServer)

	// function handler for /cgi-bin/printenv
	http.HandleFunc("/cgi-bin/printenv", printEnv)

	// deliver requests to the handlers
	err := http.ListenAndServe(":8080", nil)
	checkError(err)
	// That's it!
}

func printEnv(writer http.ResponseWriter, req *http.Request) {
	env := os.Environ()
	writer.Write([]byte("<h1>Environment</h1>\n<pre>"))
	for _, v := range env {
		fmt.Println("v:", v)
		writer.Write([]byte(v + "\n"))
	}
	writer.Write([]byte("</pre>"))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
