package main
import (
	"fmt"
	"TugasBesar/pasienmcu"
	"TugasBesar/paketmcu"
	"TugasBesar/util"
	"TugasBesar/laporan"
)

func main() {
	var dataPkt util.TabPKT
	var dataPas util.TabPAS
	var nPkt, nPas int
	var x int

	for x != 4 {
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
			MenuPasien(&dataPas, &nPas, dataPkt, nPkt)
		} else if x == 2 {
			MenuPaket(&dataPkt, &nPkt)
		} else if x == 3 {
			laporan.Laporan(dataPas, nPas)
		} else if x == 4 {
			fmt.Print("\033[H\033[2J")
			fmt.Println("---  Selamat Tinggal  ---")
		}
	}
}

func MenuPasien(pas *util.TabPAS, nPAS *int, pak util.TabPKT, nPAK int) {
	var x int
	for x != 4 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("-------------------------")
		fmt.Println("       Menu Pasien       ")
		fmt.Println("-------------------------")
		fmt.Println("1. Tambah Data           ")
		fmt.Println("2. Cari Data             ")
		fmt.Println("3. Lihat Data            ")
		fmt.Println("4. Kembali               ")
		fmt.Println("-------------------------")
		fmt.Print("Pilihan : ")
		fmt.Scan(&x)
		if x == 1 {
			pasienmcu.InPasien(pas, nPAS, pak, nPAK)
		} else if x == 2 {
			pasienmcu.FindPasien(*pas, *nPAS)
		} else if x == 3 {
			pasienmcu.CetakPasien(*pas, *nPAS, pak, nPAK)
		}
	}
}

func MenuPaket(pak *util.TabPKT, nPAK *int) {
	var x int
	for x != 3 {
		fmt.Print("\033[H\033[2J")
		fmt.Println("-------------------------")
		fmt.Println("1. Tambah Paket          ")
		fmt.Println("2. Lihat Paket           ")
		fmt.Println("3. Kembali               ")
		fmt.Println("-------------------------")
		fmt.Print("Pilihan : ")
		fmt.Scan(&x)
		if x == 1 {
			paketmcu.InPaket(pak, nPAK)
		} else if x == 2 {
			paketmcu.CetakPaket(*pak, *nPAK)
		}
	}

}