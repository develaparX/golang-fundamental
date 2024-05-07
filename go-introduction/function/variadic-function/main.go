package main

import "fmt"

func main() {
	total := sum(2, 5, 7, 8)
	fmt.Println("Total: ", total)
}

func sum(numbers ...int) int { //bisa menambahkan elemen tidak terbatas karena bertipe slice
	// fmt.Println(numbers[0])
	// fmt.Println(numbers[1])
	// fmt.Println(numbers[2])
	// fmt.Println(numbers[3])
	// fmt.Printf("%T \n", numbers) //untuk mengetahui tipe data variadic yang mana slice

	total := 0
	for _, number := range numbers {
		total += number
	}
	return total

}
