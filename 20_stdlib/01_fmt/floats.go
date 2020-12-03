package main

import "fmt"

func printBools() {
	fmt.Println("Printing floats")
	fmt.Println("--------------------------")

	var x = 1.2345
	fmt.Printf("%f\n", x)
	fmt.Printf("%v\n", x)
	fmt.Println(fmt.Sprintf("%.2f", x))
	fmt.Println(fmt.Sprintf("%10.3f", x))
}
