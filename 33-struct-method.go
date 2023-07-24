package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

/*
Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function
Method adalah function
*/

func (customer Customer) sayHi(name string){
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu(){
	fmt.Println("Huuuuuu from", a.Name)
}

func main() {
	abdul := Customer {
		Name: "Abdul",
		Address: "Bekasi",
		Age: 25,
	}
	abdul.sayHi("Eko")
	abdul.sayHuuu()

	fmt.Println(abdul.Name)
	fmt.Println(abdul.Address)
	fmt.Println(abdul.Age)

	//joko := Customer{
	//	Name:    "Joko",
	//	Address: "Cirebon",
	//	Age:     35,
	//}
	//fmt.Println(joko)
	//
	//budi := Customer{"Budi", "Jakarta", 35}
	//fmt.Println(budi)
}