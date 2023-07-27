package main

import "fmt"

func main() {
	var cities []string
	fmt.Println("cities is equal to nil: ", cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(cities)

	//cities[0]="London"
	nums := []int{2, 3, 4, 5}
	fmt.Println(nums)

	num := make([]int, 2)
	fmt.Printf("%#v\n", num)

	type names []string
	friend := names{"Dan", "Maria"}
	fmt.Println(friend)

	myFriend := friend[0]
	fmt.Println("My best friend is", myFriend)

	friend[0] = "Gabriel"
	fmt.Println("My best friend is ", friend[0])

	for i, v := range nums {
		fmt.Printf("index: %v, value: %v\n", i, v)
	}

	var n []int
	n = nums
	fmt.Println(n)
}
