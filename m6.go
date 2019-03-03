package main

import (
	"fmt"
	"log"
	"net/http"

	mx "github.com/gorilla/mux"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside ProductsHandler")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Inside ProductsHandler"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside ProductHandler")
	vars := mx.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Key: %v\n", vars["key"])
	w.Write([]byte("Inside ProductHandler"))
}

func ProductDetailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside ProductDetailsHandler")
	vars := mx.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Key: %v\n", vars["key"])
	w.Write([]byte("Inside ProductDetailsHandler"))
}

func main() {
	r := mx.NewRouter()
	s := r.PathPrefix("/products").Subrouter()
	// "/products/"
	s.HandleFunc("/", ProductsHandler)
	// "/products/{key}"
	s.HandleFunc("/{key}", ProductHandler)
	// "/products/{key}/details"
	s.HandleFunc("/{key}/details", ProductDetailsHandler)
	log.Fatalln(http.ListenAndServe("localhost:8080", r))
}
