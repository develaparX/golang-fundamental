package main

import "fmt"

type Patient struct {
	Name string
	Age  int
	Celcius
}

type Celcius float64 //custom type

func (c Celcius) toFahrenheit() float64 {
	return float64(c*9/5 + 32) // untuk mereturn float64, harus diinisiasi ulang hasilnya menggunakan float64
}

func (c *Celcius) tambah(value float64) {
	*c += Celcius(value) //mengubah value yang bertype float64 menjadi bertipe celcius
}

func main() {
	var temperature Celcius = 20.0
	fmt.Println("temperature:", temperature)
	fmt.Println("Suhu di ruangan ini", temperature.toFahrenheit(), "derajat Fahrenheit")
	temperature.tambah(3)
	fmt.Println("Suhu di ruangan ini menjadi", temperature, "derajat Celcius")

	fmt.Println("====================================================================")

	newPatient := Patient{Name: "Kenji", Age: 25, Celcius: 39.0}
	fmt.Printf("newPatient:%+v \n", newPatient)
}
