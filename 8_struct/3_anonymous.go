package main

import "fmt"

func main() {
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}

	fmt.Printf("%#v\n", diana)
	fmt.Printf("Diana's Age: %d\n", diana.age)
	///////////////////////////////////////////////////////////////
	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{"1984 by Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1)

	fmt.Println(b1.string)
	/////////////////////////////////////////////
	type Employee struct {
		name   string
		salary int
		bool
	}

	e := Employee{"John", 4000, false}
	fmt.Printf("%#v\n", e)
	e.bool = true
	fmt.Printf("%#v\n", e)
}
