package main

import "fmt"

func main() {
	const a1 = 0
	const b1 = 1.1

	fmt.Println(a1, b1)
	fmt.Printf("%T %T\n", a1, b1)

	const (
		a2 = 3 + iota
		b2
		c2
	)
	fmt.Println(a2, b2, c2)

	const (
		a3 = iota
		b3
		c3
	)
	fmt.Println(a3, b3, c3)
}
