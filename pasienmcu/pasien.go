package pasienmcu

import (
	"TugasBesar/util"
	"fmt"
)

func InPasien(pas *util.TabPAS, nPAS *int, pak util.TabPKT, nPAK int) {
	var pilih, paket string
	for pilih != "N" {
		fmt.Print("\033[H\033[2J")
		fmt.Println("-----  Tambah Data  -----")
		fmt.Print("Nama Pasien  (Spasi diganti dengan '_') : ")
		fmt.Scan(&pas[*nPAS].Nama)
		util.OutPaket(pak, nPAK)
		fmt.Printf("Pilihan Paket: ")
		fmt.Scan(&paket)
		idPak := util.CariIdxPak(pak, nPAK, paket)
		if idPak != -1 {
			pas[*nPAS].PaketPas = paket
			pas[*nPAS].Biaya = pak[idPak].Harga
		} else {
			fmt.Println("Paket belum tersedia!")
		}
		fmt.Printf("Tanggal Kunjungan (DD MM YYYY) : ")
		fmt.Scan(&pas[*nPAS].Waktu.D, &pas[*nPAS].Waktu.M, &pas[*nPAS].Waktu.Y)
		*nPAS++
		fmt.Print("\033[H\033[2J")
		fmt.Print("Lanjutkan? (Y/N) : ")
		fmt.Scan(&pilih)
	}

}

func CetakPasien(pasien *util.TabPAS, nPAS *int, pak util.TabPKT, nPAK int, user util.User) {
	var pilih string
	sortPasien(pasien, *nPAS)
	for pilih != "C" {
		fmt.Print("\033[H\033[2J")
		fmt.Println("----- Daftar Pasien -----")
		fmt.Printf("%-5s %-30s %-15s %-15s %-15s %s \n", "No.", "Nama", "Paket", "Harga", "Tanggal Kunjungan", "Hasil MCU")
		for i := 0; i < *nPAS; i++ {
			fmt.Printf("%-5d %-30s %-15s %-15d %d/%d/%-15d %s\n", i+1, (*pasien)[i].Nama, (*pasien)[i].PaketPas, (*pasien)[i].Biaya, (*pasien)[i].Waktu.D, (*pasien)[i].Waktu.M, (*pasien)[i].Waktu.Y, (*pasien)[i].Rekap)
		}
		fmt.Println("A. Edit, B. Hapus, C. Kembali")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
		if pilih == "A" {
			editPasien(pasien, *nPAS, pak, nPAK, user)
		} else if pilih == "B" {
			hapusPasien(pasien, nPAS)
		} else if (pilih != "C") && (pilih != "A") && (pilih != "B") {
			fmt.Print("Pilihan tidak valid silahkan pilih kembali : ")
			fmt.Scan(&pilih)
		}
	}
}

func editPasien(pas *util.TabPAS, nPAS int, pak util.TabPKT, nPAK int, user util.User) {
	var d, m, y int
	var name, paket, pilih string
	idx := -1
	for pilih != "Y" {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Nama Pasien : ")
		fmt.Scan(&name)
		fmt.Print("Tanggal Kunjungan (DD MM YYYY) : ")
		fmt.Scan(&d, &m, &y)
		idx = util.CariIdxPas(*pas, nPAS, name, d, m, y)
		if idx != -1 {
			var pilihan string
			if user.Role == "petugas" {
				fmt.Print("Apa yang ingin Anda edit? (Nama/Paket): ")
				fmt.Scan(&pilihan)
				if pilihan == "Nama" {
					fmt.Print("Nama Pasien Baru : ")
					fmt.Scan(&pas[idx].Nama)
					fmt.Println("Data Pasien Berhasil Diubah!")
				} else if pilihan == "Paket" {
					util.OutPaket(pak, nPAK)
					fmt.Print("Paket Baru : ")
					fmt.Scan(&paket)
					idPak := util.CariIdxPak(pak, nPAK, paket)
					if idPak != -1 {
						pas[idx].PaketPas = paket
						pas[idx].Biaya = pak[idPak].Harga
						fmt.Println("Data Pasien Berhasil Diubah!")
					} else {
						fmt.Println("Paket belum tersedia!")
					}
				}
			} else if user.Role == "lab" {
				fmt.Print("Hasil Medical Checkup : ")
				fmt.Scan(&pas[idx].Rekap)
				fmt.Println("Data Pasien Berhasil Diubah!")
			}
		} else {
			fmt.Print("\033[H\033[2J")
			fmt.Println("----------------------------")
			fmt.Println("Data Pasien Tidak Ditemukan!")
			fmt.Println("----------------------------")
		}
		fmt.Print("Kembali? (Y/N) : ")
		fmt.Scan(&pilih)
	}
}

func hapusPasien(pas *util.TabPAS, nPAS *int) {
	var name, pilih string
	var d, m, y, temp int
	for pilih != "Y" {
		fmt.Print("\033[H\033[2J")
		fmt.Print("Nama Pasien : ")
		fmt.Scan(&name)
		fmt.Print("Tanggal Kunjungan (DD MM YYYY) : ")
		fmt.Scan(&d, &m, &y)
		idx := util.CariIdxPas(*pas, *nPAS, name, d, m, y)
		if idx != -1 {
			temp = idx
			for i := 0; i < *nPAS-1; i++ {
				pas[temp] = pas[temp+1]
				temp++
			}
			*nPAS--
			fmt.Println("Data berhasil dihapus")
		} else {
			fmt.Print("\033[H\033[2J")
			fmt.Println("----------------------------")
			fmt.Println("Data Pasien Tidak Ditemukan!")
			fmt.Println("----------------------------")
		}
		fmt.Print("Kembali? (Y/N) : ")
		fmt.Scan(&pilih)
	}
}

func FindPasien(pas util.TabPAS, nPAS int) {
	var pilih, kunci, K string
	var d1, d2, m1, m2, y1, y2 int
	for K != "Y" {
		fmt.Print("\033[H\033[2J")
		fmt.Println("A. Nama, B. Paket, C.Periode")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
		if pilih == "A" || pilih == "B" {
			fmt.Print("Nama : ")
			fmt.Scan(&kunci)
			fmt.Println("----- Daftar Pasien -----")
			fmt.Printf("%-5s %-30s %-15s %-15s %-15s %s \n", "No.", "Nama", "Paket", "Harga", "Tanggal Kunjungan", "Hasil MCU")
			//Sequential Search
			for i := 0; i < nPAS; i++ {
				if pas[i].PaketPas == kunci || pas[i].Nama == kunci {
					fmt.Printf("%-5d %-30s %-15s %-15d %d/%d/%-15d %s\n", i+1, pas[i].Nama, pas[i].PaketPas, pas[i].Biaya, pas[i].Waktu.D, pas[i].Waktu.M, pas[i].Waktu.Y, pas[i].Rekap)
				}
			}
		} else if pilih == "C" {
			fmt.Print("Mulai dari : ")
			fmt.Scan(&d1, &m1, &y1)
			fmt.Print("Sampai : ")
			fmt.Scan(&d2, &m2, &y2)
			fmt.Println("----- Daftar Pasien -----")
			fmt.Printf("%-5s %-30s %-15s %-15s %-15s %s \n", "No.", "Nama", "Paket", "Harga", "Tanggal Kunjungan", "Hasil MCU")
			//Sequential Search
			for i := 0; i < nPAS; i++ {
				if (pas[i].Waktu.Y > y1 || (pas[i].Waktu.Y == y1 && pas[i].Waktu.M > m1) || (pas[i].Waktu.Y == y1 && pas[i].Waktu.M == m1 && pas[i].Waktu.D >= d1)) &&
					(pas[i].Waktu.Y < y2 || (pas[i].Waktu.Y == y2 && pas[i].Waktu.M < m2) || (pas[i].Waktu.Y == y2 && pas[i].Waktu.M == m2 && pas[i].Waktu.D <= d2)) {
					fmt.Printf("%-5d %-30s %-15s %-15d %d/%d/%-15d %s\n", i+1, pas[i].Nama, pas[i].PaketPas, pas[i].Biaya, pas[i].Waktu.D, pas[i].Waktu.M, pas[i].Waktu.Y, pas[i].Rekap)
				}
			}
		} else {
			fmt.Println("Pilihan Tidak Valid, Silahkan Mengulang Kembali!")
		}
		fmt.Print("Kembali? (Y/N) : ")
		fmt.Scan(&K)
	}
}

func sortPasien(pas *util.TabPAS, nPAS int) {
	fmt.Print("\033[H\033[2J")
	var pilih string
	fmt.Print("Pengurutan (ASC/DSC) :")
	fmt.Scan(&pilih)
	pass := 1
	if pilih == "DSC" {
		// Selection Sort Descending
		for pass < nPAS {
			idx := pass - 1
			i := pass
			for i < nPAS {
				if pas[idx].Biaya < pas[i].Biaya || (pas[idx].Biaya == pas[i].Biaya &&
					(pas[idx].Waktu.Y < pas[i].Waktu.Y || (pas[idx].Waktu.Y == pas[i].Waktu.Y &&
						pas[idx].Waktu.M < pas[i].Waktu.M || (pas[idx].Waktu.M == pas[i].Waktu.M &&
						pas[idx].Waktu.D < pas[i].Waktu.D)))) {
					idx = i
				}
				i++
			}
			temp := pas[pass-1]
			pas[pass-1] = pas[idx]
			pas[idx] = temp
			pass++
		}
	} else if pilih == "ASC" {
		//Insertion Sort Ascending
		for pass < nPAS {
			temp := pas[pass]
			i := pass
			for i > 0 && (temp.Biaya < pas[i-1].Biaya || (temp.Biaya == pas[i-1].Biaya &&
				(temp.Waktu.Y < pas[i-1].Waktu.Y || (temp.Waktu.Y == pas[i-1].Waktu.Y &&
					(temp.Waktu.M < pas[i-1].Waktu.M || (temp.Waktu.M == pas[i-1].Waktu.M &&
						temp.Waktu.D < pas[i-1].Waktu.D)))))) {
				pas[i] = pas[i-1]
				i--
			}
			pas[i] = temp
			pass++
		}
	} else {
		fmt.Print("Pilihan tidak valid silahkan pilih kembali : ")
		fmt.Scan(&pilih)
	}
}
