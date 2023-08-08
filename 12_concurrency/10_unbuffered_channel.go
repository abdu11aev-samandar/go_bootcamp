package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)    //unbuffered
	c2 := make(chan int, 3) //buffered
	_ = c2

	go func(c chan int) {
		fmt.Println("func starts")
		c <- 10
		fmt.Println("func after")
	}(c1)

	time.Sleep(time.Second * 2)

	d := <-c1
	fmt.Println(d)

	time.Sleep(time.Second * 2)
}
