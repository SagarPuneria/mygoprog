package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	log.Fatalln(http.ListenAndServe("localhost:8080", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside HomeHandler")
	w.Write([]byte("Inside HomeHandler"))
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside ProductsHandler")
	w.Write([]byte("Inside ProductsHandler"))
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside ArticlesHandler")
	w.Write([]byte("Inside ArticlesHandler"))
}
