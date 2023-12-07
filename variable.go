package main

import "fmt"

func main() {
	var name string   // initial awal variable
	name = "Daril"    // mengisi value
	fmt.Println(name) //mencetak variabel name
	name = "Agil"     // mengubah value dari var name
	fmt.Println(name) // mencetak variable setelah diubah.
	var (
		firstName  = "Daril"
		middleName = "insan"
		lastName   = "kamil"
	)
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
