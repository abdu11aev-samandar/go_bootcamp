package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var1, var2 := 'a', 'b'

	fmt.Printf("Type: %T, Value: %d\n", var1, var2)

	str := "Sara"
	fmt.Println(len(str))

	fmt.Println("Byte (not rune) at position 1", str[1])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	fmt.Println("\n" + strings.Repeat("#", 20))

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size
	}

	fmt.Println("\n" + strings.Repeat("#", 20))

	for _, r := range str {
		fmt.Printf("%c", r)
	}
}
