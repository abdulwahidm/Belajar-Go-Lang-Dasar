package main

import (
	"fmt"
	"math"
)

/*
Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung.
Sebuah interface berisikan definisi-definisi method.
Biasanya interface digunakan sebagai kontrak.
*/

type Hitung interface {
	Luas() float64
	Keliling() float64
}

type Lingkaran struct {
	Diameter float64
}

func (l Lingkaran) jariJari() float64 {
	return l.Diameter / 2
}

func (l Lingkaran) Luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l Lingkaran) Keliling() float64 {
	return math.Pi * l.Diameter
}

func main() {
	
	var bangunDatar Hitung

	bangunDatar = Lingkaran{14.0}
	fmt.Println("===== Lingkaran =====")
	fmt.Println("jari-jari	:", bangunDatar.(Lingkaran).jariJari())
	fmt.Println("Luas		:", bangunDatar.Luas())
	fmt.Println("Keliling	:", bangunDatar.Keliling())


}