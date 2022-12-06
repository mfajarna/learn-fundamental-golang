package main

import "fmt"

func main() {
	var string1 = "Budi"
	var string2 = "Budi"
	var string3 = "Tono"

	fmt.Println(string1 == string2)
	fmt.Println(string2 == string3)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 > value2)
}
