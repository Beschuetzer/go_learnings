package main

import (
	"fmt"
)

const (
	Double = 2
	Triple = 3
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

	fmt.Printf("Doubled ints using transformNumbers: %v\n", doubled)
	fmt.Printf("Triple float64 using transformNumbers: %v\n", tripled)

	doubled = transformNumbersWithNumber(&numbers, 2)
	tripled = transformNumbersWithNumber(&numbersFloat, 3)

	fmt.Printf("Doubled ints using transformNumbersWithNumber: %v\n", doubled)
	fmt.Printf("Triple float64 using transformNumbersWithNumber: %v\n", tripled)
}

func transformNumbers[T Numeric](numbers *[]T, transformer transformerFn[T]) []T {
	newNumbers := []T{}
	for _, value := range *numbers {
		newNumbers = append(newNumbers, transformer(value))
	}
	return newNumbers
}

func transformNumbersWithNumber[T Numeric](numbers *[]T, transformerType int) []T {
	newNumbers := []T{}
	for _, value := range *numbers {
		newNumbers = append(newNumbers, getTransformerFn[T](transformerType)(value))
	}
	return newNumbers
}

func getTransformerFn[T Numeric](transformerType ...int) transformerFn[T] {
	if len(transformerType) == 0 {
		return double
	}
	switch transformerType[0] {
		case Double:
			return double
		case Triple:
			return triple
		default:
			return double
	}
}

func double[T Numeric](number T) T {
	return number * 2
}

func triple[T Numeric] (number T) T {
	return number * 3
}