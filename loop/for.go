package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 { // jika kondisi belum terpenuhi maka counter akan ditambah 1
	// 	fmt.Println("Perulangan ke, ", counter) // akan ke eksekusi jika counter <= 10
	// 	counter++                               //menambahkan counter
	// }

	//counter dengan statement
	for counter2 := 1; counter2 <= 10; counter2++ {
		fmt.Println("Perulangan ke, ", counter2) // akan ke eksekusi jika counter <= 10
	}

	//for loop digunakan untuk memecah array
	names := []string{"daril", "kamil", "isnan"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	fmt.Println("=================")
	//for range digunakan unutk iterasi pada data collection, yaitu array, slice, map
	for _, name := range names {
		fmt.Println(name)
	}
	//index kalo di tipe data map yaitu key, dan name itu value
	fmt.Println("Selesai")
}
