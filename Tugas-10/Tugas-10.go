package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"sort"
	"time"
)

func kelilingSegitigaSamaSisi(sisi int, tampilkanKalimat bool) (string, error) {
	if sisi == 0 {
		if tampilkanKalimat {
			return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
		return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}

	keliling := 3 * sisi

	if tampilkanKalimat {
		return fmt.Sprintf("keliling segitiga sama sisinya dengan sisi %v cm adalah %v cm", sisi, keliling), nil
	}
	return fmt.Sprintf("%v", keliling), nil
}

func tambahAngka(	nilai int, total *int){
	*total += nilai
}

func cetakAngka(total *int){
	fmt.Println("Total angka:", *total)
}

var phones = []string{}

func addPhone(brand string, phoneList *[]string) {
	*phoneList = append(*phoneList, brand)
}

func printSortedPhones(phoneList []string) {
	sort.Strings(phoneList)
	for i, phone := range phoneList {
		fmt.Printf("%v. %v\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}

func hitungLuas(jariJari float64) int {
	luas := math.Pi * math.Pow(jariJari, 2)
	return int(math.Round(luas))
}

func hitungKeliling(jariJari float64) int {
	keliling := 2 * math.Pi * jariJari
	return int(math.Round(keliling))
}

func hitungLuasPanjang(panjang, lebar float64) float64 {
	return panjang * lebar
}

func hitungKelilingPanjang(panjang, lebar float64) float64 {
	return 2 * (panjang + lebar)
}

func main() {
	//soal 1
	kalimat := "Golang Backend Development"
	tahun := 2021

	printKalimatDanTahun(kalimat, tahun)

	//soal 2
	result1, err1 := kelilingSegitigaSamaSisi(4, true)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println(result1)
	}

	result2, err2 := kelilingSegitigaSamaSisi(8, false)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println(result2)
	}

	result3, err3 := kelilingSegitigaSamaSisi(0, true)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println(result3)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	result4, err4 := kelilingSegitigaSamaSisi(0, false)
	if err4 != nil {
		panic(err4)
	} else {
		fmt.Println(result4)
	}

	//soal 3
	// deklarasi variabel angka ini simpan di baris pertama func main
	angka := 1

	defer cetakAngka(&angka)

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	//soal 4
	// Menambahkan data ke variabel phones 
	addPhone("Xiaomi", &phones)
	addPhone("Asus", &phones)
	addPhone("IPhone", &phones)
	addPhone("Samsung", &phones)
	addPhone("Oppo", &phones)
	addPhone("Realme", &phones)
	addPhone("Vivo", &phones)

	// Mengurutkan dan menampilkan data satu persatu setiap detik
	printSortedPhones(phones)

	//soal 5
	jariJari1 := 7.0
	jariJari2 := 10.0
	jariJari3 := 15.0

	luas1 := hitungLuas(jariJari1)
	keliling1 := hitungKeliling(jariJari1)

	luas2 := hitungLuas(jariJari2)
	keliling2 := hitungKeliling(jariJari2)

	luas3 := hitungLuas(jariJari3)
	keliling3 := hitungKeliling(jariJari3)

	fmt.Printf("Jari-jari: %.0f\n", jariJari1)
	fmt.Printf("Luas Lingkaran: %d\n", luas1)
	fmt.Printf("Keliling Lingkaran: %d\n\n", keliling1)

	fmt.Printf("Jari-jari: %.0f\n", jariJari2)
	fmt.Printf("Luas Lingkaran: %d\n", luas2)
	fmt.Printf("Keliling Lingkaran: %d\n\n", keliling2)

	fmt.Printf("Jari-jari: %.0f\n", jariJari3)
	fmt.Printf("Luas Lingkaran: %d\n", luas3)
	fmt.Printf("Keliling Lingkaran: %d\n", keliling3)

	//soal 6
	var panjang, lebar float64

	// Menggunakan package flag untuk mendapatkan input panjang dan lebar dari command line
	flag.Float64Var(&panjang, "panjang", 0, "Panjang persegi panjang")
	flag.Float64Var(&lebar, "lebar", 0, "Lebar persegi panjang")
	flag.Parse()

	// Memastikan bahwa panjang dan lebar positif
	if panjang <= 0 || lebar <= 0 {
		fmt.Println("Panjang dan lebar harus lebih dari 0.")
		return
	}

	luas := hitungLuasPanjang(panjang, lebar)
	keliling := hitungKelilingPanjang(panjang, lebar)

	fmt.Printf("Panjang: %.2f\n", panjang)
	fmt.Printf("Lebar: %.2f\n", lebar)
	fmt.Printf("Luas Persegi Panjang: %.2f\n", luas)
	fmt.Printf("Keliling Persegi Panjang: %.2f\n", keliling)

}

func printKalimatDanTahun(kalimat string, tahun int) {
	defer func() {
		fmt.Printf("Eksekusi terakhir: Kalimat: %v, Tahun: %v\n", kalimat, tahun)
	}()

	// Logic atau pemrosesan lainnya
	fmt.Printf("Eksekusi awal: Kalimat: %v, Tahun: %v\n", kalimat, tahun)
	
}
