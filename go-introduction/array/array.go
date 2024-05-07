package main

import "fmt"

func main() {
	var fruits = [4]string{"Apel", "Pisang", "Anggur", "Semangka"} // indeks array dimulai dari 0 itungannya

	fmt.Println(fruits)
	fmt.Println(fruits[1])
	fruits[2] = "Jeruk" //merubah data didalam array di index 2
	fmt.Println(fruits)

	fmt.Println("===============================================")

	var scores [3]int //membuat array dengan panjang 3 element bertipe integer tanpa ada valuenya
	scores[0] = 87
	scores[1] = 78
	scores[2] = 92
	// scores[3] = 70 // jika ada value diluar nilai index atau diluar kapasitas array maka akan error

	fmt.Println(scores) // akan ada default value yaitu 0 jika tidak ada valuenya

	fmt.Println("===============================================")

	//mendeklarasikan array tanpa batas arraynya dahulu
	var days = [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	// days[7] = "Libur"
	fmt.Println(days)
	fmt.Println("Jumlah elemen :", len(days)) // melihat element dalam array days, maka jika ditambah elemen baru diluar array akan error atau dalam hal ini jika ditambah index 7 akan error

	fmt.Println("===============================================")

	// kombinasi array dan looping
	//membuat perulangan dimulai dari 0 sampai len(days), disini len(days)merupakan panjang dari array days yang mana nilainya adalah 7
	for i := 0; i < len(days); i++ {
		fmt.Printf("Elemen %d: %s\n", i, days[i])
	}

	fmt.Println("===============================================")
	// i merupakan indeks dan day yang merepresentasikan tiap elemen didalam days
	for i, day := range days {
		fmt.Printf("Elemen %d: %s\n", i, day)
	}

	fmt.Println("===============================================")

	//looping tanpa menggunakan indeks
	for _, day := range days { // yang tercetak akan nama harinya saja
		fmt.Printf("Nama Hari : %s\n", day)

	}

	fmt.Println("===============================================")

	// menggunakan indeksnya saja

	for i := range days {
		fmt.Printf("Index hari ke : %d\n", i)
	}

	fmt.Println("===============================================")

	//penggunaan looping array dan pengkondisian

	var numbers = [5]int{3, 8, 2, 7, 4}

	for _, number := range numbers {
		if number%2 == 0 { // yang tercetak akan bilangan genap
			fmt.Println(number)
		}
	}

	fmt.Println("===============================================")

	fmt.Println("Sebelum :", numbers)
	for i := 0; i < len(numbers); i++ {
		numbers[i] *= 2
	}
	fmt.Println("Setelah :", numbers)

	fmt.Println("===============================================")

	var matrix = [2][3]int{ //array multidimensi 2 baris 3 kolom atau ada array dalam array
		{3, 2, 3},
		{3, 4, 5},
	}

	fmt.Println(matrix)
	fmt.Println(matrix[1][1]) // untuk mengakses nomor 4 yaitu di index 1 dan index 1
	fmt.Println(matrix[1][2])
}
