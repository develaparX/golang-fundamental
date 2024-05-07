package main

import "fmt"

type mesin struct { //embedded struct
	tipe string
	cc   int
}

type kendaraan struct { //contoh struct
	merek string
	tahun int
	model string
	harga int
	mesin //embedded struct
}

func main() {
	var a = kendaraan{
		merek: "Toyota",
		tahun: 2023,
		model: "Camry",
		harga: 450000000,
		mesin: mesin{
			tipe: "premium",
			cc:   2000,
		},
	}
	fmt.Println("a:", a)
	fmt.Println("a mesin.tipe:", a.mesin.tipe)

}

func updateKendaraan(newKendaraan *kendaraan) {
	newKendaraan.merek = "Toyota"
	newKendaraan.model = "Camry"
	newKendaraan.harga = 650000000
	newKendaraan.tahun = 2023
	fmt.Println("kendaraan di function updateKendaraan:", newKendaraan)
}
