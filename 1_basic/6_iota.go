package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)

	const (
		c11 = iota
		c22
		c33
	)
	fmt.Println(c1, c2, c3)

	const (
		North = iota
		East
		South
		West
	)
	fmt.Println(West)

	const (
		a = (iota * 2) + 1
		b
		c
	)
	fmt.Println(a, b, c)

	const (
		x = -(iota + 2)
		_
		y
		z
	)

}
