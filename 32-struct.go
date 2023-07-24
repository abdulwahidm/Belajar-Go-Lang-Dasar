package main

import "fmt"

/**
Struct adalah sebuah template data yang digunakan untuk menggabungkan nol 
atau lebih tipe data lainnya dalam satu kesatuan
Struct biasanya representasi data dalam program aplikasi yang kita buat
Data di struct disimpan dalam field  
Sederhananya struct adalah kumpulan dari field
*/

type Person struct {
	Name, Address 	string
	Age				int
}

/**
Struct adalah template data atau prototype data
Struct tidak bisa langsung digunakan
Namun kita bisa membuat data/object dari struct yang telah kita buat
*/

func main() {
	var abdul Person
	abdul.Name = "Abdul Wahid Muhaemin"
	abdul.Address = "Bekasi"
	abdul.Age = 25

	// fmt.Println(abdul.Name)
	// fmt.Println(abdul.Address)
	// fmt.Println(abdul.Age)

	// Struct Literal
	budi := Person {
		Name : "Budiman",
		Address: "Jakarta",
		Age: 21,
	}
	 fmt.Println(budi)
}

