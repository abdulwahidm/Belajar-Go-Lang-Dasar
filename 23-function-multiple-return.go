package main

import "fmt"
import "math"

func main() {
	diameter := 10.00
	luas, keliling := calculate(diameter)

	// fmt.Println("luas lingkaran:",luas,"\ndan kelilingnya adalah:",keliling)
	fmt.Printf("luas lingkaran\t\t: %.2f \n", luas)
	fmt.Printf("keliling lingkaran\t: %.2f \n", keliling)
}

func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d / 2, 2)
	// hitung keliling
	var circumference = math.Pi * d
	// kembalikan 2 nilai
	return area, circumference
	}