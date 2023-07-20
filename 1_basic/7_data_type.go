package main

import "fmt"

func main() {
	// int
	var i1 int8 = -128
	fmt.Printf("%T\n", i1)

	// uint
	var i2 uint16 = 65535
	fmt.Printf("%T\n", i2)

	// Float
	var f1, f2, f3 float64 = 1.1, -.2, 5.
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	// Rune
	var r rune = 'f'
	fmt.Printf("%T\n", r)
	fmt.Println(r)
	fmt.Printf("%x\n", r)

	//String
	var b bool = true
	fmt.Printf("%T\n", b)

	// String
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)

	// Array
	var numbers = [4]int{4, 5, 6, 7}
	fmt.Printf("%T\n", numbers)

	// Slice
	var cities = []string{"London", "Tokyo"}
	fmt.Printf("%T\n", cities)

	// Map
	balances := map[string]float64{
		"USD": 2332.2,
		"EUR": 511.11,
	}
	fmt.Printf("%T\n", balances)

	// Struct
	type Person struct {
		name string
		age  int
	}
	var you Person
	fmt.Printf("%T\n", you)

	// Pointer
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr)

	// Function
	fmt.Printf("%T\n", f)
}

func f() {

}
