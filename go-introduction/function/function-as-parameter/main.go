package main

import "fmt"

func main() {
	// fmt.Println("Genap:", isEven(8))
	// fmt.Println("Ganjil:", isOdd(8))

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := filter(numbers, isEven)
	fmt.Println("Genap:", evenNumbers)

	oddNumbers := filter(numbers, isOdd)
	fmt.Println("Ganjil:", oddNumbers)
}

func filter(numbers []int, f func(int) bool) []int {
	var result []int

	for _, num := range numbers {
		if f(num) {
			result = append(result, num)
		}
	}
	return result
}

func isEven(number int) bool {
	return number%2 == 0
}

func isOdd(number int) bool {
	return number%2 != 0
}
