package main

import "fmt"

func main() {
	if false {
		fmt.Println("Kode dijalankan")
	}

	fmt.Println("Done")

	fmt.Println("=====================================")

	if result := "lulus"; result == "lulus" {
		fmt.Println("Selamat anda,", result)
	} else {
		fmt.Println("Maaf Anda,", result)
	}

	// fmt.Println(result) //akan error

	fmt.Println("=====================================")

	hour := 6

	if hour > 8 && hour < 17 {
		fmt.Println("Masih dalam rentang waktu yang diperbolehkan,", hour)
	} else {
		fmt.Println("Diluar rentang waktu")
	}

	fmt.Println("=====================================")

	GPA := 2.5
	var graduationStatus string

	if GPA >= 3.5 && GPA <= 4.0 {
		graduationStatus = "CUMLAUDE"
	} else if GPA >= 3.0 && GPA < 3.5 {
		graduationStatus = "MEMUASKAN"
	} else if GPA > 2.75 && GPA < 3.0 {
		graduationStatus = "CUKUP MEMUASKAN"
	} else {
		graduationStatus = "KURANG MEMUASKAN"
	}
	fmt.Printf("Selamat kamu lulus dengan predikat %s dengan IPK %.2f\n", graduationStatus, GPA)

	fmt.Println("=====================================")

	//nested condition atau condition dalam condition
	x := 3
	y := -1

	// memiliki alur pengecekan x dahulu baru y
	if x > 0 {
		if y > 0 {
			fmt.Println("x dan y lebih besar dari 0")
		} else {
			fmt.Println("x lebih besar dari 0 dan y kurang dari atau sama dengan 0")
		}
	}

	fmt.Println("=====================================")

	//switch case
	var poin = 9
	// switch poin {
	// case 8:
	// 	{
	// 		fmt.Println("bagus")
	// 	}
	// case 7, 6, 5:
	// 	{
	// 		fmt.Println("cukup")
	// 	}
	// default:
	// 	{
	// 		fmt.Println("kurang")
	// 	}
	// }

	switch {
	case poin >= 8:
		{
			fmt.Println("bagus")
		}
		fallthrough
	case poin >= 6 && poin < 8:
		{
			fmt.Println("cukup")
		}
	case poin >= 4 && poin < 6:
		{
			fmt.Println("Kurang")
		}
	default:
		{
			fmt.Println("gagal")
			fmt.Println("Kamu perlu belajar lebih giat lagi")
		}
	}

	// coding challange enigma
	//Bantulah Susi untuk menentukan jadwal sehari-harinya dengan sebuah program sederhana. Akan diberikan inputan jam berupa integer lalu tampilkanlah pada console kegiatan yang akan dilakukan Susi sesuai jam yang diberikan. Kegiatan yang akan dilakukan Susi adalah berikut :

	// Jika jam 4-5 akan mencetak “Bangun Pagi”

	// Jika jam 6-7 akan mencetak “Mandi dan Sarapan”.

	// Jika jam 8-11 akan mencetak “Berangkat Sekolah”.

	// Jika jam 12 akan mencetak “Pulang Sekolah”.

	// Jika jam 13-15 akan mencetak “Tidur Siang”.

	// Jika Jam 16-17 akan mencetak “Waktu Main”

	// Selain dari itu akan mencetak “Waktu Belajar dan Istirahat”

	// Bila user menginput melebihi 24 jam maka akan mencetak “Hanya ada 24 jam dalam sehari”

	fmt.Println("=====================================")

	jam := 6

	if jam >= 4 && jam <= 5 {
		fmt.Println("Bangun Pagi")
	} else if jam >= 6 && jam <= 7 {
		fmt.Println("Mandi dan Sarapan")
	} else if jam >= 8 && jam <= 11 {
		fmt.Println("Berangkat Sekolah")
	} else if jam == 12 {
		fmt.Println("Pulang Sekolah")
	} else if jam >= 13 && jam <= 15 {
		fmt.Println("Tidur Siang")
	} else if jam >= 16 && jam <= 17 {
		fmt.Println("Waktu Main")
	} else if jam >= 24 {
		fmt.Println("Hanya ada 24 jam dalam sehari")
	} else {
		fmt.Println("Waktu Belajar dan Istirahat")
	}
}
