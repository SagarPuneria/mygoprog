//https://golang.org/doc/articles/wiki/#tmp_0
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
http://127.0.0.1:8082/
404 page not found[Here it will not invoke viewHandler(){}]

http://127.0.0.1:8082/TestPage/
404 page not found[Here it will not invoke viewHandler(){}]

http://127.0.0.1:8082/view/ [OR] http://127.0.0.1:8082/view
r.URL.Path = /view/
r.URL.Path[:] = /view/
r.URL.Path[1:] = view/
r.URL.Path[len("/view/"):] =
r.URL.Path[len("/views/"):] = panic
r.URL.Path[len("/TestPage/"):] = panic

http://127.0.0.1:8082/view/TestPage
r.URL.Path[:] = /view/TestPage
r.URL.Path[len("/view/"):] = TestPage

http://127.0.0.1:8082/view/TestPage/
r.URL.Path[:] = /view/TestPage/
r.URL.Path[len("/view/"):] = TestPage/
*/
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ioutil.ReadFile, Error:", err)
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DEBUG2", r.URL.Path[:]) //DEBUG2 /view/
	title := r.URL.Path[len("/view/"):]
	fmt.Println("title:", title) //title:
	//	title = r.URL.Path[len("/TestPage/"):] //panic
	//	fmt.Println("title:",title)
	p, err := loadPage(title) //If you want to get title as 'TestPage' then url must be http://127.0.0.1:8082/view/TestPage
	//	p, _ := loadPage("TestPage")
	if err == nil {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body) //Framing output in html form i.e.,<h1>%s</h1><div>%s</div>
	}
}
func main() {
	http.HandleFunc("/view/", viewHandler)
	fmt.Println("DEBUG")
	http.ListenAndServe(":8082", nil)
}
