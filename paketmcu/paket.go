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

func cetakPaket(paket util.TabPKT, nPAK int) {
	var pilih string
	for pilih != "C" {
		fmt.Print("\033[H\033[2J")
		fmt.Println("----- Daftar  Paket -----")
		fmt.Printf("%-5s %-30s %s\n", "No.", "Nama Paket", "Harga")
		for i := 0; i < nPAK; i++ {
			fmt.Printf("%-5d %-30s %d\n", i+1, paket[i].Jenis, paket[i].Harga)
		}
		fmt.Println("A. Edit, B. Hapus, C. Kembali")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
		if pilih == "A" {
			editPaket(&paket, nPAK)
		} else if pilih == "B" {
			hapusPaket(&paket, &nPAK)
		} else if pilih != "C" && pilih != "A" && pilih != "B" {
			fmt.Print("Pilihan Tidak Valid, Pilihan : ")
			fmt.Scan(&pilih)
		}
	}
}

func editPaket(paket *util.TabPKT, nPAK int) {
	var pilih, name string
	for pilih != "Y" {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Nama Paket : ")
		fmt.Scan(&name)
		idx := util.CariIdxPak(*paket, nPAK, name)
		if idx != -1 {
			fmt.Print("Nama Paket Baru : ")
			fmt.Scan(&paket[idx].Jenis)
			fmt.Print("Harga Baru : ")
			fmt.Scan(&paket[idx].Harga)
			fmt.Println("Paket berhasil diubah")
		} else {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Paket tidak terubah/tidak ditemukan!")
		}
		fmt.Print("Kembali? (Y/N) : ")
		fmt.Scan(&pilih)
	}
}

func hapusPaket(A *util.TabPKT, n *int) {
	var name, pilih string
	var idx, temp int
	for pilih != "Y" {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Nama Paket : ")
		fmt.Scan(&name)
		idx = util.CariIdxPak(*A, *n, name)
		if idx != -1 {
			temp = idx
			for i := 0; i < *n-1; i++ {
				A[temp] = A[temp+1]
				temp++
			}
			*n--
			fmt.Println("Paket berhasil dihapus")
		} else {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Paket tidak terhapus/tidak ditemukan!")
		}
		fmt.Print("Kembali? (Y/N) : ")
		fmt.Scan(&pilih)
	}
}
