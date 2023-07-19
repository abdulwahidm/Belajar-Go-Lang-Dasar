package main

import "fmt"

// Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
// Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah const, bukan var
// Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya


func main() {
	const firstName string = "Abdul"
	const lastName = "Wahid"

	fmt.Println(firstName)
	fmt.Println(lastName)

	firstName = "Muhamin" // error (cannot assign to firstName (constant))
	fmt.Println(firstName) 
}