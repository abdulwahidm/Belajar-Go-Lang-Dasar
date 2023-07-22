package main
 
import "fmt"

func main() {
// 	Pada Array atau Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
// Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan
// Cara menggunakan map cukup dengan menuliskan keyword map diikuti tipe data key dan value-nya.
// Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak-banyaknya, asalkan kata kunci nya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru.

	person := map[string]string{
		"name" : "Abdul",
		"address" : "Bekasi",
	}
	person["title"] = "programmer" // menambahkan data baru

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Println(len(person)) // fungsi len mengembalikan jumlah data pada map
	delete(person, "address") // menghapus sebuah data  pada map
	fmt.Println("jumlah data map setlah dihapus:", len(person))

	// Cara lain dalam membuat map
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))
}