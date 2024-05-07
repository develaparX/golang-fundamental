package main

import "fmt"

func main() {
	var sum func(int, int) int = add //tipe func harus sama dengan valuenya yaitu di func add yg dipanggil

	fmt.Println("Hasil :", sum(3, 4))
}

func add(a, b int) int {
	return a + b
}
