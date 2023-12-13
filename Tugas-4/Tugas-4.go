package main

import (
	"fmt"
)

func main() {
	//soal 1
	jumlahPerulangan := 20

	for i := 1; i <= jumlahPerulangan; i++ {

		if i%2 == 0 {
			fmt.Println(i, "- Berkualitas")
		} else if i%3 == 0 {
			fmt.Println(i, "- I Love Coding")
		} else {
			fmt.Println(i, "- Santai")
		}
	}

	//soal2 
	tinggi := 7

	for i := 1; i <= tinggi; i++ {
		for j := 1; j <= i ; j++ {
			fmt.Print("#")
		}
	fmt.Println()
	}

	//soal 3
	kalimat := [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	potonganKalimat := kalimat[2:7]

	fmt.Println(potonganKalimat)

	//soal 4
	var sayuran = []string{}

	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for i, s := range sayuran {
		fmt.Printf("%d. %s\n", i+1, s)
	}

	//soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	// Menampilkan nilai variabel dengan looping
	for key, value := range satuan {
		fmt.Printf("%s = %d\n", key, value)
	}

	// Menghitung volume balok
	volumeBalok := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]

	// Menampilkan volume balok
	fmt.Printf("volume balok = %d\n", volumeBalok)
}