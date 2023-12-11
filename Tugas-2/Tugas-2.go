package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	//soal 1
	var word1 = "Bootcamp"
	var word2 = "Digital"
	var word3 = "Skill"
	var word4 = "Sanbercode"
	var word5 = "Golang"

	fmt.Println(word1 + " " + word2 + " " + word3 + " " + word4 + " " + word5)

	//soal 2
	halo := "Halo Dunia"
	var haloBaru = strings.Replace(halo, "Dunia", "Golang", -1)
	fmt.Println(haloBaru)

	//soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kataKedua = strings.Title(kataKedua)
	kataKetiga = kataKetiga[:len(kataKetiga)-1] + strings.ToUpper(kataKetiga[len(kataKetiga)-1:])
	kataKeempat = strings.ToUpper(kataKeempat)

	kataGabungan := kataPertama + " " + kataKedua + " " + kataKetiga + " " + kataKeempat

	// Menampilkan hasil
	fmt.Println(kataGabungan)
	
	//soal 4
	var angkaPertama= "8";
	var angkaKedua= "5";
	var angkaKetiga= "6";
	var angkaKeempat = "7";

	var pertama, _ = strconv.Atoi(angkaPertama)
	var kedua, _ = strconv.Atoi(angkaKedua)
	var Ketiga, _ = strconv.Atoi(angkaKetiga)
	var Keempat, _ = strconv.Atoi(angkaKeempat)

	var hasil = pertama + kedua + Ketiga + Keempat

	fmt.Println(hasil)

	//soal 5
	var kalimat = "halo halo bandung"
	var angka = 2021

	// Mengubah "halo" menjadi "Hi" dengan fungsi strings.Replace
	kalimat = strings.Replace(kalimat, "halo", "Hi", -1)

	// Menggabungkan kalimat dan angka dengan tanda "-" di antara
	var result = fmt.Sprintf("%s - %d", kalimat, angka)

	// Menampilkan result
	fmt.Println(result)

}