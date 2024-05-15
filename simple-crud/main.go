package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Id          int
	Title       string
	Author      string
	ReleaseYear string
	Pages       int
}

var books []Book // menampung data buku

var fileName string = "data.csv" // ingin menimpan data dalam bentuk csv

func main() {

	//membuat file data.csv
	// Panggil fungsi createFile
	// createFile(fileName)
	// err := createFile(fileName)
	// if err != nil {
	// 	fmt.Println("Gagal membuat file:", err)
	// 	return
	// }

	// addNewBook() // menambahkan buku baru

	//melihat semua buku
	loadDataFromCSV(fileName)
	viewAllBook()

	updateBookById(1)

}

func createFile(fileName string) error {
	// Memeriksa apakah file sudah ada
	if _, err := os.Stat(fileName); err == nil {
		// Jika file sudah ada, langsung kembalikan nil
		fmt.Println("File sudah ada:", fileName)
		return nil
	} else if os.IsNotExist(err) {
		// Jika file belum ada, buat file
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer file.Close()

		fmt.Println("File", fileName, "berhasil dibuat")
		return nil
	} else {
		// Jika terjadi kesalahan selain file tidak ada
		fmt.Println(err)
		return err
	}
}

func addNewBook() error { //menambah buku baru
	var newBook Book

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Book Details :")
	fmt.Print("Book Id :")
	scanner.Scan()
	newBook.Id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Book Title :")
	scanner.Scan()
	newBook.Title = scanner.Text()

	fmt.Print("Book Author :")
	scanner.Scan()
	newBook.Author = scanner.Text()

	fmt.Print("Release Year :")
	scanner.Scan()
	newBook.ReleaseYear = scanner.Text()

	fmt.Print("Book pages :")
	scanner.Scan()
	newBook.Pages, _ = strconv.Atoi(scanner.Text())

	_, err := findBookById(newBook.Id)
	if err != nil {
		books = append(books, newBook)
	} else {
		return fmt.Errorf("Book with id : %d already exist", newBook.Id)
	}

	err = saveDataToCSV(fileName)
	if err != nil {
		return err
	}

	fmt.Println("Book added successfully")

	return nil
}

func saveDataToCSV(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}
	defer file.Close()

	for _, book := range books {
		row := strconv.Itoa(book.Id) + "," + book.Title + "," + book.Author + "," + book.ReleaseYear + "," + strconv.Itoa(book.Pages) + "\n" //data perbaris dalam data.csv
		file.WriteString(row)
	}
	return nil
}

func loadDataFromCSV(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening csv file : %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() { //menampilkan data.csv

		record := strings.Split(scanner.Text(), ",") //membagi string yang ditangkap scanner.text berdasarkan "," koma
		// fmt.Println(record) //menampilkan data apa adanya (string semua)

		id, _ := strconv.Atoi(record[0])    //convert ke int
		pages, _ := strconv.Atoi(record[4]) // convert ke int

		book := Book{
			Id:          id,
			Title:       record[1],
			Author:      record[2],
			ReleaseYear: record[3],
			Pages:       pages,
		}
		books = append(books, book)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error opening csv file : %w", err)
	}

	return nil
}

func viewAllBook() error { //melihat semua buku
	if len(books) == 0 {
		return fmt.Errorf("no books available")
	}
	for i, book := range books {
		fmt.Println(strings.Repeat("=", 50)) //pembatas
		fmt.Println("Book - ", i+1)
		fmt.Println("Book Id :", book.Id)
		fmt.Println("Book Title : ", book.Title)
		fmt.Println("Book Author :", book.Author)
		fmt.Println("Release Year :", book.ReleaseYear)
		fmt.Println("Pages :", book.Pages)
		fmt.Println(strings.Repeat("=", 50)) // pembatas

	}
	return nil
}

func findBookById(id int) (Book, error) { //pengecekan apakah id sudah ada apa belum
	for _, book := range books {
		if book.Id == id {
			return book, nil

		}
	}
	return Book{}, fmt.Errorf("id : %d not found", id)
}

func updateBookById(id int) error {
	for i, book := range books {
		if book.Id == id {
			fmt.Println("Book found, enter new details:")
			var updatedBook Book

			scanner := bufio.NewScanner(os.Stdin)

			fmt.Print("New Book Title :")
			scanner.Scan()
			updatedBook.Title = scanner.Text()

			fmt.Print("New Book Author :")
			scanner.Scan()
			updatedBook.Author = scanner.Text()

			fmt.Print("New Release Year :")
			scanner.Scan()
			updatedBook.ReleaseYear = scanner.Text()

			fmt.Print("New Book pages :")
			scanner.Scan()
			updatedBook.Pages, _ = strconv.Atoi(scanner.Text())

			books[i] = updatedBook

			err := saveDataToCSV(fileName)
			if err != nil {
				return err
			}

			fmt.Println("Book updated successfully")
			return nil
		}
	}
	return fmt.Errorf("Book with id : %d not found", id)
}
