package main

import "fmt"

type kendaraan struct { //contoh struct
	merek string
	tahun int
	model string
	harga int
}

func main() {
	var a kendaraan
	a.merek = "Toyota"
	a.tahun = 2023
	a.model = "Innova"
	a.harga = 700000000

	fmt.Println("a:", a)           // akan tercetak semua field dalam struct kendaraan
	fmt.Println("merek:", a.merek) // memanggil field dalam struct spesifik
	fmt.Println("tahun:", a.tahun)
	fmt.Println("model:", a.model)
	fmt.Println("harga:", a.harga)

	fmt.Println("=================================================")

	var b = kendaraan{} //cara pemanggilan struct kedua
	b.merek = "Wuling"
	b.tahun = 2023
	b.model = "Conferro"
	b.harga = 450000000

	fmt.Println("b:", b)

	var c = kendaraan{"Honda", 2023, "Brio", 250000000} // cara penggunaan struct ketiga, harus urut dan harus diisi semua

	fmt.Println("c:", c)

	fmt.Println("=================================================")

	// menuliskan tidak berurutan harus ditulis variablenya, tapi hasilnya akan berurutan sesuai field struct kendaraan
	var d = kendaraan{model: "Xenia", merek: "Daihastsu", harga: 300000000, tahun: 2023}
	fmt.Println("d:", d)

	fmt.Println("=================================================")

	//pass by value menggunakan struct

	var x = kendaraan{merek: "Honda", tahun: 2023, model: "HRV", harga: 350000000}
	var y kendaraan = x

	fmt.Println("Sebelum diubah =====================")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Printf("alamat x: %p\n", &x)
	fmt.Printf("alamat y: %p\n", &y)

	fmt.Println("Setelah diubah =====================")
	y.model = "CRV"

	fmt.Println("x:", x)
	fmt.Println("y:", y)

	// fmt.Println("=================================================")
	// var z = kendaraan{merek: "Honda", tahun: 2023, model: "Civic", harga: 700000000}
	// updateKendaraan(z)
	// fmt.Println("kendaraan di function main:", z)

	// fmt.Println("=================================================")

	// //pass by reference dalam struct
	// var z2 *kendaraan = &z
	// fmt.Printf("alamat z: %p \n", &z)
	// fmt.Printf("alamat z2: %p \n", z2)
	// z2.model = "CRV"
	// fmt.Println("z:", z)
	// fmt.Println("z2:", z2)

	fmt.Println("=================================================")

	//pass by reference dalam function dan struct

	var z = kendaraan{merek: "Honda", tahun: 2023, model: "Civic", harga: 700000000}
	updateKendaraan(&z) //yang tercetak akan ini karena pass by reference karena yg dikirim &z
	fmt.Println("kendaraan di function main:", &z)

	fmt.Println("=================================================")

	var k = new(kendaraan) //tidak perlu menggunakan & karena sudah menggunakan new()
	fmt.Printf("alamat k: %p \n", k)
}

func updateKendaraan(newKendaraan *kendaraan) {
	newKendaraan.merek = "Toyota"
	newKendaraan.model = "Camry"
	newKendaraan.harga = 650000000
	newKendaraan.tahun = 2023
	fmt.Println("kendaraan di function updateKendaraan:", newKendaraan)
}
