package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const maxTurns = 10

	// arg1 := os.Args[1]

	rand.Seed(time.Now().Unix())
	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
