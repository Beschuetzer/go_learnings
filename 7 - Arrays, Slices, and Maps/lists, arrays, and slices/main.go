package main

import (
	"fmt"
)

type Product struct {
	Name string
	Aisle int
	Quantity int
}

func main() {
	products := []Product{
		{
			Name: "My Product",
			Aisle: 1,
			Quantity: 1,
		},
		{
			Name: "My Product 2",
			Aisle: 2,
			Quantity: 2,
		},
	}
	fmt.Println(products)

	//This is how to add a new element to a slice/array when using literal
	newProductsUsingLiterals := append(products, Product{
		Name: "My Product 3",
		Aisle: 3,
		Quantity: 3,
	}, Product{
		Name: "My Product 4",
		Aisle: 4,
		Quantity: 4,
	})

	fmt.Println(newProductsUsingLiterals)

	
	//This is how to add a new element to a slice/array when you have another slice/array
	newProductsWithAnotherSliceOrArray := append(products, newProductsUsingLiterals...)
	fmt.Println(newProductsWithAnotherSliceOrArray)
}
