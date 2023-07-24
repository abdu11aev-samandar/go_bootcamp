package main

import (
	"fmt"
	"time"
)

func main() {
	//////////////////////////////
	lang := "golang"

	switch lang {
	case "Python":
		fmt.Println("You are learning Python!")
	case "Go", "golang":
		fmt.Println("Good, go for Go!")
	default:
		fmt.Println("Any other programming lang is a good start!")
	}

	/////////////////////////
	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even num!")
	case n%2 != 0:
		fmt.Println("Odd num!")
	default:
		fmt.Println("Never here!")
	}

	///////////////////////////////
	hour := time.Now().Hour()
	//fmt.Println(hour)
	switch {
	case hour < 12:
		fmt.Println("Good Morning!")
	case hour < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}
