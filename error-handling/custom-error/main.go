package main

import "fmt"

type Item struct {
	id    int
	name  string
	price int
}

type ItemNotFoundError struct { //custom-error harus punya struct
	Id int
}

func (i *ItemNotFoundError) Error() string { // kemudian custom error harus seperti ini
	return fmt.Sprintf("Item dengan id: %d tidak ada", i.Id)
}

var items = []Item{
	{1, "Laptop", 20000},
	{2, "Jam", 5000},
	{3, "Hp", 10000},
}

func getItemById(id int) (Item, error) { //akan error jika mencari id yang tidak ada di item
	for _, item := range items {
		if item.id == id {
			return item, nil
		}
	}
	return Item{}, &ItemNotFoundError{Id: id}
}

func main() {
	result, err := getItemById(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
