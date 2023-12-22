package tugas8

import (
	"fmt"
	"math"
	"strings"
)

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (s SegitigaSamaSisi) Luas() int {
	return (s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return 3 * s.Alas
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * math.Pi * t.JariJari * (t.JariJari + t.Tinggi)
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * (float64(b.Panjang*b.Lebar) + float64(b.Panjang*b.Tinggi) + float64(b.Lebar*b.Tinggi))
}

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

type DisplayData interface {
	Display() string
}

func (p Phone) Display() string {
	colorsStr := strings.Join(p.Colors, ", ")

	return fmt.Sprintf("name:   %q,\nbrand:  %q,\nyear:   %d,\ncolors: %s", p.Name, p.Brand, p.Year, colorsStr)
}

func LuasPersegi(sisi int, tampilkanKalimat bool) interface{} {
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

func Hasil(bangun interface{}) {
	switch v := bangun.(type) {
	case HitungBangunDatar:
		fmt.Printf("Luas: %d\n", v.Luas())
		fmt.Printf("Keliling: %d\n", v.Keliling())
	case HitungBangunRuang:
		fmt.Printf("Volume: %.2f\n", v.Volume())
		fmt.Printf("Luas Permukaan: %.2f\n", v.LuasPermukaan())
	}
}

func Display(d DisplayData) string {
	return d.Display()
}
