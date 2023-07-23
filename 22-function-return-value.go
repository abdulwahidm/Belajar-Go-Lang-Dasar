package main

import (
	"fmt"
	"math/rand"
	"time"
)
/*
Sebuah fungsi bisa dirancang tidak mengembalikan nilai balik (void), atau bisa
mengembalikan suatu nilai. Fungsi yang memiliki nilai kembalian, harus
ditentukan tipe data nilai baliknya pada saat deklarasi.
Program berikut merupakan contoh penerapan fungsi yang memiliki return value.
*/

func main() {

	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	
}

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}