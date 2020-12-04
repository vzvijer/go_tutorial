package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Print(a, b, " -> ")
	t := a
	a = b
	b = t
	fmt.Println(a, b)

	a, b = 3, 4
	fmt.Print(a, b, " -> ")
	a, b = b, a
	fmt.Println(a, b)
}
