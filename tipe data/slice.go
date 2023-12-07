package main

import "fmt"

func main() {
	names := [...]string{"daril", "insan", "kamil", "alesha", "athaya", "ajo"}

	slice := names[4:6]
	fmt.Println(slice) // slicing data ke 4 sampai ke 6

	slice2 := names[3:]
	fmt.Println(slice2) // slicing data ke 3 hingga data terakhirnya

	slice3 := names[:3]
	fmt.Println(slice3) // slicing data ke 3 hingga sebelum data ke 3

	slice4 := names[:]
	fmt.Println(slice4) // slicing semua data

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:]
	fmt.Println(daySlice1) // [sabtu, minggu]

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1) // [Sabtu Baru, Minggu Baru]
	fmt.Println(days)      // [senin, selasa, rabu, kamis, jumat,Sabtu Baru, Minggu Baru]

	daySlice2 := append(daySlice1, "Libur baru")
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	//membuat slice
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "daril"
	newSlice[1] = "daril"
	// newSlice[2] = "daril" // error. harus menggunkan append()

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "daril")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	//perbedaan array dan slice
	iniArray := [...]string{"hello", "world", "im", "back"}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
