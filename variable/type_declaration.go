package main

import "fmt"

func main() {
	type NOKTP string // inisial alias atau type

	var ktpDaril NOKTP = "11111"
	//merubah tipe data contoh menjadi NOKTP
	var contoh string = "222222"
	var contohKtp NOKTP = NOKTP(contoh)

	fmt.Println(ktpDaril)
	fmt.Println(contohKtp)
}
