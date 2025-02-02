package main

import "fmt"

func main() {
	factorialsToGet := []int{-1,2,3,4,5}

	for _, value := range factorialsToGet {
		fmt.Printf("%v factorial is %d\n", value, getFactorial(value))
	}

}

func getFactorial(number int) int {
	if number < 0 {
		return -1
	}
	if number == 0 {
		return 1
	}
	return number * getFactorial(number-1)
}