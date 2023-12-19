package main

import "fmt"

type Buah struct {
	Nama       string
	Warna      string
	AdaBijinya bool
	Harga      int
}

type segitiga struct{
	alas, tinggi int
  }
  
  type persegi struct{
	sisi int
  }
  
  type persegiPanjang struct{
	panjang, lebar int
  }

  func (s segitiga) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (p persegi) luas() int {
	return p.sisi * p.sisi
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

// Method untuk menambahkan warna ke property colors pada phone
func (p *phone) addColor(newColor string) {
	p.colors = append(p.colors, newColor)
}

type movie struct {
	title    string
	genre    string
	duration int
	year     int
}

var dataFilm = []movie{}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	filmBaru := movie{
		title:    title,
		genre:    genre,
		duration: duration,
		year:     year,
	}
	*dataFilm = append(*dataFilm, filmBaru)
}

func main() {
	//soal 1
	nanas := Buah{"Nanas", "Kuning", false, 9000}
	jeruk := Buah{"Jeruk", "Oranye", true, 8000}
	semangka := Buah{"Semangka", "Hijau & Merah", true, 10000}
	pisang := Buah{"Pisang", "Kuning", false, 5000}

	tampilkanBuah(nanas)
	tampilkanBuah(jeruk)
	tampilkanBuah(semangka)
	tampilkanBuah(pisang)

	//soal 2
	segitiga1 := segitiga{alas: 5, tinggi: 8}
	persegi1 := persegi{sisi: 4}
	persegiPanjang1 := persegiPanjang{panjang: 6, lebar: 10}

	luasSegitiga := segitiga1.luas()
	luasPersegi := persegi1.luas()
	luasPersegiPanjang := persegiPanjang1.luas()

	fmt.Printf("Luas Segitiga: %d\n", luasSegitiga)
	fmt.Printf("Luas Persegi: %d\n", luasPersegi)
	fmt.Printf("Luas Persegi Panjang: %d\n", luasPersegiPanjang)

	//soal 3
	myPhone := phone{
		name:   "Galaxy A35",
		brand:  "Samsung",
		year:   2022,
		colors: []string{"Black", "Silver"},
	}

	fmt.Println("Warna sebelum ditambahkan:", myPhone.colors)
	myPhone.addColor("Blue")
	fmt.Println("Warna setelah ditambahkan:", myPhone.colors)

	//soal 4
	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. Title: %s\n", i+1,film.title)
		fmt.Printf("Genre: %s\n", film.genre)
		fmt.Printf("Duration: %d minutes\n", film.duration)
		fmt.Printf("Year: %d\n\n", film.year)
	}

}

func tampilkanBuah(buah Buah) {
	fmt.Printf("Nama: %s\n", buah.Nama)
	fmt.Printf("Warna: %s\n", buah.Warna)
	fmt.Printf("Ada Bijinya: %t\n", buah.AdaBijinya)
	fmt.Printf("Harga: %d\n\n", buah.Harga)
}