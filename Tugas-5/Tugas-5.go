package main

import (
	"fmt"
	"strings"
)

func luasPersegiPanjang(panjang, lebar int)int{
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar int) int{
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang, lebar, tinggi int) int{
	return panjang * lebar * tinggi
}

func introduce(nama, jenisKelamin, pekerjaan, usia string) string {
	var panggilan string

	if jenisKelamin == "laki-laki" {
		panggilan = "Pak"
	} else if jenisKelamin == "perempuan" {
		panggilan = "Bu"
	} else {
		panggilan = ""
	}

	result := fmt.Sprintf("%s %s adalah seorang %s yang berusia %s tahun", panggilan, nama, pekerjaan, usia)

	return result
}

func buahFavorit(nama string, buah ...string) string {
	buahStr := strings.Join(buah, ", ")
	return fmt.Sprintf("Halo, nama saya %s dan buah favorit saya adalah %s", nama, buahStr)
}

var dataFilm = []map[string]string{}

func tambahDataFilm() func(string, string, string, string) {
	return func(title, jam, genre, tahun string) {
		film := map[string]string{
			"title":  title,
			"jam": jam,
			"genre":  genre,
			"tahun":  tahun,
		}
		dataFilm = append(dataFilm, film)
	}
}


func main(){
	
	//soal 1
	panjang := 12
	lebar := 4
	tinggi := 8
	
	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas) 
	fmt.Println(keliling)
	fmt.Println(volume)

	//soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) 
	
	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	//soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	//soal 4
	tambahDataFilm := tambahDataFilm()

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("Avenger", "2 jam", "action", "2019")
	tambahDataFilm("Spiderman", "2 jam", "action", "2004")
	tambahDataFilm("Juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
