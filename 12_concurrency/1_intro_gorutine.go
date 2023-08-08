package main

import (
	"fmt"
	"time"
)

func doSomething() {
	fmt.Println("This function says: \"Hello\"")
}

func main() {
	doSomething()

	go doSomething()

	time.Sleep(time.Second * 2)
}
