package main

import (
	"fmt"
)

/*
Panic function adalah function yang bisa kita gunakan untuk menghentikan program
Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

Recover adalah function yang bisa kita gunakan untuk menangkap data panic
Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
*/

func endApp(){
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool){
	defer endApp() // recover() disisipkan pada fungsi defer
	if error {
		panic("APLIKASI ERROR") // jika panic terpanggil maka alur program akan berhenti
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Abdul")
}
