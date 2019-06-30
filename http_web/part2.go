package main

import (
	"fmt"
	"net/http"
)

/*
http://127.0.0.1:8083/ [OR] http://127.0.0.1:8083
r.URL.Path = /
r.URL.Path[:] = /
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
	fmt.Fprintf(w, "Wow, I did it %s!\n", r.URL.Path[1:pathLenght-1])
	fmt.Fprintf(w, "Wow, I got it %s!\n", r.URL.Path[1:])
	fmt.Println("DEBUG4")
}

func main() {
	fmt.Println("DEBUG1")
	http.HandleFunc("/", handler)
	fmt.Println("DEBUG2")
	http.ListenAndServe(":8083", nil)
	fmt.Println("END")
}

/*OUTPUT:
DEBUG1
DEBUG2
DEBUG3 / $
DEBUG3  $
pathLenght: 1
DEBUG4
DEBUG3 /favicon.ico $
DEBUG3 favicon.ico $
DEBUG4
DEBUG3 /view/TestPage/ $
DEBUG3 view/TestPage/ $
DEBUG4
DEBUG3 /favicon.ico $
DEBUG3 favicon.ico $
DEBUG4
^Csignal: interrupt
root@vadmin01:~/DATA/CODE/http_web#
*/
