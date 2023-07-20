package main

import "fmt"

func main() {
	var i1 int8 = -128
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535
	fmt.Printf("%T\n", i2)

	var f1, f2, f3 float64 = 1.1, -.2, 5.
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	var r rune = 'f'
	fmt.Printf("%T\n", r)
	fmt.Println(r)
	fmt.Printf("%x\n", r)

	var b bool = true
	fmt.Printf("%T\n", b)

	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)
}
