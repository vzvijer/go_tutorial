package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func max(a []int) (maximum int, index int) {
	if len(a) == 0 {
		return 0, -1
	}

	index = 0
	maximum = a[index]
	for i, v := range a {
		if v > maximum {
			index = i
			maximum = v
		}
	}
	return maximum, index
}

func main() {
	f, err := ioutil.ReadFile("numbers.txt")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(f))
	numbers := strings.Split(string(f), " ")

	nums := []int{}
	for _, number := range numbers {
		x, _ := strconv.Atoi(number)
		nums = append(nums, x)
	}
	fmt.Println(nums)
	max, ind := max(nums)
	fmt.Println("The biggest number is", max, "Its index is: ", ind)
}
