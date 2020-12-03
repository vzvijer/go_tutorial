package main

import "fmt"

func printInts() {
	fmt.Println("Printing ints")
	fmt.Println("--------------------------")

	var a = 97
	fmt.Printf("%d\n", a)
	fmt.Printf("%v\n", a)
	fmt.Printf("%10b\n", a)
	fmt.Printf("%o\n", a)
	fmt.Printf("%q\n", a)
	fmt.Printf("%X\n", a)
}
