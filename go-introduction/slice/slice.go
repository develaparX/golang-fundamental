package main

import "fmt"

func main() {
	numbersSlice := []int{2, 7, 9, 4}
	fmt.Println(numbersSlice)

	fmt.Println("Panjang Slice:", len(numbersSlice))
	fmt.Println("Kapasitas Slice:", cap(numbersSlice))

	fmt.Println("=========================================")

	scores := make([]int, 4, 7) // jika int tidak diisi value maka akan tercetak default value yaitu 0
	scores[0] = 70
	scores[1] = 80
	scores[2] = 95
	scores[3] = 65
	fmt.Println(scores)
	fmt.Println("Panjang Scores:", len(scores))
	fmt.Println("Kapasitas Scores:", cap(scores))

	fmt.Println("=========================================")

	scores2 := make([]int, 4) // jika tidak diisi kapasitasnya, maka akan sama dengan panjangnya
	scores2[0] = 70
	scores2[1] = 80
	scores2[2] = 95
	scores2[3] = 65
	fmt.Println(scores2)
	fmt.Println("Panjang Scores 2:", len(scores2))
	fmt.Println("Kapasitas Scores 2:", cap(scores2))

	fmt.Println("=========================================")

	//perbedaan array dan slice

	heroes := [4]string{"Superman", "Batman", "Spiderman", "Iron Man"}
	// heroes[5] = "CatWoman" // di array akan error saat penambahan

	fmt.Println("Heroes :", heroes)

	villain := make([]string, 3, 5)
	villain[0] = "Thanos"
	villain[1] = "Joker"
	villain[2] = "Homelander"
	fmt.Println("Villain :", villain)
	fmt.Println("Panjang Villain :", len(villain))
	fmt.Println("Kapasitas Villain :", cap(villain))
	villain = append(villain, "Ultron")
	villain = append(villain, "Sauron", "Voldemort")
	// jika menambah melebihi kapasitas yang dideklarasikan, maka kapasitas secara default akan berubah dua kali lipat
	fmt.Println("Villain :", villain)
	fmt.Println("Panjang Villain :", len(villain))
	fmt.Println("Kapasitas Villain :", cap(villain))

	fmt.Println("=========================================")
	// array : pass by value
	// slice : pass by reference

	//contoh di array pass by value
	var numbers = [4]int{2, 1, 6, 8}
	var anotherNumbers = numbers
	fmt.Println("Numbers :", numbers)
	fmt.Println("Another Numbers :", anotherNumbers)
	anotherNumbers[1] = 11 // hanya merubah anotherNumbers, dan di array numbers tidak berubah
	fmt.Println("Numbers :", numbers)
	fmt.Println("Another Numbers :", anotherNumbers)

	fmt.Println("=========================================")
	//contoh di slice by reference

	var prices = []int{20000, 13000, 10000}
	var anotherPrices = prices
	//pass by reference, atau jika merubah elemen dalam another prices maka di prices akan beruubah juga
	fmt.Println("Prices :", prices)
	fmt.Println("Another Prices :", anotherPrices)
	anotherPrices[1] = 15000
	fmt.Println("Prices :", prices)
	fmt.Println("Another Prices :", anotherPrices)

	fmt.Println("=========================================")

	// membuat slice dari array yang sudah ada
	ages := [4]int{20, 22, 25, 19}
	sliceAges := ages[0:3] //membuat slices dari index 0, dan mengambil 3 jumlah elemen dari index 0 di ages
	fmt.Println("Ages :", ages)
	fmt.Println("Slice Ages : ", sliceAges)
	sliceAges[2] = 33 // jika dirubah slicenya maka akan terjadi pass by reference di array seperti slice
	fmt.Println("Ages :", ages)
	fmt.Println("Slice Ages : ", sliceAges)
	sliceAges = append(sliceAges, 37)
	fmt.Println("Ages :", ages)
	fmt.Println("Slice Ages : ", sliceAges)

}
