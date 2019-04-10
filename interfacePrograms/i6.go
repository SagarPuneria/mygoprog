package main

import (
	"fmt"
	"math"
)

/* define an interface */
type Shape interface {
	area() float64
}

/* define a circle */
type Circle struct {
	x, y, radius float64
}

/* define a rectangle */
type Rectangle struct {
	width, height float64
}

/* define a method for circle (implementation of Shape.area())*/
func (circle *Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

/* define a method for rectangle (implementation of Shape.area())*/
func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}

/* define a method for shape */
func getArea(shape Shape) float64 {
	return shape.area()
}
func main() {
	circle := &Circle{x: 0, y: 0, radius: 5}
	rectangle := &Rectangle{width: 10, height: 5}
	fmt.Println("Circle area:", getArea(circle))       //Circle area: 78.53981633974483
	fmt.Println("Rectangle area:", getArea(rectangle)) //Rectangle area: 50
	fmt.Println("-------(OR)------------")
	Shapes := []Shape{&Circle{x: 0, y: 0, radius: 5}, &Rectangle{width: 10, height: 5}}
	//Shapes := []Shape{circle, rectangle}
	for _, v := range Shapes {
		fmt.Println(v.area())
	}
}

/*
Circle area: 78.53981633974483
Rectangle area: 50
-------(OR)------------
78.53981633974483
50
*/
