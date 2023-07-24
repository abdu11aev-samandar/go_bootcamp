package main

import "fmt"

func main() {
	i := 0

loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	//	THIS IS AN ERROR
	//	goto todo
	//	x := 5
	//todo:
	//	fmt.Println("something here")
}
