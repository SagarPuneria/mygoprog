//https://golang.org/doc/articles/wiki/#tmp_0
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
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
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	defer func() {
		if err1 := recover(); err1 != nil {
			fmt.Println("!!!!!!!!!!!!!!!!Panic Occured and Recovered in renderTemplate(), Error Info: ", err1)
		}
	}()
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err1 := recover(); err1 != nil {
			fmt.Println("!!!!!!!!!!!!!!!!Panic Occured and Recovered in viewHandler(), Error Info: ", err1)
		}
	}()
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		//If title name is wrong or other than TestPage
		fmt.Println("viewHandler: If title name is wrong or other than TestPage")
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err1 := recover(); err1 != nil {
			fmt.Println("!!!!!!!!!!!!!!!!Panic Occured and Recovered in editHandler(), Error Info: ", err1)
		}
	}()
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		//If title name is wrong or other than TestPage
		fmt.Println("editHandler: If title name is wrong or other than TestPage")
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err1 := recover(); err1 != nil {
			fmt.Println("!!!!!!!!!!!!!!!!Panic Occured and Recovered in saveHandler(), Error Info: ", err1)
		}
	}()
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8082", nil)
	//http.ListenAndServe("192.168.0.39:8082", nil)
}
