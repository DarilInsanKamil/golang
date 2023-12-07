package main

import "fmt"

func main() {
	var nama = "daril"
	var nama2 = "darils"

	var res bool = nama == nama2
	fmt.Println(res) //true
	var res2 bool = nama != nama2
	fmt.Println(res2) //false

	var angka1 = 20
	var angka2 = 10

	var resAngka = angka1 < angka2
	fmt.Println(resAngka) //false
	var resAngka2 = angka1 > angka2
	fmt.Println(resAngka2) //true
}
