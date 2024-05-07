package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var firstName string
	var lastName string
	var age int

	fmt.Print("Enter Your Name : ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Print("Enter Your Age : ")
	fmt.Scanln(&age)
	birthYear := (2023 - age)

	fmt.Println("Your Name is", firstName, lastName, "You are born in", birthYear)

	fmt.Println("===========================================================")

	var fullName string
	var address string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Data diri pegawai enigma camp")
	fmt.Print("Masukkan Nama Anda : ")
	scanner.Scan()
	fullName = scanner.Text()
	fmt.Print("Masukkan Umur Anda : ")
	scanner.Scan()
	age, _ = strconv.Atoi(scanner.Text())
	fmt.Print("Masukkan Alamat Anda :")
	scanner.Scan()
	address = scanner.Text()

	fmt.Printf("%s | %d | %s", fullName, age, address)

}
