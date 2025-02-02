package main

import (
	"fmt"
)

//this is creating a new type that would be the same as int | float32 | float64 in TS if those types existed
type Numeric interface {
	~float32 |
	~float64 |
	~int
}

type transformerFn[T Numeric] func (T) T

func main() {
	numbers := []int{1, 2, 3, 4}
	numbersFloat := []float64{1.1, 2.2, 3.3, 4.4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbersFloat, triple)

	fmt.Printf("Doubled ints: %v\n", doubled)
	fmt.Printf("Triple float64: %v\n", tripled)
}

func transformNumbers[T Numeric](numbers *[]T, transformer transformerFn[T]) []T {
	doubleNumbers := []T{}
	for _, value := range *numbers {
		doubleNumbers = append(doubleNumbers, transformer(value))
	}
	return doubleNumbers
}

func double[T Numeric](number T) T {
	return number * 2
}

func triple[T Numeric] (number T) T {
	return number * 3
}