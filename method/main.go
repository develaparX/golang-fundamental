package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

// membuat struct triangle dengan method area
func (t triangle) area() float64 { //value receiver : area()
	// fmt.Println("Luas:", 0.5*t.base*t.height) // jika ingin langsung cetak tanpa inisiasi di #2
	return 0.5 * t.base * t.height // cetak dengan return float64 dan harus inisiasi di #2 atau contohnya di #3
}

//pointer receiver
func (t *triangle) increaseSize() { //harus (*)pointer receiver untuk hasil yang berubah di f main atau di hasil #4, bisa lihat #4
	t.base += 1
	t.height += 1
	fmt.Println("base:", t.base)
	fmt.Println("height:", t.height)
}

func main() {
	instanceTriangle := triangle{10, 12} // ini siasi value di triangle #1
	// instanceTriangle.area()              //memanggil method area #2
	area := instanceTriangle.area() // var area baru untuk nilai return di method #3
	fmt.Println("area:", area)

	fmt.Println("=========================================================================")

	fmt.Println("instanceTriangle :", instanceTriangle)
	instanceTriangle.increaseSize() // inisiasi atau pemanggilan method #4
	fmt.Println("instanceTriangle :", instanceTriangle)

}
