package main

import "fmt"

func main() {
	type book struct {
		title  string
		author string
		year   int
	}

	lastBook := book{title: "Anna"}
	fmt.Println(lastBook.title)
	//page := lastBook.pages
	fmt.Printf("%#v\n", lastBook)

	lastBook.author = "Leo"
	lastBook.year = 1878
	fmt.Printf("%+v\n", lastBook)

	aBook := book{title: "Anna", author: "Leo", year: 1878}
	fmt.Println(aBook == lastBook)

	myBook := aBook
	myBook.year = 2020
	fmt.Println(myBook, aBook)
}
