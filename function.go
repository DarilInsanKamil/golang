package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

// function dengan parameter
func penjumlahan(a int, b int) {
	var c = a + b
	fmt.Println(c)
}

// function with return (mengembalikan value atau data)
func getHello(name string) string {
	hello := "hello " + name
	return hello
}

// function with multiple value (mengembalikan value lebih dari satu)
func getFullName() (string, string) {
	return "daril", "insan"
}

func main() {
	a := 5
	b := 5
	result := getHello("Daril")
	fmt.Println(result)
	fmt.Println(getHello("insan"))
	fmt.Println(getHello("kamil"))
	sayHello()
	penjumlahan(a, b)
	penjumlahan(7, 4)

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
