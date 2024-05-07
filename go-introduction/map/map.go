package main

import "fmt"

func main() {
	user := map[string]string{
		"username": "kenjif",
		"email":    "kenji@mail.com",
	}
	fmt.Println(user)

	fmt.Println("======================================================")

	var scores = make(map[string]int)
	fmt.Println(scores)

	scores["java"] = 85
	scores["react"] = 87
	scores["kotlin"] = 90

	fmt.Println("Scores :", scores)

	fmt.Println("Nilai Java :", scores["java"]) // mencetak scores yang keynya java
	fmt.Println("Nilai Kotlin :", scores["kotlin"])
	fmt.Println("Nilai Golang :", scores["golang"]) // mencetak yang keynya gaada, output akan ada default value yaitu 0

	fmt.Println("======================================================")

	scores["java"] = 90 // reassign value
	fmt.Println("Scores :", scores)

	fmt.Println("======================================================")

	// menghapus value di map
	delete(scores, "kotlin")
	fmt.Println("Scores :", scores)

	fmt.Println("======================================================")

	names := map[int]string{
		1: "john",
		2: "jane",
		3: "lili",
		4: "ruby",
	}

	for _, value := range names {
		fmt.Println("Value :", value)
	}

}
