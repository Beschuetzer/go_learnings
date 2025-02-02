package main

import "fmt"

type Number interface { 
	~int |
	~float32 |
	~float64
}
type tranformerFn[T Number] func(T) T 

func main() {
	numbers := []int{1, 2, 3}

	//using an anonymous functions
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	printNicely(2, transformed)

	//using the factory function pattern
	transformedFive := transformNumbers(&numbers, createTransformNumbersTransformer(5))
	transformedSix := transformNumbers(&numbers, createTransformNumbersTransformer(6))

	printNicely(5, transformedFive)
	printNicely(6, transformedSix)
}

func printNicely[T Number](number T, values []T) {
	fmt.Printf("Transformed %vx: %v\n", number, values)
}

func transformNumbers[T Number](numbers *[]T, transformer tranformerFn[T]) []T {
	dNumbers := []T{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transformer(val))
	}

	return dNumbers
}

func createTransformNumbersTransformer[T Number](factor T) func(T) T {
	return func(number T) T {
		return number * factor
	}
}