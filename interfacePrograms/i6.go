package main

import (
	"fmt"
	"math"
)

// Shape ...define an interface
type Shape interface {
	area() float64
}

// Circle ...define a circle
type Circle struct {
	x, y, radius float64
}

// Rectangle ...define a rectangle
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
	fmt.Println("Circle area:", getArea(circle), circle.area()) // Circle area: 78.53981633974483 78.53981633974483

	shape := Shape(circle)
	fmt.Println("Circle area:", shape.area()) // Circle area: 78.53981633974483

	circle2, ok := shape.(*Circle)
	if ok {
		fmt.Println("Circle area:", circle2.area()) // Circle area: 78.53981633974483
	}

	rectangle := &Rectangle{width: 10, height: 5}
	fmt.Println("Rectangle area:", getArea(rectangle), rectangle.area()) // Rectangle area: 50 50

	shape = Shape(rectangle)
	fmt.Println("Rectangle area:", shape.area()) // Rectangle area: 50

	rectangle2, ok := shape.(*Rectangle)
	if ok {
		fmt.Println("Rectangle area:", rectangle2.area()) // Rectangle area: 50
	}

	fmt.Println("-------(OR)------------")
	// Shapes := []Shape{&Circle{x:0, y:0, radius:5}, &Rectangle{width:10, height:5}}
	// Shapes := []Shape{circle, rectangle}
	Shapes := []Shape{circle2, rectangle2}
	for _, v := range Shapes {
		fmt.Println(v.area()) // fmt.Println(getArea(v)) }
	}
}

/*Output:
Circle area: 78.53981633974483 78.53981633974483
Circle area: 78.53981633974483
Circle area: 78.53981633974483
Rectangle area: 50 50
Rectangle area: 50
Rectangle area: 50
-------(OR)------------
78.53981633974483
50
*/
