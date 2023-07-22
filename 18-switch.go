package main

import "fmt"

func main() {

	// Selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan Switch Expression
	// Switch expression sangat sederhana dibandingkan if
	// Biasanya switch  expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable
	point := 1
	switch point {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4: // case dapat menampung banyak kondisi
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")	
	}

	//  Switch dengan kondisi:
	var value = 1
	switch {
	case value > 8:
		fmt.Println("great!")
	case value == 8:
		fmt.Println("perfect")
	case (value < 8) && (value > 3):
		fmt.Println("awesome")
	case value < 5:
		fmt.Println("you need to learn more")
	default: // gunakan {} jika nilai kembalian lebih dari satu agar terlihat lebih clean
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}