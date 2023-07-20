package main

import "fmt"

func main() {
	const a float64 = 5.1 //typed
	const b = 6.7         //untyped

	const c float64 = a * b
	const str = "Hello" + "Go!"

	const d = 5 > 10
	fmt.Println(d)

	//const x int = 5
	//const y float64 = 2.2 * x

	const x = 5
	const y = 2.2 * 5
	fmt.Printf("%T\n", y)

	var i int = x
	var j float64 = x
	var p byte = x
	fmt.Println(i, j, p)

	const r = 5
	var rr = r
	fmt.Printf("%T", rr)
}
