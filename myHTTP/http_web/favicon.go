//https://gist.github.com/nguyendangminh/bd6e1e01df3c6cff139b1609fc1a646c
package main

import (
	"fmt"
	"net/http"
)

/*
http://127.0.0.1:8083/
r.URL.Path[:] = /
r.URL.Path = /
r.URL.Path[1:] =
r.URL.Path[2:] = panic

http://127.0.0.1:8083/view/
r.URL.Path[:] = /view/
r.URL.Path = /view/

http://127.0.0.1:8083/TestPage/
r.URL.Path[:] = /TestPage/
r.URL.Path = /TestPage/

http://127.0.0.1:8083/view/TestPage/
r.URL.Path[:] = /view/TestPage/
r.URL.Path = /view/TestPage/
r.URL.Path[1:] = view/TestPage/
r.URL.Path[2:] = iew/TestPage/
*/

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1faviconHandler", r.URL.Path[:], "$")
	fmt.Println("2faviconHandler", r.URL.Path[len("/"):], "$")
	http.ServeFile(w, r, "path/favicon.ico")
	fmt.Println("3faviconHandler")
	/*
		*b = append(*b, 5, 6)
		(*b)[0] = 9
	*/
}

func agentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1agentHandler", r.URL.Path[:], "$")
	fmt.Println("2agentHandler", r.URL.Path[len("/"):], "$")
	http.ServeFile(w, r, "path/agent.ico")
	fmt.Println("3agentHandler")
}

func shieldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1shieldHandler", r.URL.Path[:], "$")
	fmt.Println("2shieldHandler", r.URL.Path[len("/"):], "$")
	http.ServeFile(w, r, "path/shield.ico")
	fmt.Println("3shieldHandler")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DEBUG3", r.URL.Path[:], "$") // DEBUG3 / $
	// len("/") = 1
	fmt.Println("DEBUG3", r.URL.Path[len("/"):], "$") // DEBUG3  $
	fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[:])
	pathLenght := len(r.URL.Path)
	if pathLenght == 1 {
		fmt.Println("pathLenght:", pathLenght)
		pathLenght = 2
	}
	fmt.Fprintf(w, "Wow, I did it %s!", r.URL.Path[1:pathLenght-1])
	fmt.Println("DEBUG4")
}

func main() {
	fmt.Println("DEBUG1")
	http.HandleFunc("/", handler)
	http.HandleFunc("/icon1.ico", faviconHandler)
	http.HandleFunc("/icon2.ico", agentHandler)
	http.HandleFunc("/icon3.ico", shieldHandler)
	fmt.Println("DEBUG2")
	http.ListenAndServe(":8083", nil)
	fmt.Println("END")
}

/*OUTPUT:
root@vadmin01:~/DATA/CODE/http_web# go run part2.go
DEBUG1
DEBUG2
DEBUG3 / $
DEBUG3  $
DEBUG4
DEBUG3 /favicon.ico $
DEBUG3 favicon.ico $
DEBUG4
^Csignal: interrupt
root@vadmin01:~/DATA/CODE/http_web#
*/
