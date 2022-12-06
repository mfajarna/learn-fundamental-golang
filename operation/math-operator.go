package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 2
	var c = a * b
	fmt.Println(c)

	// Augmented Assigment
	var z = 10
	z += 10
	fmt.Println(z)

	// Unary Operator
	z++
	fmt.Println(z)

	var negative = -100
	var positive = +100
	fmt.Println(negative)
	fmt.Println(positive)
}
