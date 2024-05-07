package main

import "fmt"

func main() {
	var numbers = [4]int{4, 3, 7, 1}
	var anotherNumbers = numbers

	fmt.Println("Sebelum =====================================================")
	fmt.Println("numbers:", numbers)
	fmt.Println("anotherNumbers:", anotherNumbers)

	numbers[1] = 9 // array secara default pass by value
	fmt.Println("Sesudah =====================================================")
	fmt.Println("numbers:", numbers)
	fmt.Println("anotherNumbers:", anotherNumbers)

	fmt.Println("==============================================================")

	var scores = [4]int{7, 8, 6, 9}
	multiplyBy10(scores)
	fmt.Println("scores:", scores)
}

func multiplyBy10(n [4]int) {
	for i := range n {
		n[i] = n[i] * 10
	}
	fmt.Println("n in function:", n)
}
