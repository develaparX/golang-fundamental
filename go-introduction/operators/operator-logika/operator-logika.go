package main

import "fmt"

func main() {
	fmt.Println("Operator Logika")

	a := true
	b := false
	fmt.Println(a && b)                 // true bila kanan kiri true
	fmt.Println(a || b)                 // mereturn true bila salah satu true
	fmt.Println(!a)                     //meretrun kebalikan dari boolean
	fmt.Println(false || true && false) // ketika menggunakan lebih dari satu logika, yang dirun akan and(&&) dahulu
}
