package main

import "fmt"

func main() {
	// Anonymous Function
	func() {
		fmt.Println("Hellow World")
	}()

	// Anonymous FUnction dalam variable
	halo := func() { //jika assign anony func di var maka var jadi bertipe func
		fmt.Println("Halo Dunia!")
	}

	halo()

	// Passing argumen ke dalam anonymous function
	func(word string) {
		fmt.Println(word)
	}("Enigma Camp")

	// Passing argument dalam variable function
	hello := func(name string) {
		fmt.Println("Hello, ", name)
	}
	hello("Joni")

	// Passing anonymous FUnction sebagai argumen
	greetEnglish := func(name string) string {
		return "Hello " + name
	}
	greetRussian := func(name string) string {
		return "previet " + name
	}
	greetKorean := func(name string) string {
		return "Anyeong " + name
	}
	greet("Rick", greetEnglish)
	greet("Rick", greetKorean)
	greet("Rick", greetRussian)

	add := func(num1 int, num2 int) int {
		return num1 + num2
	}
	multiply := func(num1 int, num2 int) int {
		return num1 * num2
	}
	calculate(3, 2, add)
	calculate(4, 5, multiply)
}

func greet(name string, f func(name string) string) {
	fmt.Println(f(name))
}

func calculate(a int, b int, operator func(x int, y int) int) {
	fmt.Println(operator(a, b))
}
