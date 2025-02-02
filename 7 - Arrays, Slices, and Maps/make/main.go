package main

import (
	"fmt"
)

// the 'make' function tells Go "I'm going to be adding X number of elements to the array soon" so start with a bigger array 
// so as not to have to remake the array when I do
func main() {
	//this will crash the app since the above creates an empty array
	// userNames := []string{}
	// userNames[0] = "Adam"


	// this means create an array with 2 empty slots but with capacity of 5
	// this makes appending items more efficient
	// this only works if you know in advance roughly how many items to expect
	// this will pre-allocated memory for 5 items and create first 2 slots
	userNames := make([]string, 2, 5)
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Adam")

	fmt.Println(userNames)

	//making maps can impore performance if you know roughly how many items to expect
	//this will pre-allocate memory to accomodate for 100 items
	courseRatings := make(map[string]float64, 100)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.2

	fmt.Println(courseRatings)
}