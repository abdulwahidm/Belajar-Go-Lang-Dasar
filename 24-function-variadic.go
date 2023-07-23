package main

import "fmt"
import "strings"

/*
Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array.
Apa bedanya dengan parameter biasa dengan tipe data Array?
Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
JIka parameter menggunakan varargs, kita bisa langsung mengirim data nya, 
jika lebih dari satu, cukup gunakan tanda koma
*/
func main() {
	total := sumAll(10, 90, 30, 50, 40)
	fmt.Println(total)

	slice := []int{10,20,30,40,50}
	total = sumAll(slice...)
	fmt.Println(total)

	// Memanggil function yourhobbies
	yourHobbies("Abdul", "Coding", "Reading", "Haking")
}

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")
	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}
