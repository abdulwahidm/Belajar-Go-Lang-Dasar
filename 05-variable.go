package main

import "fmt"


// Variable adalah tempat untuk menyimpan data
// Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
// Di Go-Lang Variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa variable
// Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan nama variable dan tipe datanya

func main() {

	// 1. Dengan menggunakan keyword var maka wajib menuliskan tipe datanya
	var name string = "Abdul Wahid Muhaemin"
	fmt.Println(name)

	// 2. Menggunakan operator assignment := 
	address := "Jakarta"
	fmt.Println(address)

	// 3. Deklarasi multiivariable 
	var (
		age = 24
		gender = "male"
	)

	fmt.Println(age)
	fmt.Println(gender)	
}