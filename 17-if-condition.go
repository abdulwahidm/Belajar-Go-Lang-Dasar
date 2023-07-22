package main

import "fmt"

func main() {
	
	var point int = 3

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	// Variabel Trmporary pada if else condition 
	// adalah variabel yang hanya bisa digunakan pada blok seleksi
	// kondisi di mana ia ditempatkan saja. Penggunaan variabel ini membawa
	// beberapa manfaat, antara lain:
	// Scope atau cakupan variabel jelas, hanya bisa digunakan pada blok seleksi kondisi itu saja
	// Kode menjadi lebih rapi
	// Ketika nilai variabel tersebut didapat dari sebuah komputasi, perhitungan
	// tidak perlu dilakukan di dalam blok masing-masing kondisi.
	
	value := 8840.0

	if percent := value / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}