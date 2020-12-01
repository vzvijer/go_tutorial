package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Please enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Print("Hello ", name)
	}
}
