package main

import "fmt"

type Person struct {
	Name, Address 	string
	Age				int
}

func main() {
	var abdul Person
	abdul.Name = "Abdul Wahid Muhaemin"
	abdul.Address = "Bekasi"
	abdul.Age = 25

	// fmt.Println(abdul.Name)
	// fmt.Println(abdul.Address)
	// fmt.Println(abdul.Age)
}