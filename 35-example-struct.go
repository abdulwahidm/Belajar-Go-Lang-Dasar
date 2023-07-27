package main

import (
	"math"
)

type Hitung interface {
	Luas 		float64
	Keliling 	float64
}

type Lingkaran struct {
	Diameter 	float64
}

func (l Lingkaran) jariJari() float64{
	return l.Diameter / 2
}

func (l Lingkaran) luas() float64{
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func main() {
	circle := Lingkaran{ 10.0 }
	
	fmt.Println(circle.jariJari())
	fmt.Println(circle.luas())
	
}