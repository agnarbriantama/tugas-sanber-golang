package main

import (
	"fmt"
	"math"
	"strings"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return (s.alas * s.tinggi) / 2
}

func (s segitigaSamaSisi) keliling() int {
	return 3 * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}


func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * (float64(b.panjang*b.lebar) + float64(b.panjang*b.tinggi) + float64(b.lebar*b.tinggi))
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type displayData interface {
	display() string
}

func (p phone) display() string {
	colorsStr := strings.Join(p.colors, ", ")

	return fmt.Sprintf("name:   %q,\nbrand:  %q,\nyear:   %d,\ncolors: %s", p.name, p.brand, p.year, colorsStr)
}

func luasPersegi(sisi int, tampilkanKalimat bool) interface{} {
	if sisi == 0 {
		if tampilkanKalimat {
			return "Maaf anda belum menginput sisi dari persegi"
		}
		return nil
	}

	luas := sisi * sisi

	if tampilkanKalimat {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	}
	return luas
}


func main() {
	//soal 1
	segitiga := segitigaSamaSisi{alas: 4, tinggi: 3}
	persegi := persegiPanjang{panjang: 5, lebar: 6}
	silinder := tabung{jariJari: 2.5, tinggi: 4.0}
	kubus := balok{panjang: 3, lebar: 3, tinggi: 3}

	fmt.Println("Luas dan Keliling Segitiga:")
	hasil(segitiga)

	fmt.Println("\nLuas dan Keliling Persegi Panjang:")
	hasil(persegi)

	fmt.Println("\nVolume dan Luas Permukaan Silinder:")
	hasil(silinder)

	fmt.Println("\nVolume dan Luas Permukaan Kubus:")
	hasil(kubus)

	//soal 2
	myPhone := phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}


	result := display(myPhone)
	fmt.Println(result)

	//soal 3
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	//soal 4
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

func hasil(bangun interface{}) {
	switch v := bangun.(type) {
	case hitungBangunDatar:
		fmt.Printf("Luas: %d\n", v.luas())
		fmt.Printf("Keliling: %d\n", v.keliling())
	case hitungBangunRuang:
		fmt.Printf("Volume: %.2f\n", v.volume())
		fmt.Printf("Luas Permukaan: %.2f\n", v.luasPermukaan())
	}
}

func display(d displayData) string {
	return d.display()
}