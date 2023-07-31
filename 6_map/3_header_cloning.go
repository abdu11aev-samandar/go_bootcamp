package main

import "fmt"

func main() {
	friends := map[string]int{"Dan": 40, "Maria": 25}

	neighbors := friends

	friends["Dan"] = 50
	fmt.Println(neighbors)

	people := make(map[string]int)

	for k, v := range friends {
		people[k] = v
	}

	people["Anne"] = 18
	fmt.Printf("%#v ------------ %#v\n", people, friends)

	delete(friends, "Dan")
	fmt.Println(friends)
}
