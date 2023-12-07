package main

import "fmt"

// function with multiple value (mengembalikan value lebih dari satu)
func getFullName() (string, string) {
	return "daril", "insan"
}

// function with multiple value akan tetapi kita hanya butuh bebrapa datanya saja

func main() {
	// firstName, lastName := getFullName()
	_, lastName := getFullName()
	fmt.Println(lastName)
}
