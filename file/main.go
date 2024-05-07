package main

import (
	"fmt"
	"strconv"
)

var names []string

func main() {
	var input string
	fmt.Println("names:", names)
	fmt.Print("Masukkan jumlah nama : ")
	fmt.Scanln(&input)
	len, _ := strconv.Atoi(input) //mengkonversi input dan disimpan dalam var len
	for i := 0; i < len; i++ {
		fmt.Print("Masukkan Nama : ")
		fmt.Scanln(&input)
		names = append(names, input) //menyimpan dalam slice names
	}
	fmt.Println("names:", names)
	// jika di run maka data tidak tersimpan/hilang atau hanya bersifat sementara/volatile

}
