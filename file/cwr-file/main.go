package main

import (
	"bufio"
	"fmt"
	"os"
)

var path = "C:/enigmacamp/Golang Fundamental/data/"
var fileName = "example.txt"
var filePath = path + fileName

func main() {

	// var input string
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Print("Masukkan nama :")
	// scanner.Scan()
	// input = scanner.Text()

	// writeFile(input)
	//ini keatas jika ingin menulis ke file

	readFIle()

}

func createFile() {

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Println("File", fileName, "berhasil dibuat.")
	}
}

func writeFile(input string) {

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666) //membuka file, disetela filepath settingan agar tidak replace data sebelumnya
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() //menutup file

	writer := bufio.NewWriter(file)  //menulis argumen ke file
	writer.WriteString(input + "\n") //ada \n agar data memiliki baris berbeda
	writer.Flush()                   //menulis ke file

	//agar data tidak tertimpa

}

func readFIle() {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
