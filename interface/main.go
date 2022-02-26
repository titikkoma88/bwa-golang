package main

import "fmt"

type Luas interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {
	// bangunDatar := Persegi{Sisi: 4}
	bangunDatar := PersegiPanjang{Panjang: 6, Lebar: 5}

	luas := SeberapaLuas(bangunDatar)
	fmt.Println("Luas : ", luas)
}

func SeberapaLuas(bangunDatar PersegiPanjang) int {
	return bangunDatar.HitungLuas()
}

func SeberapaLuas2(bangunDatar Persegi) int {
	return bangunDatar.HitungLuas()
}
