package main

import "fmt"

func main() {
	// Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama
	// Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh Array tersebut
	// Daya tampung Array tidak bisa bertambah setelah Array dibuat

	var names [3]string
	names[0] = "Abdul"
	names[1] = "Wahid"
	names[2] = "Muhaemin"
	fmt.Println(names) // [Abdul Wahid Muhaemin]

	var lagi [10]string
	fmt.Println(len(lagi)) // 10

	// Cara inisialisai yang lain
	languages := [...] string {
		"Javascript",
		"Python",
		"Golang",
		"PHP",
		"Ruby",
		"C++",
	}
	fmt.Println(languages) // [Javascript Python Golang PHP Ruby C++]
	fmt.Println(len(languages)) // 6

	for index, lang := range languages {
		fmt.Println(index+1, lang)
		// 1 Javascript
		// 2 Python
		// 3 Golang
		// 4 PHP
		// 5 Ruby
		// 6 C++
	}
}