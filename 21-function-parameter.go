package main

import "fmt"



func main() {
	firstName := "Abdul"
	sayHelloTo(firstName, "Wahid")
	sayHelloTo("Budi", "Nugraha")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

/*
Khusus untuk fungsi yang tipe data parameternya sama, bisa ditulis dengan gaya
yang unik. Tipe datanya dituliskan cukup sekali saja di akhir. Contohnya bisa
dilihat pada kode berikut.
*/

func nameOfFunc(paramA type, paramB type, paramC type) returnType
func nameOfFunc(paramA, paramB, paramC type) returnType

func randomWithRange(min int, max int) int
func randomWithRange(min, max int) int