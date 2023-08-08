package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now().UnixNano() / 1000000
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Hello!"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Salom!"
	}()

	time.Sleep(time.Second * 2)
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received:", msg2)
		default:
			fmt.Println("no activity")
		}
	}
	end := time.Now().UnixNano() / 1000000
	fmt.Println(end - start)
}
