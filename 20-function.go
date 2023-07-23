package main

import (
	"fmt"
	"strings"
)

/*
Fungsi merupakan aspek penting dalam pemrograman. Definisi fungsi sendiri
adalah sekumpulan blok kode yang dibungkus dengan nama tertentu. Penerapan
fungsi yang tepat akan menjadikan kode lebih modular dan juga dry (kependekan
dari don't repeat yourself), tak perlu menuliskan banyak kode yang kegunaannya
berkali-kali, cukup sekali saja lalu panggil sesuai kebutuhan.
*/

func main() {
	names := []string{"John", "Wick"}
	printMessage("halo", names)		
}

/*
Pada kode di atas, sebuah fungsi baru dibuat dengan nama printMessage
memiliki 2 buah parameter yaitu string message dan slice string arr .
Fungsi tersebut dipanggil dalam main , dengan disisipkan 2 buah data sebagai
parameter, data pertama adalah string "hallo" yang ditampung parameter
message , dan parameter ke 2 adalah slice string names
oleh parameter arr .
*/

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}