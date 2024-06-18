package paketmcu

import (
	"TugasBesar/util"
	"fmt"
)

func InPaket(pak *util.TabPKT, nPAK *int) {
	var pilih string
	for pilih != "N" {
		fmt.Print("\033[H\033[2J")
		fmt.Println("-------------------------")
		fmt.Println("    Tambah Data Paket    ")
		fmt.Println("-------------------------")
		fmt.Print("Nama Paket (Spasi diganti dengan '_') : ")
		fmt.Scan(&pak[*nPAK].Jenis)
		fmt.Print("Harga : ")
		fmt.Scan(&pak[*nPAK].Harga)
		*nPAK++
		fmt.Print("\033[H\033[2J")
		fmt.Print("Lanjutkan? (Y/N): ")
		fmt.Scan(&pilih)
	}
	sortingPaket(pak, *nPAK)
}

func CetakPaket(paket *util.TabPKT, nPAK *int) {
	var pilih string
	for pilih != "C" {
		fmt.Print("\033[H\033[2J")
		fmt.Println("----- Daftar  Paket -----")
		fmt.Printf("%-5s %-30s %s\n", "No.", "Nama Paket", "Harga")
		for i := 0; i < *nPAK; i++ {
			fmt.Printf("%-5d %-30s %d\n", i+1, (*paket)[i].Jenis, (*paket)[i].Harga)
		}
		fmt.Println("A. Edit, B. Hapus, C. Kembali")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
		if pilih == "A" {
			editPaket(paket, *nPAK)
		} else if pilih == "B" {
			hapusPaket(paket, nPAK)
		} else if (pilih != "C") && (pilih != "A") && (pilih != "B") {
			fmt.Print("Pilihan Tidak Valid, Pilihan : ")
			fmt.Scan(&pilih)
		}
	}
}

func editPaket(paket *util.TabPKT, nPAK int) {
	var pil, name string
	for pil != "Y" {
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
		fmt.Scan(&pil)
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

func sortingPaket(paket *util.TabPKT, n int) {
	//Selection Descending
	pass := 1
	for pass < n {
		i := pass
		idx := pass - 1
		for i < n {
			if paket[idx].Harga < paket[i].Harga {
				idx = i
			}
			i++
		}
		temp := paket[pass-1]
		paket[pass-1] = paket[idx]
		paket[idx] = temp
		pass++
	}
}

func FindPaket(paket util.TabPKT, n int) {
	var name, pilih string
	for pilih != "Y" {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Nama Paket : ")
		fmt.Scan(&name)
		found := -1
		//Sequential Search
		for i := 0; i < n && found == -1; i++ {
			if paket[i].Jenis == name {
				found = i
			}
		}
		if found != -1 {
			fmt.Println("Paket Ditemukan!")
			fmt.Printf("%-30s %s\n", "Nama Paket", "Harga")
			fmt.Printf("%-30s %d\n", paket[found].Jenis, paket[found].Harga)
		} else {
			fmt.Println("Paket Tidak Ditemukan!")
		}
		fmt.Print("Kembali? (Y/N) : ")
		fmt.Scan(&pilih)
	}

}
