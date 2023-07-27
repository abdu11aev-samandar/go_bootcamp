package main

import (
	"fmt"
	"unsafe"
)

func main() {
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println(cars, newCars)

	s1 := []int{10, 20, 30, 40, 50}
	newSlice := s1[0:3]
	fmt.Println(len(newSlice), cap(newSlice))

	newSlice = s1[2:5]
	fmt.Println(len(newSlice), cap(newSlice))

	fmt.Printf("%p\n", &s1)

	fmt.Printf("%p %p \n", &s1, &newSlice)

	newSlice[0] = 1000
	fmt.Println("s1: ", s1)

	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("arrays's size in butes: %d \n", unsafe.Sizeof(a))
	fmt.Printf("slice's size in butes: %d \n", unsafe.Sizeof(s))

}
