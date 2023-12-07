package main

import "fmt"

func main() {
	const (
		firstName string = "Daril"
		lastName         = "Kamil"
	) // buat multiple variable pakai const
	fmt.Println(firstName)
	fmt.Println(lastName)
	//error
	// firstName = "Komang"
	// lastName = "Ucul"
	// 	.\constant.go:6:2: cannot assign to firstName (neither addressable nor a map index expression)
	// .\constant.go:7:2: cannot assign to lastName (neither addressable nor a map index expression)

}
