package main

import ("fmt"
		"strconv"
	)

func main() {

	//soal1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	// ubah var string menjadi integer
	panjang, err1 := strconv.Atoi(panjangPersegiPanjang)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	lebar, err2 := strconv.Atoi(lebarPersegiPanjang)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	alas, err3 := strconv.Atoi(alasSegitiga)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	tinggi, err4 := strconv.Atoi(tinggiSegitiga)
	if err4 != nil {
		fmt.Println(err4)
		return
	}

	// Deklarasi variabel integer
	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int
	// luas dan keliling persegi panjang
	luasPersegiPanjang = panjang * lebar
	kelilingPersegiPanjang = 2 * (panjang + lebar)
	// luas segitiga
	luasSegitiga = alas * tinggi / 2

	fmt.Println("Luas persegi panjang =", luasPersegiPanjang)
	fmt.Println("Keliling persegi panjang =", kelilingPersegiPanjang)
	fmt.Println("Luas segitiga =", luasSegitiga)

	//soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	var indeksNilaiJohn string
	var indeksNilaiDoe string

	if nilaiJohn >= 80{
		indeksNilaiJohn = "indeksnya A"
	}else if nilaiJohn >= 70{
		indeksNilaiJohn = "indeksnya B"
	}else if nilaiJohn >= 60 {
		indeksNilaiJohn = "indeksnya C"
	}else if nilaiJohn >= 50{
		indeksNilaiJohn = "indeksnya D"
	} else{
		indeksNilaiJohn= "indeksnya E"
	}

	
	if nilaiDoe >= 80{
		indeksNilaiDoe = "indeksnya A"
	}else if nilaiDoe >= 70{
		indeksNilaiDoe = "indeksnya B"
	}else if nilaiDoe >= 60 {
		indeksNilaiDoe = "indeksnya C"
	}else if nilaiDoe >= 50{
		indeksNilaiDoe = "indeksnya D"
	} else{
		indeksNilaiDoe= "indeksnya E"
	}

	fmt.Println("Nilai John " + indeksNilaiJohn)
	fmt.Println("Nilai Doe " + indeksNilaiDoe)
	
	//soal 3 

	var tanggal = 12
	var bulan = 7
	var tahun = 2001

	output := fmt.Sprintf("%d ", tanggal)

	switch bulan {
	case 1:
		output += "Januari"
	case 2:
		output += "Februari"
	case 3:
		output += "Maret"
	case 4:
		output += "April"
	case 5:
		output += "Mei"
	case 6:
		output += "Juni"
	case 7:
		output += "Juli"
	case 8:
		output += "Agustus"
	case 9:
		output += "September"
	case 10:
		output += "Oktober"
	case 11:
		output += "November"
	case 12:
		output += "Desember"
	default:
		output += "Bulan tidak valid"
	}

	output += fmt.Sprintf(" %d", tahun)

	fmt.Println(output)

	//soal 4
	tahunKelahiran := 2001

	if tahunKelahiran >= 1944 && tahunKelahiran <= 1964 {
		fmt.Println("Anda termasuk dalam generasi Baby Boomer.")
	} else if tahunKelahiran >= 1965 && tahunKelahiran <= 1979 {
		fmt.Println("Anda termasuk dalam generasi X.")
	} else if tahunKelahiran >= 1980 && tahunKelahiran <= 1994 {
		fmt.Println("Anda termasuk dalam generasi Y (Millennials).")
	} else if tahunKelahiran >= 1995 && tahunKelahiran <= 2015 {
		fmt.Println("Anda termasuk dalam generasi Z.")
	} else {
		fmt.Println("Tahun kelahiran Anda di luar rentang yang didefinisikan.")
	}


}