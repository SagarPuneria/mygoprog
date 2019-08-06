//https://golang.org/doc/articles/wiki/#tmp_0
package main

import (
	//	"html/template"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	defer func() {
		if err1 := recover(); err1 != nil {
			fmt.Println("!!!!!!!!!!!!!!!!Panic Occured and Recovered in loadPage(), Error Info: ", err1)
		}
	}()
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ioutil.ReadFile, Error:", err)
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err1 := recover(); err1 != nil {
			fmt.Println("!!!!!!!!!!!!!!!!Panic Occured and Recovered in viewHandler(), Error Info: ", err1)
		}
	}()
	title := r.URL.Path[len("/view/"):] //http://127.0.0.1:8082/view/TestPage
	p, err := loadPage(title)
	if err == nil {
		//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body) //Framing output in html form i.e.,<h1>Title</h1><div>Body</div>

		fmt.Fprintf(w, "<h1>%s</h1><p>[<a href=\"/edit/%s\">edit</a>]</p><div>%s</div>", p.Title, p.Title, p.Body) //This html have edit option to redirect url from http://127.0.0.1:8082/view/TestPage to http://127.0.0.1:8082/edit/TestPage
	}
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err1 := recover(); err1 != nil {
			fmt.Println("!!!!!!!!!!!!!!!!Panic Occured and Recovered in editHandler(), Error Info: ", err1)
		}
	}()
	title := r.URL.Path[len("/edit/"):] //http://127.0.0.1:8082/edit/TestPage
	// title name is TestPage
	p, err := loadPage(title)
	if err != nil {
		//If title name is wrong or other than TestPage
		fmt.Println("p:", p) //p: <nil>
		p = &Page{Title: title}
		fmt.Println("p:", p) //p: &{TestPages []}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body) //Framing output in HTML form
}
func main() {
	http.HandleFunc("/view/", viewHandler) //http://127.0.0.1:8082/view/TestPage
	http.HandleFunc("/edit/", editHandler) //http://127.0.0.1:8082/edit/TestPage
	//http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8082", nil)
}
