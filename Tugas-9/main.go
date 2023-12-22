package main

import (
	"fmt"
	tugas8 "tugas-9-module/Tugas-8"
)

func main() {
	// soal 1
	segitiga := tugas8.SegitigaSamaSisi{Alas: 4, Tinggi: 3}
	persegi := tugas8.PersegiPanjang{Panjang: 5, Lebar: 6}
	silinder := tugas8.Tabung{JariJari: 2.5, Tinggi: 4.0}
	kubus := tugas8.Balok{Panjang: 3, Lebar: 3, Tinggi: 3}

	fmt.Println("Luas dan Keliling Segitiga:")
	tugas8.Hasil(segitiga)

	fmt.Println("\nLuas dan Keliling Persegi Panjang:")
	tugas8.Hasil(persegi)

	fmt.Println("\nVolume dan Luas Permukaan Silinder:")
	tugas8.Hasil(silinder)

	fmt.Println("\nVolume dan Luas Permukaan Kubus:")
	tugas8.Hasil(kubus)

	// soal 2
	myPhone := tugas8.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	result := tugas8.Display(myPhone)
	fmt.Println(result)

	// soal 3
	fmt.Println(tugas8.LuasPersegi(4, true))
	fmt.Println(tugas8.LuasPersegi(8, false))
	fmt.Println(tugas8.LuasPersegi(0, true))
	fmt.Println(tugas8.LuasPersegi(0, false))

	// soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	angkaPertama, ok1 := kumpulanAngkaPertama.([]int)
	angkaKedua, ok2 := kumpulanAngkaKedua.([]int)

	if !ok1 || !ok2 {
		fmt.Println("Type assertion gagal.")
		return
	}

	total := 0
	for _, angka := range angkaPertama {
		total += angka
	}
	for _, angka := range angkaKedua {
		total += angka
	}

	prefixStr, ok := prefix.(string)
	if !ok {
		fmt.Println("Type assertion untuk prefix gagal.")
		return
	}

	fmt.Printf("%s%d = %d\n", prefixStr, total, total)
}
