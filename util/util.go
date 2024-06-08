package util

import "fmt"

const NMAX int = 10000

type paket struct {
	Jenis, PktPas string
	Harga         int
}

type pasien struct {
	Nama, Rekap string
	Waktu       tanggal
	PaketMCU    paket
	Biaya, ID   int
}

type tanggal struct {
	D, M, Y int
}

type TabPAS [NMAX]pasien
type TabPKT [NMAX]paket

func CariIdxPas(pas TabPAS, nPas int, x string, h, b, t int) int {
	found := -1
	var i int
	for i < nPas && found == -1 {
		if pas[i].Nama == x && pas[i].Waktu.D == h &&  pas[i].Waktu.M == b && pas[i].Waktu.Y == t {
			found = i
		} else {
			i++
		}
	}
	return found
}

func CariIdxPak(paket TabPKT, nPAS int, x string) int {
	found := -1
	var i int
	for i < nPAS && found == -1 {
		if paket[i].Jenis == x {
			found = i
		} else {
			i++
		}
	}
	return found
}

func OutPaket(paket TabPKT, nPAK int) {
	if nPAK == 0 {
		fmt.Println("Belum ada paket yang ditambahkan!")
	} else {
		fmt.Printf("%-5s %-30s %s\n", "No.", "Nama Paket", "Harga")
		for i := 0; i < nPAK; i++ {
			fmt.Printf("%-5d %-30s %d\n", i+1, paket[i].Jenis, paket[i].Harga)
		}
	}
}
