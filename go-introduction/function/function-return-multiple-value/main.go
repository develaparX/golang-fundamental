package main

import "fmt"

func main() {
	num := []int{4, 1, 3, 7, 5}
	kecil, _ := minMax(num) //kecil untuk return min, dan besar untuk return max // gunakan underscore untuk memanggil salah satu
	fmt.Println("Kecil :", kecil)
	// fmt.Println("Besar :", besar)
}

func minMax(numbers []int) (min int, max int) {
	min = numbers[0] //karena variable sudah di declare di return
	max = numbers[0]

	for _, n := range numbers {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return //karena variable sudah di declare jadi tidak perlu ditulis
}
