// https://github.com/gorilla/mux#testing-handlers
// endpoints.go
package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside HealthCheckHandler")
	// A very simple health check.
	w.WriteHeader(http.StatusOK)
	//w.WriteHeader(http.StatusConflict)
	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
	//io.WriteString(w, `{"alive2": true}`)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler)

	err := http.ListenAndServe("localhost:8080", r)
	fmt.Println("We got ERROR", err)
	if err != nil {
		fmt.Println("http.ListenAndServe, Error:", err)
	}
}
