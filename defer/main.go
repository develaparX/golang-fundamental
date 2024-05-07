package main

import "fmt"

func main() {
	// defer fmt.Println("Done") //jika ditambah defer akan selalu di run di akhir setelah func, bisa diletakkan dimanapun
	// fmt.Println("Start")
	// fmt.Println("Processing")

	// =============================================================================================

	// num := 4

	// defer fmt.Println("Result Defer:", num) //akan di run setelah ini, tapi ditampilkan di akhir

	// num += 10
	// num *= 2

	// fmt.Println("Result Main:", num)

	// =============================================================================================

	// num1 := 4
	// num2 := 0

	// defer fmt.Println("Done") // jika diletakkan dibawah error, tidak akan tampil. tapi jika disini tetap tampil
	// fmt.Println("Start")
	// fmt.Println(num1 / num2) //akan error

	// =============================================================================================
	// defer bisa digunakan untuk memanggil function, tetap akan dijalankan di akhir
	// defer add(7, multiply(2, 3)) // jika seperti ini akan mencetak multiply(func) dulu baru defernya
	// add(3, 4)

	// =============================================================================================
	//banyak defer akan dieksekusi secara LIFO (Last In First Out)
	// fmt.Println("Start")
	// defer fmt.Println("Done") //4
	// defer add(1, 2)           // lalu ini #3
	// defer add(3, 4)           //kemudian ini #2
	// defer multiply(3, 3)      //akan ini dulu #1

	// =============================================================================================
	// fmt.Println("Start")
	// defer loop()
	// fmt.Println("Done")

	//cc
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

func add(num1 int, num2 int) {
	result := num1 + num2
	fmt.Println(result)
}

func multiply(num1 int, num2 int) int {
	result := num1 * num2
	fmt.Println(result)
	return result
}

func loop() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done Loop")
}
