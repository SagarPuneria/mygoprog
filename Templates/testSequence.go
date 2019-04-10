/**
 * Sequence.go
 * Copyright Roger Peppe
 */

package main

import (
	"errors"
	"fmt"
	"os"
	"text/template"
)

/*var tmpl = `
{{range $}}
	@{{.}}@
{{end}}`*/

var tmpl = `{{$comma := sequence "" ", "}}
{{range $}}
	{{$comma.Next}}@{{.}}@
{{end}}
{{$comma := sequence "" ", "}}
{{$colour := cycle "black" "white" "red"}}
{{range $}}
	{{$comma.Next}}{{.}} in {{$colour.Next}}
{{end}}`

/*var tmpl = `{{$comma := sequence "" ", "}}
{{range $}}{{$comma.Next}}@{{.}}@{{end}}
{{$comma := sequence "" ", "}}
{{$colour := cycle "black" "white" "red"}}
{{range $}}{{$comma.Next}}{{.}} in {{$colour.Next}}{{end}}
`*/

var fmap = template.FuncMap{
	"sequence": sequenceFunc,
	"cycle":    cycleFunc,
}

func main() {
	t, err := template.New("Template Paser").Funcs(fmap).Parse(tmpl)
	if err != nil {
		fmt.Printf("parse error: %v\n", err)
		return
	}
	err = t.Execute(os.Stdout, []string{"a", "b", "c", "d", "e", "f"})
	if err != nil {
		fmt.Printf("exec error: %v\n", err)
	}
}

type generator struct {
	ss []string
	i  int
	f  func(s []string, i int) string
}

func (seq *generator) Next() string {
	fmt.Println("##Inside (seq *generator) Next() and seq.i:", seq.i)
	s := seq.f(seq.ss, seq.i)
	fmt.Println("##Inside (seq *generator) Next() and s:", s)
	seq.i++
	return s
}

func sequenceGen(ss []string, i int) string {
	fmt.Println("##Inside sequenceGen ss:", ss, "and len(ss):", len(ss), " and i:", i)
	for i, s := range ss {
		fmt.Println("i:", i, "s:", s)
	}
	if i >= len(ss) {
		fmt.Println("##ss[len(ss)-1]:", ss[len(ss)-1])
		return ss[len(ss)-1]
	}
	fmt.Println("##ss[i]:", ss[i])
	return ss[i]
}

func cycleGen(ss []string, i int) string {
	fmt.Println("##Inside cycleGen ss:", ss, " and i:", i)
	return ss[i%len(ss)]
}

func sequenceFunc(ss ...string) (*generator, error) {
	fmt.Println("##Inside sequenceFunc ss:", ss, "and len(ss):", len(ss))
	for i, s := range ss {
		fmt.Println("i:", i, "s:", s)
	}
	if len(ss) == 0 {
		return nil, errors.New("sequence must have at least one element")
	}
	return &generator{ss, 0, sequenceGen}, nil
}

func cycleFunc(ss ...string) (*generator, error) {
	fmt.Println("##Inside cycleFunc ss:", ss, "and len(ss):", len(ss))
	for i, s := range ss {
		fmt.Println("i:", i, "s:", s)
	}
	if len(ss) == 0 {
		return nil, errors.New("cycle must have at least one element")
	}
	return &generator{ss, 0, cycleGen}, nil
}
