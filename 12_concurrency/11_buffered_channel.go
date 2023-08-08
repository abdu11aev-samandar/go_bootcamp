package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Println(i)
			c <- i
			fmt.Println(i)
		}
		close(c)
	}(c)

	time.Sleep(time.Second * 2)
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println(<-c)

	//c <- 10
}
