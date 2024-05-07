package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("z")
	}()
	fmt.Println("b")
	panic("a")
}
