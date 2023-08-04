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

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
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
	var s shape = circle{radius: 2.5}
	fmt.Printf("%T\n", s)

	//s.volume()
	ball, ok := s.(circle)

	if ok == true {
		fmt.Printf("Ball Volume: %v\n", ball.volume())
	}

	s = rectangle{width: 3.4, height: 2.2}
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)
	}
}
