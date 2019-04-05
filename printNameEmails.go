/**
 * PrintNameEmails
 */

package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name   string
	Emails []string
}

const templ = `{{$name := .Name}}
{{range .Emails}}
	Name is {{$name}}, email is "{{.}}"
{{end}}`

func main() {
	person := Person{
		Name:   "jan",
		Emails: []string{"jan@newmarch.name", "jan.newmarch@gmail.com"},
	}

	/*t := template.New("Person template")
	t, err := t.Parse(templ)
	*/
	t, err := template.New("Person template").Parse(templ)
	checkError("Parse", err)

	err = t.Execute(os.Stdout, person)
	checkError("Execute", err)
}

func checkError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
