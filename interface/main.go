package main

import "fmt"

type Rectangle struct {
	width  float64
	length float64
}
type Shape interface { //interface, kumpulan method tanpa implementasi
	getArea() float64
	getPerimeter() float64
}

func (r Rectangle) getArea() float64 {
	return r.width * r.length
}

func (r Rectangle) getPerimeter() float64 {
	return 2 * (r.width + r.length)
}

type Square struct {
	side float64
}

func (s Square) getArea() float64 {
	return s.side * s.side
}
func (s Square) getPerimeter() float64 {
	return 4 * s.side
}

// func getAreaOfRectangle(r Rectangle) {
// 	fmt.Println("Luas:", r.getArea())
// }
// func getAreaOfSquare(s Square) {
// 	fmt.Println("Luas:", s.getArea())
// }

func getArea(s Shape) { //menggantikan dua func diatas
	fmt.Println("Luas:", s.getArea())
}

func main() {
	var shape1 Shape = Square{side: 5}
	var shape2 Shape = Rectangle{width: 4, length: 3}
	var shape3 Shape = Rectangle{width: 7, length: 9}
	// ini ke atas upcasting

	fmt.Printf("shape1: %#v \n", shape1)
	fmt.Printf("shape2: %#v \n", shape2)
	fmt.Printf("shape3: %#v \n", shape3)

	fmt.Println("======================================================================")

	shapes := []Shape{shape1, shape2, shape3}

	for _, shape := range shapes {
		fmt.Printf("%#v memiliki luas %.1f dan keliling %.1f \n", shape, shape.getArea(), shape.getPerimeter())
	}

	fmt.Println("======================================================================")

	// getAreaOfRectangle(Rectangle{width: 3, length: 7})
	// getAreaOfSquare(Square{side: 9})

	getArea(Rectangle{width: 2, length: 3})
	getArea(Square{side: 7})

	fmt.Println("======================================================================")

	//downcasting

	var square1 Shape = Square{side: 6}
	fmt.Println("getArea:", square1.getArea())
	fmt.Println("getPerimeter:", square1.getPerimeter())
	// fmt.Println("side:",square1.side) //tidak akan bisa mendapatkan value side

	var originalSquare Square = square1.(Square) //downcasting interface agar bisa mengakses semuanya termasuk side
	fmt.Println("getArea:", originalSquare.getArea())
	fmt.Println("getPerimeter:", originalSquare.getPerimeter())
	fmt.Println("side:", originalSquare.side)

	fmt.Println("======================================================================")
	//interface kosong, bisa menampung berbagai jenis data atau biasa disebut dynamic typing

	// var anything interface{} //sama dengan dibawah
	var anything any //bentuk lain interfdace{}di golang diatas 1.18

	anything = "Kenji" //string
	fmt.Printf("anything %T: %v \n", anything, anything)

	anything = 20 // re assign nilai yang tadinya string jadi int, bisa dilihat di hasil run untuk bedanya
	fmt.Printf("anything %T: %v \n", anything, anything)

	anything = false //menjadi boolean
	fmt.Printf("anything %T: %v \n", anything, anything)

	anything = []string{"Golang", "Java", "Python"} // slice
	fmt.Printf("anything %T: %v \n", anything, anything)

}
