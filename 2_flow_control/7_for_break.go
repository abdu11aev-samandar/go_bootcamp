package main

import "fmt"

func main() {
	count := 0

	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13 \n", i)
			count++
		}
		if count == 10 {
			break
		}
	}
	fmt.Println("Just a message after the for loop")
}
