package main

import "fmt"

func main() {
	s1 := "I love Golang!"
	fmt.Println(s1[2:5])

	s2 := ""
	fmt.Println(s2[0:2])

	rs := []rune(s2)
	fmt.Printf("%T\n", rs)

	fmt.Println(string(rs[0:3]))
}
