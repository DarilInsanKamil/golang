package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fitltredName := filter(name)
	fmt.Println("Hello ", fitltredName)
}

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	fitltredName := filter(name)
// 	fmt.Println("Hello ", fitltredName)
// }

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}
func main() {
	sayHelloWithFilter("daril", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("anjing", filter)

}
