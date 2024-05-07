package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ { //pengulangan akan diulang sampai i<=5 jadi false oleh i++
		fmt.Println("Angka", i)
	}

	fmt.Println("==========================================")

	var i = 6     // definisikan nilai i
	for i <= 10 { // lopping dengan kondisi ini, jika true akan di eksekusi sampai <=10
		fmt.Println("Angka", i)
		i++ // tanpa ini akan ada infinite loop
	}

	fmt.Println("==========================================")

	var index = 0 //tanpa ini akan infinite looping
	for {
		fmt.Println("Enigma")
		index++         // tanpa ini dan diatas juga akan infinite looping
		if index == 5 { // ini akan dijalankan sampai bernilai true yaitu dalam hal ini 5 kali
			break
		}
	}

	fmt.Println("==========================================")

	//nested looping
	for i := 1; i <= 3; i++ { // looping paling luar untuk menentukan berapa barisnya
		for j := 1; j <= 5; j++ { // looping dalam menentukan jumlah data tiap barisnya atau mencetak 1 sampai 5
			fmt.Print(j, " ")
		}
		fmt.Println() // berpindah baris baru, sampai nilai i <= 3
	}

	fmt.Println("==========================================")

	//perbedaan break dan continue
	// break untuk keluar dari looping
	// continue untuk lanjut ke perulangan selanjutnya tanpa mengeksekusi kode didalam looping

	// contoh break
	fmt.Println("Start")

	for i := 1; i <= 10; i++ {
		if i == 5 {
			break // jika i loop bernilai 5, if break ini akan bernilai True dan program keluar dari line if ini dan eksekusi line 55
		}
		fmt.Print(i, " ")
	}

	fmt.Println()
	fmt.Println("Done")

	fmt.Println("==========================================")

	//contoh continua

	fmt.Println("Start")

	for i := 1; i <= 10; i++ { // perulangan akan berlanjut sampai 10
		if i == 5 { // jika nilai ini bernilai true, maka tidak dijalankan/dicetak dan kemudian lanjut sampai 10
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	fmt.Println("Done")

	fmt.Println("==========================================")

	//break dalam nested loop
	fmt.Println("Start")

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 5; j++ {
			if j == 4 {
				// break //akan berhenti di 4
				continue // angka 4 akan dilewati
			}
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	fmt.Println("Done")
	//looping berlaku dalam statement yang mengurung hal tersebut, tidak diluar
}
