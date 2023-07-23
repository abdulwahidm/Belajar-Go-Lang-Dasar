package main

import "fmt"

func main() {
	// Perulangan adalah proses mengulang-ulang eksekusi blok kode tanpa henti,
	// selama kondisi yang dijadikan acuan terpenuhi. Biasanya disiapkan variabel
	// untuk iterasi atau variabel penanda kapan perulangan akan diberhentikan.

	// Di Go keyword perulangan hanya for saja, tetapi meski demikian,
	// kemampuannya merupakan gabungan for , foreach , dan while ibarat bahasa pemrograman lain.

	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// For dengan argumen hanya kondisi
	/*
	var i = 0
	for i < 5 {
		fmt.Println("Angka", i)
		i++
	} 
	*/

	// For dengan tanpa argumen
	/*
	var i = 0
	for {
		fmt.Println("Angka", i)
		i++

		if i == 5 {
			break
		}
	}
	*/

	// Keyword break dan continue
	for i := 1; i <= 12; i++ {
		if i % 2 == 1 {
			continue
		}
		if i > 10 {
			break
		}
		fmt.Println("Angka:", i)
	}

	// For loop dengan teknik pemeberian label
	outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
			break outerLoop
		}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}