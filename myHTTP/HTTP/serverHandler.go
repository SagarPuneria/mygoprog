/* ServerHandler
 */

package main

import (
	"fmt"
	"net/http"
)

func main() {

	myHandler := http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		fmt.Println("Inside HandlerFunc and http.StatusNoContent:", http.StatusNoContent)
		// Just return no content - arbitrary headers can be set, arbitrary body
		rw.WriteHeader(http.StatusNoContent)
	})

	http.ListenAndServe("127.0.0.1:8080", myHandler)
}
