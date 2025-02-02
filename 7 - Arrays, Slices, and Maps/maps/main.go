package main

import (
	"fmt"
)

// this is functioning like an enum in TS
const (
	LinkedIn = "linkedIn"
	Aws = "aws"
	Google = "google"
)

type WebsitesMapKey string
type WebsitesMap map[WebsitesMapKey]string

func main() {
	//Go allows any type as a key (e.g. a struct)
	websites := WebsitesMap{
		LinkedIn: "https://google.com",
		Aws: "https://aws.com",
	}
	fmt.Printf("Map is: { %v }\n", websites)
	websites.printNicely(Aws)
	websites[LinkedIn] = "https://linkedin.com"
	websites.printNicely(LinkedIn)
	delete(websites, LinkedIn)
	websites.printNicely(LinkedIn)
}

func (websitesMap WebsitesMap) printNicely(key WebsitesMapKey) {
	//accessing keys is the same as JS
	fmt.Printf("%v: %v\n", key, websitesMap[key])
}

