package main

import (
	"fmt"
	"net/http"
)

type Service interface {
	Sum(http.ResponseWriter, *http.Request)
	//Concat(http.ResponseWriter, *http.Request)
	//ServeHTTP(http.ResponseWriter, *http.Request)
}
type sumConcatService struct {
	a int
	b int
}

type concatService struct {
	a string
	b string
}

func (s sumConcatService) Sum(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DEBUG1")
	fmt.Fprintf(w, "Hi there, Sum of number is %d\n", s.a+s.b)
	fmt.Println("DEBUG2")
}

func (c concatService) Concat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DEBUG1")
	fmt.Fprintf(w, "Hi there, concat is %d\n", c.a+c.b)
	fmt.Println("DEBUG2")
}

type basicService struct {
	t1 string
}

func (s basicService) SumbasicService(a, b int) (int, error) {
	return a + b, nil
}

func (c basicService) ConcatbasicService(a, b string) (string, error) {
	return a + b, nil
}

func (s basicService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	val, _ := s.SumbasicService(1, 2)
	str, _ := s.ConcatbasicService("Sagar", " Puneria")
	fmt.Println("DEBUG1")
	fmt.Fprintf(w, "%s.Hi there, Myself %s and my number is %d\n", s.t1, str, val)
	fmt.Println("DEBUG2")
}

func main() {
	//basic := basicService{t1: "wow"}
	var basic2 Service
	//basic2 = basic
	//http.HandleFunc("/view/", basic2.ServeHTTP)
	//http.HandleFunc("/view/", basic.ServeHTTP)
	sum := sumService{a: 1, b: 2}
	basic2 = sum
	http.HandleFunc("/sum/", basic2.Sum)
	fmt.Println("Begin")
	http.ListenAndServe(":8083", nil)
	fmt.Println("END")
}
