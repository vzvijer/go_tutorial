package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	fmt.Print("Enter two numbers: ")
	var reader = bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	checkError(err)
	numbers := strings.Split(strings.Trim(input, " \t\n"), " ")
	a, err := strconv.ParseFloat(numbers[0], 64)
	checkError(err)
	b, err := strconv.ParseFloat(numbers[1], 64)
	checkError(err)
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
	fmt.Printf("%v - %v = %v\n", a, b, a-b)
	fmt.Printf("%v * %v = %v\n", a, b, a*b)
	fmt.Printf("%v / %v = %.2v\n", a, b, a/b)
}
