package main

import (
	"fmt"
	"errors"
)

// Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interface nya adalah error.
// Untuk membuat error, kita tidak perlu manual.
// Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors 
// (Package akan kita bahas secara detail di materi tersendiri)

func main() {
	hasil, _ := Pembagi(100, 25)
	fmt.Println("Hasil:", hasil)
}

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	} else {
		result := nilai/pembagi
		return result, nil
	}
}