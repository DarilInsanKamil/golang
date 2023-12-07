package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are Blocked")
	} else {
		fmt.Println("welcome")
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Daril", blacklist)

	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
