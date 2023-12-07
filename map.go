package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "daril"
	// person["age"] = "20"

	person := map[string]string{
		"name": "Daril",
		"age":  "20",
	}

	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person)

	//membuat map
	book := make(map[string]string)

	book["title"] = "ubur-ubur lembur"
	book["author"] = "raditya dika"
	book["wrong"] = "Ups"
	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)
}
