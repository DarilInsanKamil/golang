package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Daril"
	names[1] = "Insan"
	names[2] = "Kamil"

	fmt.Println(names)    //mencetak seluruh array
	fmt.Println(names[0]) //mencetak data berdasarkan indexnya

	names[0] = "Komang" //mengubah dengan cara mengakses indexnya

	fmt.Println(names[0]) //mencetak data berdasarkan indexnya
	fmt.Println(names[1]) //mencetak data berdasarkan indexnya
	fmt.Println(names[2]) //mencetak data berdasarkan indexnya

	//membuat array secara langsung
	var values = [...]int{90, 100, 56, 909} // kan ini menetapkan 3 data jika kita hanya mengisi 2 data maka akan mencetak nilai defaultnya yaitu 0
	fmt.Println(values)

	fmt.Println(len(values))

}
