package main

import (
	"fmt"
)

/*
Panic function adalah function yang bisa kita gunakan untuk menghentikan program
Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi
*/

func endApp(){
	// message := recover()
	// if message != nil {
	// 	fmt.Println("Error dengan message:", message)
	// }
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Abdul")
}
