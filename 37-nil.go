package main

import "fmt"

// Biasanya di dalam bahasa pemrograman lain, object yang belum diinisialisasi maka secara otomatis nilainya adalah null atau nil
// Berbeda dengan Go-Lang, di Go-Lang saat kita buat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan default value nya
// Namun di Go-Lang ada data nil, yaitu data kosong
// Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer dan channel

func main() {

	var abdul map[string] string = nil
	fmt.Println(abdul) // map []
}