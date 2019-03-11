package main

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a *author) fullName() string {
	return fmt.Sprintf("%s, %s", a.firstName, a.lastName)
}

func (a *post) fullName() string {
	return fmt.Sprintf("%s, %s", a.title, a.content)
}

type post struct {
	title   string
	content string
	*author
}

//Note: One of the field author is anonymous(i.e.Field with NO variable name) structure.
func (p *post) details() {
	fmt.Println("Title: ", p.title)                       // Title: Go doesn’t support Inheritance
	fmt.Println("Content: ", p.content)                   // Content: Go supports composition instead of inheritance
	fmt.Println("post fullName: ", p.fullName())          // post fullName:  Go doesn’t support Inheritance, Go supports composition instead of inheritance
	fmt.Println("Author fullName: ", p.author.fullName()) // Author fullName:  Naveen, Ramanathan
	fmt.Println("Bio: ", p.author.bio)                    // Bio:  Golang Enthusiast
}
func main() {
	author1 := &author{"Naveen", "Ramanathan", "Golang Enthusiast"}
	post1 := &post{
		"Go doesn’t support Inheritance",
		"Go supports composition instead of inheritance",
		author1,
	}
	post1.details()
}

/*
Title:  Go doesn’t support Inheritance
Content:  Go supports composition instead of inheritance
post fullName:  Go doesn’t support Inheritance, Go supports composition instead of inheritance
Author fullName:  Naveen, Ramanathan
Bio:  Golang Enthusiast
*/
