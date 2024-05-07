package main

import "fmt"

func main() {

	// fmt.Println("Hasil Penjumlahan :", add(5, 2)) // work

	total := add(5, 2) //menyimpan nilai yang di return func add di variable total
	fmt.Println("Hasil Penjualan :", total)

	total2 := multiply(3, 8)
	fmt.Println("Hasil Perkalian :", total2)

	calculation := add(7, multiply(4, 2))
	fmt.Println("Hasil Kalkulasi : ", calculation)
}

func add(a int, b int) int {
	result := a + b
	// result := "a+b" // jika data tidak sesuai tipe data, maka akan error
	return result
}

func multiply(a, b int) int {
	result := a * b
	return result
}
