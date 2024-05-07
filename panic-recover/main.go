package main

import "fmt"

func main() {
	//defer anonymous func
	// defer func() { // jika ada ini maka yang tampil pesan errornya dan tidak dengan errornya
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Terjadi Panic:", r)
	// 	}
	// }()

	var input string
	fmt.Print("Masukkan Nama :")
	fmt.Scanln(&input)

	fmt.Println("Start")
	if !isEmpty(input) { //memanggil func isEmpty yang mana berisi panic, code setelah panic tak akan dijalankan
		fmt.Println(input)
	}
	fmt.Println("Done")
}

func isEmpty(input string) (empty bool) {
	defer func() { // jika defer disini, maka program akan terus berjalan dan "done" akan tercetak
		if r := recover(); r != nil {
			fmt.Println("Terjadi Panic:", r)
			empty = true
		}
	}()

	if input == "" {
		panic("input tidak boleh kosong")

	}
	empty = false
	return
}
