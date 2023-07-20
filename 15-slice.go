package main

import "fmt"

func main() {
	// Tipe data Slice adalah potongan dari data Array
	// Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
	// Slide dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di Array

	// Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity
	// Pointer adalah penunjuk data pertama di array para slice
	// Length adalah panjang dari slice, dan
	// Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

	names := []string {
		"abdul", 
		"wahid",
		"muhaemin",
	}
	fmt.Println(names)
	names = append(names, "budi") // menambahkan elemen
	names = append(names, "john")
	fmt.Println(names) // [abdul wahid muhaemin budi]
	slice := names[1:3] // memotong elemen ke 1 sampai sebelum 3
	slice1 := names[1:] // memotong elemen ke 1 sampai akhir
	fmt.Println(slice) // [wahid muhaemin]
	fmt.Println(slice1) // [wahid muhaemin budi john]

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Minggu"
	newSlice[1] = "Senin"
	fmt.Println(newSlice) // [Minggu Senin]

	// Saat membuat Array, kita harus berhati-hati, jika salah, maka yang kita buat bukanlah Array, melainkan Slice
	iniArray := [...]int {1, 2, 3, 4, 5}
	iniSlice := [] int {6, 7, 8, 9}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}

