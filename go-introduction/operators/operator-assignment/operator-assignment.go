package main

import "fmt"

func main() {
	fmt.Println("Operator Assignment")

	// var a int = 10
	// a = a - 3
	// fmt.Println("Nilai a =", a)

	var a int = 5
	var b int

	b = a
	fmt.Println("Nilai b =", b)
	b += a // b = b + a
	fmt.Println("Nilai dari b =", b)
	b -= a // b = b - a
	fmt.Println("Nilai dari b =", b)
	b *= a // b= b * a
	fmt.Println("Nilai dari b =", b)
	b /= a // b = b / a
	fmt.Println("Nilai dari b =", b)
	b %= a // b = b % a
	fmt.Println("Nilai dari b =", b)

}
