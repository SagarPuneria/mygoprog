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

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside ArticleHandler")
	vars := mx.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "id: %v\n", vars["id"])
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
	w.Write([]byte("Inside ArticleHandler"))
}

func main() {
	r := mx.NewRouter()
	s := r.Host("localhost").Subrouter()
	s.HandleFunc("/products", ProductsHandler)
	s.HandleFunc("/products/{key}", ProductHandler)
	s.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	log.Fatalln(http.ListenAndServe("localhost:8080", r))
}
