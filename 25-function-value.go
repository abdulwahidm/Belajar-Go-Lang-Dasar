package main

import "fmt"

// Function adalah first class citizen
// Function juga merupakan tipe data, dan bisa disimpan di dalam variable

func main() {
	greeting := getGoodMorning("Abdul")
	fmt.Println(greeting)
}

func getGoodMorning(name string) string {
	return "Good morning "+ name
}