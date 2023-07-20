package main

import "fmt"

func main() {

	nilai := 90
	absensi := 81

	nilaiLulus := nilai > 80
	absensiLulus := absensi > 80

	lulus := nilaiLulus && absensiLulus
	fmt.Println(lulus) //true	
}