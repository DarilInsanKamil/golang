package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80 // cek nilai akhir. output = true
	var lulusNilaiAbsensi bool = absensi > 80  // cek nilai absensi. output = false

	var lulus bool = lulusNilaiAkhir && lulusNilaiAbsensi //cek kelulusan. output = false
	//karena salah satu false maka hasilnya false.

	fmt.Println(lulus)
}
