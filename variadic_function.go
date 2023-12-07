package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	//menggunakan varargs
	fmt.Println(sumAll(10, 10, 10, 10))

	//menggunakan slice menjadi varargs
	numbers := []int{10, 10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}
