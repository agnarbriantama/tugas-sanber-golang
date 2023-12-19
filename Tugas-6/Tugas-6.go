package main

import (
	"fmt"
	"math"
)

var luasLingkaran float64
var kelilingLingkaran float64

func hitungLuasKelilingLingkaran(jariJari *float64) {
	luasLingkaran = math.Pi * math.Pow(*jariJari, 2)
	kelilingLingkaran = 2 * math.Pi * (*jariJari)
}

func introduce(sentence *string, name, gender, occupation, age string) {
	if gender == "laki-laki" {
		*sentence = fmt.Sprintf("Pak %s adalah seorang %s yang berusia %s tahun", name, occupation, age)
	} else if gender == "perempuan" {
		*sentence = fmt.Sprintf("Bu %s adalah seorang %s yang berusia %s tahun", name, occupation, age)
	} else {
		*sentence = "Gender tidak valid"
	}
}

func tambahDataFilm(title, duration, genre, releaseYear string, dataFilm *[]map[string]string) {
	film := map[string]string{
		"title":       title,
		"duration":    duration,
		"genre":       genre,
		"releaseYear": releaseYear,
	}
	*dataFilm = append(*dataFilm, film)
}

func tambahBuah(buah *[]string, namaBuah string) {
	*buah = append(*buah, namaBuah)
}

func main() {
	//soal 1
	jariJari := 5.0
	fmt.Printf("Jari-jari lingkaran: %.2f\n", jariJari)
	
	hitungLuasKelilingLingkaran(&jariJari)

	fmt.Printf("Luas Lingkaran: %.2f\n", luasLingkaran)
	fmt.Printf("Keliling Lingkaran: %.2f\n", kelilingLingkaran)

	//soal 2
	var sentence string

	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

	//soal 3
	var buah = []string{}

	tambahBuah(&buah, "Jeruk")
	tambahBuah(&buah, "Semangka")
	tambahBuah(&buah, "Mangga")
	tambahBuah(&buah, "Strawberry")
	tambahBuah(&buah, "Durian")
	tambahBuah(&buah, "Manggis")
	tambahBuah(&buah, "Alpukat")

	for i, namaBuah := range buah {
		fmt.Printf("%d. %s\n", i+1, namaBuah)
	}

	//soal 4
	var dataFilm = []map[string]string{}

	// Menambahkan data film menggunakan fungsi tambahDataFilm
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	// Menampilkan hasil
	for i, film := range dataFilm {
		fmt.Printf("%d. Title: %s\n", i+1, film["title"])
		fmt.Printf("Duration: %s\n", film["duration"])
		fmt.Printf("Genre: %s\n", film["genre"])
		fmt.Printf("Release Year: %s\n\n", film["releaseYear"])
	}
}