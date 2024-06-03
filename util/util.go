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

