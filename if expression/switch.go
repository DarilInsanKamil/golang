package main

import "fmt"

func main() {
	name := "darilinsankamil"

	switch name {
	case "komang":
		fmt.Println("hello, komang")
	case "daril":
		fmt.Println("Hello, Daril")
	default:
		fmt.Println("Hello User")
	}

	//shorthand switch
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	length := len(name)
	//switch tanpa kondisi
	//jika tidak menggunakan ekspresi maka pengkondisian ada di case

	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
