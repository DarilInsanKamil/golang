package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecrusive(value int) int {
	if value == 1 { //jika valuenya satu maka return 1
		return 1
	} else { // jika valuenya bukan satu maka value * factorialRecrusive dan akan terus memanggil fungsi tsb sampai 1
		return value * factorialRecrusive(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecrusive(10))
}
