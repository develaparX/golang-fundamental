package main

import "fmt"

type Student struct {
	// Isi struct ini
	name string
	age  int
}

func (s *Student) Introduction() {
	// Tulis kode disini
	fmt.Printf("Hello, my name is %s. Im %d years old.\n", s.name, s.age)
}

func main() {
	// Tulis kode disini
	var nama string
	var umur int
	fmt.Scanln(&nama)
	fmt.Scanln(&umur)

	student1 := Student{name: nama, age: umur}
	student1.Introduction()
}
