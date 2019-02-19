package main

import "fmt"

type Mover interface {
	Move()
}

type Car struct {
	price int
	Mover
	speed int
}

type Wheel struct {
}

func (w *Wheel) Move() {
	fmt.Println("Inside wheel")
}

// promoted method
func (c *Car) Move() {
	fmt.Println("Inside Move")
	c.Mover.Move() // Crashes, nil pointer dereference. If Field 'Mover' is nil
}

func main() {
	var m Mover = &Car{Mover: &Wheel{}}
	//var m Mover = &Car{} // Field 'Mover' is nil.
	m.Move()
}
