package main

import (
	"fmt"
	"strings"
)

func f2(a ...int) {
	a[0] = 50
}

func sumAndProduct(a ...float64) (float64, float64) {
	sum := 0.
	product := 1.
	for _, v := range a {
		sum += v
		product *= v
	}
	return sum, product
}

func personInfo(age int, names ...string) string {
	fullName := strings.Join(names, " ")

	returnString := fmt.Sprintf("Age: %d, Full name: %s", age, fullName)
	return returnString
}

func main() {
	nums := []int{1, 2}
	nums = append(nums, 3, 4)

	f2(nums...)
	fmt.Println("Nums:", nums)

	s, p := sumAndProduct(2., 5., 10.)
	fmt.Println(s, p)

	info := personInfo(30, "Wolf", "Wolf2", "Wolf3")
	fmt.Println(info)
}
