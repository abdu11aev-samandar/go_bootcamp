package main

import (
	"fmt"
	"strings"
)

func main() {
	num := [3]int{10, 20, 30}
	fmt.Printf("%#v\n", num)

	num[0] = 7
	fmt.Printf("%#v\n", num)

	// num[5] = 100 // error

	////////////////////////////////////
	for i, v := range num {
		fmt.Println("index", i, " value:", v)
	}

	fmt.Println(strings.Repeat("#", 20))

	for i := 0; i < len(num); i++ {
		fmt.Println("index:", i, " value:", num[i])
	}

	/////////////////////////////////////////////////
	balances := [2][3]int{
		{5, 6, 7},
		[3]int{8, 9, 10},
	}
	fmt.Println(balances)

	/////////////////////////////////////////////////
	m := [3]int{1, 2, 3}
	n := m
	fmt.Println("n is equal to m:", n == m)
	m[0] = -1
	fmt.Println("n is equal to m:", n == m)
}
