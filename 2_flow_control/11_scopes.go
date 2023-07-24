package main

import "fmt"
import f "fmt"

const done = false // package scoped

var b int = 10

func main() {
	var task = "Running" //local (block) scoped
	fmt.Println(task, done)

	const done = true // local scoped
	f.Println("done in main() is %v\n", done)
	f1()

	f.Println("Bye Bye!")
}

func f1() {
	const done = true
	f.Printf("done is f1(): %v\n", done) //this is done from package scope
	a := 10
	_ = a

}
