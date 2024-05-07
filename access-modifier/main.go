package main

import (
	"access-modifier/library" //mengimport package di semua folder/package library
	"fmt"
)

func main() {
	fmt.Println("HourInADay:", library.HourInADay) //mengakses variable HourInADay didalam library.go
	// fmt.Println("secondsInMinutes:", library.secondsInMinutes) // akan error karena variable unExported

	fmt.Println("Name:", library.Name)

	fmt.Println(library.DaysToMinutes(3))
}
