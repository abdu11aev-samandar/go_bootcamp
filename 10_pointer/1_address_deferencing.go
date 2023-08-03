package main

import (
	"fmt"
)

func main() {
	name := "Anna"
	fmt.Println(&name)

	var x int = 2
	ptr := &x

	fmt.Printf("ptr is of type %T with a value of %v and address %p\n", ptr, ptr, &ptr)
	fmt.Printf("address of x is %p\n", &x)

	var ptr1 *float64
	_ = ptr1

	p := new(int)

	x = 100
	p = &x

	fmt.Printf("p is of type %T with a value of %v\n", p, p)
	fmt.Printf("address of x is %p\n", &x)

	*p = 90
	fmt.Println(x, *p)
	fmt.Println(*p == x)

	*p = 10
	*p = *p / 2
	fmt.Println(x)
}
