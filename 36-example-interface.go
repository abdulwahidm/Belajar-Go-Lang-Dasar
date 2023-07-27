package main

import "fmt"

type BangunDatar interface {
	Luas()		int
}

type Persegi struct {
	Sisi 		int
}

type Lingkaran struct {
	Diameter 	int
}

func (l Lingkaran) jariJari() int {
	return l.Diameter /2
}

func (persegi Persegi) Luas() int {
	return persegi.Sisi * persegi.Sisi
}



func main() {

	var newPersegi BangunDatar
	 
	newPersegi = Persegi{Sisi: 4}
	fmt.Println(newPersegi.Luas())
	
}

