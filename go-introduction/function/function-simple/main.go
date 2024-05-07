package main

import "fmt"

var name = "Kenji"

func main() {
	helloWorld()
	fmt.Println("main: ", name)

	greet("Kenji", 25)
	greet("John", 19)
	greet("Doe", 20)
	// greet(21) // akan error, pastikan argumen sesuai type data parameter

	// total := add(2,6)
}

func helloWorld() {
	//  var name = "Kenji" // memiliki scope local atau hanya dapat diakses di func helloWorld
	var name = "Ferdian" // akan jadi prioritas daripada yang global
	fmt.Println("Hello World!")
	fmt.Println("HelloWorld :", name)
}

func greet(name string, age int) {
	fmt.Println("Hello my name is", name, "I'm", age, "years old")
}

func add(a int, b int) {
	fmt.Println("Result Add :", a+b)
}
