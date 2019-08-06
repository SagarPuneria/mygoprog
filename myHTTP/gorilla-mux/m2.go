package main

import (
	"fmt"
	"log"
	"net/http"

	mx "github.com/gorilla/mux"
)

func main() {
	r := mx.NewRouter()
	r.HandleFunc("/products/{key}", ProductsHandler)
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticlesHandler)
	log.Fatalln(http.ListenAndServe("localhost:8080", r))
}

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside ArticlesCategoryHandler")
	vars := mx.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
	w.Write([]byte("Inside ArticlesCategoryHandler"))
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside ProductsHandler")
	vars := mx.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Key: %v\n", vars["key"])
	w.Write([]byte("Inside ProductsHandler"))
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside ArticlesHandler")
	vars := mx.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category id: %v\n", vars["id"])
	w.Write([]byte("Inside ArticlesHandler"))
}
