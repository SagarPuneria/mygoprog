package main

import "fmt"

//https://golang.org/doc/faq#closures_and_goroutines

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		/*//v := v // create a new 'v'.
		go func() {
			fmt.Println(v)
			done <- true
		}()*/
		go func(u string) {
			fmt.Println(u)
			done <- true
		}(v)
	}

	// wait for all goroutines to complete before exiting
	for _ = range values {
		<-done
	}
}
