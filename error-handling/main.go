package main

import (
	"errors"
	"fmt"
)

func main() {
	//akan terjadi panic! dan dibawah ini tidak akan di jalankan, jika tidak ada error handling
	result, err := divide(5, 0) // jika nilai b bukan 0 maka error tidak akan tampil
	if err != nil {
		fmt.Println("Error:", err) //jika error akan ada ini
		return
	}
	fmt.Println("Hasil Pembagian :", result)
}

func divide(a, b int) (int, error) { // return error handling di declare pesannya di atas
	if b == 0 {
		return 0, errors.New("pembagian dengan angka 0") //error jika b sama dengan 0
	}
	return a / b, nil // nil sama dengan kosong atau jika tidak error akan mereturn nilai kosong
}
