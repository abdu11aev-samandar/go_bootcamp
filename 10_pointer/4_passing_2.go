package main

import "fmt"

func changeValue(quality int, price float64, name string, sold bool) {
	quality = 3
	price = 500.4
	name = "Phone"
	sold = false
}

func changeValueByPointer(quality *int, price *float64, name *string, sold *bool) {
	*quality = 3
	*price = 500.4
	*name = "Phone"
	*sold = false
}

type Product struct {
	price       float64
	productName string
}

func changeProduct(p *Product) {
	(*p).price = 300
	p.productName = "Bicycle"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}

func main() {
	quality, price, name, sold := 5, 300.4, "Laptop", true
	fmt.Println("before calling changeValues():", quality, price, name, sold)
	changeValue(quality, price, name, sold)
	fmt.Println("after calling changeValues():", quality, price, name, sold)

	changeValueByPointer(&quality, &price, &name, &sold)
	fmt.Println("after calling changeValues():", quality, price, name, sold)

	gift := Product{
		price:       100,
		productName: "Watch",
	}

	changeProduct(&gift)

	fmt.Println(gift)

	fmt.Println("before calling changeProductByPointer():", gift)
	//changeValueByPointer(&gift)
	fmt.Println("after calling changeProductByPointer():", gift)

	prices := []int{1, 2, 3}
	changeSlice(prices)
	fmt.Println("prices slice after calling changeSlice():", prices)

	myMap := map[string]int{"a": 1, "b": 2}
	changeMap(myMap)
	fmt.Println("myMap after calling changeMap():", myMap)
}
