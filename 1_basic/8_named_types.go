package main

import "fmt"

type km float64
type miles float64

func main() {
	type speed uint

	var s1 speed = 10
	var s2 speed = 10

	fmt.Println(s2 - s1)

	var x uint
	//x = s1 Error

	x = uint(s1)
	_ = x

	var s3 speed = speed(x)
	_ = s3

	var parisToLondon km = 465
	var distanceInMiles miles

	distanceInMiles = miles(parisToLondon) / 0.621
	fmt.Println(distanceInMiles)
}
