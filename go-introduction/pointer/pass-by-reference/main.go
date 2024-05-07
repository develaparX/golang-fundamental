package main

import "fmt"

func main() {
	var x int = 4
	var y *int = &x

	fmt.Println("x:", x)
	fmt.Println("alamat x:", &x)
	fmt.Println("y:", y)
	fmt.Println("alamat y:", &y)

	fmt.Println("Nilai reference dari pointer y:", *y)

	*y = *y + 1 //jika merubah nilai reference y maka di x juga akan berubah // pass by reference

	fmt.Println("x:", x)
	fmt.Println("*y:", *y)

	fmt.Println("==================================================")

	var a int = 3
	incerease(&a)
	fmt.Println("a:", a) //nilai ini akan sama yaitu 4, ini merupakan pass by reference
}

func incerease(n *int) {
	*n = *n + 1
	fmt.Println("n di function increase:", *n)
}
