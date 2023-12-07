package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	//augmented assignment
	c += 10
	fmt.Println(c)
	c -= 10
	fmt.Println(c)
	//unary operator ++ -- !
	c++
	fmt.Println(c)
	c--
	fmt.Println(c)
}
