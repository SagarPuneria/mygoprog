/*
Processing HTTP requests with Go is primarily about two things:
(1) ServeMux aka Request Router aka MultiPlexor
(2) Handlers

SERVEMUX
ServeMux = HTTP request router = multiplexor = mux
compares incoming requests against a list of predefined URL paths,
and calls the associated handler for the path whenever a match is found.

HANDLERS
responsible for writing response headers and bodies.
Almost any type ("object") can be a handler, so long as it satisfies the http.Handler interface.
In lay terms, that simply means it must have a ServeHTTP method with the following signature:
ServeHTTP(http.ResponseWriter, *http.Request)

*/

package main

import (
	"fmt"
	"net/http"
)

type person struct {
	fName string
}

func (p *person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("First Name: " + p.fName)
	w.Write([]byte("First Name: " + p.fName))
}

func main() {

	//http.ListenAndServe(":8080", &person{fName: "Jim"})

	/*personOne := &person{fName: "Jim"}
	http.ListenAndServe(":8080", personOne)*/

	/*http.Handle("/", &person{fName: "Jim"})
	http.ListenAndServe(":8080", nil)*/

	/*mux := http.NewServeMux()
	mux.Handle("/", &person{fName: "Jim"})
	http.ListenAndServe(":8080", mux)*/

	//-------------------------------------
	/*http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello universe 2")
		w.Write([]byte("Hello universe 2"))
		//req.Body.Close() // NOT NEEDED
		//req.Close = true // NOT NEEDED FOR SERVER SIDE
	})*/
	http.HandleFunc("/", someFunc)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello universe")
	w.Write([]byte("Hello universe"))
}

/*
We were able to do this because ServeMux also has a ServeHTTP method,
meaning that it too satisfies the Handler interface.

For me it simplifies things to think of a ServeMux as just being a special kind of handler,
which instead of providing a response itself passes the request on to a second handler.

This isn't as much of a leap as it first sounds –
chaining handlers together is fairly commonplace in Go.

from:
http://www.alexedwards.net/blog/a-recap-of-request-handling
*/
