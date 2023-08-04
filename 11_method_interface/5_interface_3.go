package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

/*func printCircle(c circle) {
	fmt.Println("Shape", c)
	fmt.Println("Area", c.area())
	fmt.Println("Perimeter", c.perimeter())
}

func printRectangle(r rectangle) {
	fmt.Println("Shape", r)
	fmt.Println("Area", r.area())
	fmt.Println("Perimeter", r.perimeter())
}*/

func printShape(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %#v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {
	var s shape
	fmt.Printf("%T\n", s)

	ball := circle{radius: 2.5}
	s = ball

	printShape(s)
	fmt.Printf("Type of s: %T\n", s)

	room := rectangle{width: 2, height: 3}
	s = room
	fmt.Printf("Type of s: %T\n", s)
}
