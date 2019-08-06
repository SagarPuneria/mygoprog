package main

import (
	"fmt"
	"log"
	"net/http"

	mx "github.com/gorilla/mux"
)

func specificHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside specificHandler")
	w.Write([]byte("Inside specificHandler"))
}

type person struct {
	fName string
}

func (p *person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("First Name: " + p.fName)
	w.Write([]byte("First Name: " + p.fName))
}

func main() {
	r := mx.NewRouter()
	r.HandleFunc("/specific", specificHandler)
	personOne := &person{fName: "Jim"}
	r.PathPrefix("/").Handler(personOne) // r.PathPrefix("/").Handler(personOne) is not same as r.Handle("/", personOne)
	log.Fatalln(http.ListenAndServe("localhost:8080", r))
}
