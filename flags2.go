package main

import (
	"flag"
	"fmt"
	"net/url"
)

type URLValue struct {
	URL *url.URL
}

func (v *URLValue) String() string {
	if v.URL != nil {
		fmt.Println("Inside String(), v.URL:", v.URL, ", v.URL.String():", v.URL.String()+";")
		return v.URL.String()
	}
	return ""
}

func (v *URLValue) Set(s string) error {
	fmt.Println("Inside Set(), s:", s)
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	fmt.Println("Inside Set(), v.URL:", v.URL, ", v.URL.String():", v.URL.String())
	return nil
}

var u = &url.URL{}

func main() {
	fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	fmt.Println("Debug 1")
	fs.Var(&URLValue{u}, "url", "URL to parse")

	fmt.Println("Debug 2, u:", u)
	fs.Parse([]string{"-url", "https://golang.org/pkg/flag/"})
	fmt.Println("Debug 3, u:", u)

	fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)

}
