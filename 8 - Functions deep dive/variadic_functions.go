package main

import "fmt"

func main() {
	integers := []int{1,2,3,4,5}

	fmt.Println(sumUp(1, 2, 3, 4))
	fmt.Println(sumUp(integers...))
}

func sumUp(numbers ...int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}