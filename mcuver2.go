package main

import "fmt"

const NMAX int = 10000

type paket struct {
	jenis, PktPas string
	harga         int
}

type pasien struct {
	nama, rekap string
	waktu       tanggal
	paketMCU    paket
	biaya, ID   int
}

type tanggal struct {
	D, M, Y int
}

type tabPAS [NMAX]pasien
type tabPKT [NMAX]paket

func main() {
	var dataPas tabPAS
	var dataPkt tabPKT
	var nPas, nPkt int
	var x int

	for  {
		fmt.Println("-------- WELCOME --------")
		fmt.Println("     Medical Checkup     ")
		fmt.Println("        Main Menu        ")
		fmt.Println("-------------------------")
		fmt.Println("1. Data Pasien           ")
		fmt.Println("2. Paket Layanan         ")
		fmt.Println("3. Laporan Pemasukan     ")
		fmt.Println("4. Keluar                ")
		fmt.Println("-------------------------")
		fmt.Print("Pilihan: ")
		fmt.Scan(&x)
		if x == 1 {
			menu_pasien(&dataPas, &nPas, &dataPkt, &nPkt)
		} else if x == 2 {
			menu_paket()
		} else if x == 3 {
			laporan(data, nData)
		} else if x == 4 {
			fmt.Println("---  Selamat Tinggal  ---")
		}
	}
	
}

func menu_pasien(pas *tabPAS, nPAS *int, pak *tabPKT, nPAK *int) {
	var x int

	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("-------------------------")
		fmt.Println("       Menu Pasien       ")
		fmt.Println("-------------------------")
		fmt.Println("1. Tambah Data Pasien    ")
		fmt.Println("2. Cari Data Pasien      ")
		fmt.Println("3. Lihat Data Pasien     ")
		fmt.Println("4. Kembali               ")
		fmt.Println("-------------------------")
		fmt.Print("Pilihan : ")
		fmt.Scan(&x)
		if x == 1 {
			inPasien(&pas, &nPAS, *pak, *nPAK)
		} else if x == 2 {
			findPasien(data, nData)
		} else if x == 3 {
			tampilkan_Pasien(data, nData)
		} else if x == 4 {
			main()
		}
	}
}
