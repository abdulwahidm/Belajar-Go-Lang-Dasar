package main

import "fmt"

func main() {
	greeting := getGoodMorning("Abdul")
	fmt.Println(greeting)
}

func getGoodMorning(name string) string {
	return "Good morning "+ name
}