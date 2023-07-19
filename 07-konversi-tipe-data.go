package main 

import "fmt"

func main() {
	// Di Go-Lang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
	// Misal kita ingin mengkonversi tipe data int32 ke int63, dan lain-lain

	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)	
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Abdul Wahid Muhaemin"
	var a = name[0]
	var aString = string(a)

	fmt.Println(name) //A
	fmt.Println(aString)
}