package main

import "fmt"

func change(a *int) *float64 {
	*a = 100

	b := 5.5
	return &b
}

func changeVar(a int) {
	a = 66
}

func main() {
	x := 8
	p := &x

	fmt.Println("value of x before calling change():", x)
	change(p)
	fmt.Println("value of x after calling change():", x)

	fmt.Println("value of x after calling changeVar():", x)
	changeVar(x)
	fmt.Println("value of x after calling changeVar():", x)

}
