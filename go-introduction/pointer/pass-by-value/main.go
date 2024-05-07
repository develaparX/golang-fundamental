package main

import "fmt"

//contoh pass by value
func main() {
	var x = 4
	var y = x

	y = y + 1 //hanya reassign nilai y saja karena alamat memori y dan x berbeda
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	fmt.Println("alamat memori x:", &x)
	fmt.Println("alamat memori y:", &y)

	fmt.Println("========================================")

	var a = 3
	increase(a) //pass by value, memanggil func increase dengan mengassign a ke var n di func increase

	fmt.Println("a:", a)
}

func increase(n int) {
	n = n + 1
	fmt.Println("n:", n)
}
