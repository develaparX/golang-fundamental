package main

import "fmt"

func main() {
	var firstName string = "Kenjiro"

	var lastName string = "Kenzata"

	fmt.Println(firstName, lastName)
	fmt.Printf("Halo, Kejiro Kenzata\n")
	fmt.Printf("Halo, %s %s\n", firstName, lastName)

	var (
		age     int
		address string
	)

	age = 25
	address = "Jakarta"

	fmt.Printf("Halo, saya %s %s berumur %d dan saya tinggal di %s \n", firstName, lastName, age, address)

	var (
		bootcampName, bootcampAddress = "Enigma Camp", "Jakarta Selatan"
	)
	fmt.Println(bootcampName, bootcampAddress)

	day := "Monday"
	date := "31"
	month := "October"

	fmt.Println(day, date, month+" 2022")

	var number = 21
	number = 20

	const phi = 3.14

	fmt.Println(number, phi)

	bootcamp, _ := "Enigma Camp", "Aktif7"
	fmt.Println(bootcamp)

}
