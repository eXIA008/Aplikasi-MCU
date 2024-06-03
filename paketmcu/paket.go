package paketmcu

import (
	"TugasBesar/util"
	"fmt"
)

func MenuPaket(paket *util.TabPKT, nPAK *int) {
	var x int
	var dataPkt util.TabPKT
	var nPkt int
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
			inPaket(&dataPkt, &nPkt)
		} else if x == 2 {
			cetakPaket(dataPkt, nPkt)
		}
	}

}

func inPaket(paket *util.TabPKT, nPAK *int) {
	var pilih string
	for pilih != "N" {
		fmt.Print("\033[H\033[2J")
		fmt.Println("-------------------------")
		fmt.Println("    Tambah Data Paket    ")
		fmt.Println("-------------------------")
		fmt.Print("Nama Paket (Spasi diganti dengan '_') : ")
		fmt.Scan(&paket[*nPAK].Jenis)
		fmt.Print("Harga : ")
		fmt.Scan(&paket[*nPAK].Harga)
		*nPAK++
		fmt.Print("\033[H\033[2J")
		fmt.Print("Lanjutkan? (Y/N): ")
		fmt.Scan(&pilih)
	}

}

func cetakPaket(paket util.TabPKT, nPak int) {
	var pilih string
	for pilih != "C" {
		fmt.Println("----- Daftar  Paket -----")
		fmt.Printf("%-5s %-30s %s\n", "No.", "Nama Paket", "Harga")
		for i := 0; i < nPak; i++ {
			fmt.Printf("%-5d %-30s %d\n", i+1, paket[i].Jenis, paket[i].Harga)
		}
		fmt.Println("A. Edit, B. Hapus, C. Kembali")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
		if pilih == "A"  {
			//editPaket(&paket, nPak)
		} else if pilih == "B"{
			//hapusPaket(&paket, &nPak)
		} else {
			fmt.Print("Pilihan Tidak Valid, Pilihan")
			fmt.Scan(&pilih)
		}
	}
}
