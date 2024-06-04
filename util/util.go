package util

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
