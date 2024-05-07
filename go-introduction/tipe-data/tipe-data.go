package main

import "fmt"

func main() {
	fmt.Println("Tipe Data")
	fmt.Println("Numerik")
	var positiveNumber uint8 = 90 //tipe data ini tidak bisa minus (browsing tentang tipe data untuk lebih lengkap)
	var negativeNumber int = -90  // bisa punya value negatif dan lebih banyak, tapi makan memory lebih besar
	fmt.Printf("Bilangan Positif : %d \n", positiveNumber)
	fmt.Printf("Bilangan Negatif : %d \n", negativeNumber)

	var decimalNumber = 3.90
	fmt.Printf("Bilangan Pecahan : %f \n", decimalNumber)   // %f disini untuk memanggil bilangan pecahan atau bilangan desimal
	fmt.Printf("Bilangan Pecahan : %.2f \n", decimalNumber) // %.2f disini untuk mencetak bilangan dibelakang koma 2 angka saja

	fmt.Println("======================================================")

	fmt.Println("Boolean")
	var exist = true
	fmt.Printf("exist ? %t\n", exist) // %t untuk memanggil tipe data boolean

	fmt.Println("======================================================")

	fmt.Println("String")
	var message string = "Halo"
	fmt.Printf("message : %s\n", message) // %s untuk memanggil tipe data string
	var otherMessage = `Nama saya "Kenji", 
	Salam Kenal, 
	Mari belajar di "Enigma Camp"` // menggunakan ``(Backtick) untuk membuat semua karakter didalamnya tidak di escape atau semua akan terdeteksi sebagai string

	fmt.Println(otherMessage)

	fmt.Println("Kenji" + "Ferdian")

	fmt.Println(len(message)) //len untuk mengetahui jumlah karakter dalam string
}
