package main

import "fmt"

func main() {
	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"Dan", "Diana", "Paul", "John"}
	fmt.Printf("%#v\n", a3)

	a4 := [4]string{"x", "y"}
	fmt.Printf("%#v\n", a4)

	a5 := [...]int{1, 2, 5, 1, -10, 66}
	fmt.Printf("%#v\n", a5)
	fmt.Println("Lengs", len(a5))

	a6 := [...]int{
		1,
		2,
		3,
		4, // this comma is madatory
	}
	fmt.Println(a6)
}
