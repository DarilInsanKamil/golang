package main

import "fmt"

func main() {
	name := "ani"

	if name == "daril" {
		fmt.Println("Hello,", name)
	} else if name == "komang" {
		fmt.Println("hello, komang")
	} else {
		fmt.Println("hello, user")
	}

	length := len(name)
	if length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}

	//short statement dalam if
	//jika length hanya dipakai untuk pengkondisian ini saja maka gunakan shorthand
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
