package main

import "fmt"

//secara default, slice sudah pass by reference
func main() {
	var numbers = []int{4, 3, 7, 11}
	var anotherNumbers = numbers

	fmt.Println("numbers:", numbers)
	fmt.Println("anotherNumbers:", anotherNumbers)

	numbers[1] = 9
	fmt.Println("numbers:", numbers)
	fmt.Println("anotherNumbers:", anotherNumbers)

	fmt.Println("=================================================")

	var scores = []int{7, 8, 6, 9}
	multiplyBy10(scores)

	fmt.Println("scores:", scores) //hasil var scores akan sudah berubah
}

func multiplyBy10(numbers []int) {
	for i := range numbers {
		numbers[i] = numbers[i] * 10
	}
	fmt.Println("numbers di function:", numbers)
}
