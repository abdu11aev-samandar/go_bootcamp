package main

import (
	"fmt"
	"log"
	"os"
)

func foo() {
	fmt.Println("This is foo()")
}

func bar() {
	fmt.Println("This is bar()")
}

func foobar() {
	fmt.Println("This is foobar()")
}

func main() {
	defer foo()
	bar()

	fmt.Println("Just a string after defering foo() and bar() and calling bar()")

	defer foobar()

	file, err := os.Open("4_defer_stmt.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
