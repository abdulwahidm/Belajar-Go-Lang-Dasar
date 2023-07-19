package main

import "fmt"

func main() {

	// Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
	// Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti

	type NoKTP string

	var ktpSaya = "134782478470"
	fmt.Println(ktpSaya)
	fmt.Println(NoKTP("123467894345"))
	
}