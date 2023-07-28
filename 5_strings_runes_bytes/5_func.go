package main

import (
	"fmt"
	"strings"
)

func main() {

	p := fmt.Println

	result := strings.Contains("I love go Programming!", "love")
	p(result)

	result = strings.ContainsAny("success", "xy")
	p(result)

	result = strings.ContainsRune("golang", 'g')
	p(result)

	n := strings.Count("cheese", "e")
	p(n)

	n = strings.Count("Five", "")
	p(n)

	p(strings.ToLower("GO PYTHON jaVA"))
	p(strings.ToUpper("GO PYTHON jaVA"))

	p("go" == "go")
	p("GO" == "go")

	p(strings.ToLower("GO") == strings.ToLower("go"))

	p(strings.EqualFold("GO", "go"))

}
