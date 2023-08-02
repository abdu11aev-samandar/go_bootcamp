package main

import "fmt"

func f1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", a)
}

func main() {
	f1(1, 2, 3, 4)
	f1()

	nums := []int{1, 2}
	nums = append(nums, 3, 4)

	f1(nums...)

}
