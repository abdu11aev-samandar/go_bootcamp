package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//fmt.Println("os.Args", os.Args)
	//fmt.Println("Path:", os.Args[0])
	//fmt.Println("1st arg", os.Args[1])
	//fmt.Println("2st arg", os.Args[2])
	//fmt.Println("No if items inside os.Args:", len(os.Args))

	var results, err = strconv.ParseFloat(os.Args[1], 64)

	//fmt.Println(results)
	fmt.Printf("%T\n", os.Args[1])
	fmt.Printf("%T\n", results)
	_ = err
}
