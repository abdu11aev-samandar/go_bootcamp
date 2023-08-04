package main

import "fmt"

type emptyInterface interface {
}

type person struct {
	info interface{}
}

func main() {
	var empty interface{}

	empty = 5
	fmt.Println(empty)

	empty = "Go"
	fmt.Println(empty)

	empty = []int{4, 5, 6}
	fmt.Println(empty)
	fmt.Println(len(empty.([]int)))

	you := person{}
	you.info = "Your name"
	you.info = 40
	you.info = []float64{5.6, 5.}
	fmt.Println(you.info)
}
