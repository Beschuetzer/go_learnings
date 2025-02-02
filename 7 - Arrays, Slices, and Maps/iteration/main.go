package main

import (
	"fmt"
)

const (
	LinkedIn = "linkedIn"
	Aws = "aws"
	Google = "google"
)

type WebsitesMapKey string
type WebsitesMap map[WebsitesMapKey]string

func main () {
	websites := WebsitesMap{
		LinkedIn: "https://google.com",
		Aws: "https://aws.com",
	}

	userNames := make([]string, 2, 10)
	userNames = append(userNames, "Adam")
	userNames = append(userNames, "Test")

	for range userNames {
		fmt.Println("The most basic type of iteration when you want to do something the same number of times as a map or array but don't need to use the key or value")
	}

	//this is just iteration over the keys
	for index := range userNames {
		fmt.Printf("userName's index is '%v'\n", index)
	}

	//this is iteration when you want key and value
	for index, value:= range userNames {
		fmt.Printf("userName's index '%d' is '%v'\n", index, value)
	}

	//the same syntax for iteratin maps
	for key, value := range websites {
		fmt.Printf("website's key '%v' is '%v'\n", key, value)
	}
}
