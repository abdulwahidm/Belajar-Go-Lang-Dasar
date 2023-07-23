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

/*
Selain sebagai penanda nilai balik, keyword return juga bisa dimanfaatkan
untuk menghentikan proses dalam blok fungsi di mana ia dipakai. Contohnya bisa
dilihat pada kode berikut.
*/

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}
	var res = m / n
		fmt.Printf("%d / %d = %d\n", m, n, res)
	}