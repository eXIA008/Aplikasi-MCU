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
			menu_pasien(&dataPas, &nPas, dataPkt, nPkt)
		} else if x == 2 {
			//menu_paket()
		} else if x == 3 {
			//laporan(data, nData)
		} else if x == 4 {
			fmt.Print("\033[H\033[2J")
			fmt.Println("---  Selamat Tinggal  ---")
		}
	}

}

/*CUT THIS OUT [PASIEN.GO]---------------------------------------------------------------------*/
func menu_pasien(pas *tabPAS, nPAS *int, pak tabPKT, nPAK int) {
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
			inPasien(pas, nPAS, pak, nPAK)
		} else if x == 2 {
			findPasien(*pas, *nPAS)
		} else if x == 3 {
			cetakPasien(*pas, *nPAS, pak, nPAK)
		}
	}
}

func inPasien(pas *tabPAS, nPAS *int, pak tabPKT, nPAK int) {
	var pilih string
	for pilih != "N" {
		fmt.Print("\033[H\033[2J")
		fmt.Println("-----  Tambah Data  -----")
		fmt.Print("ID dan Nama Pasien  (Spasi diganti dengan '_') : ")
		fmt.Scan(&pas[*nPAS].ID, &pas[*nPAS].nama)
		//cetakPaket(pak, nPAK)
		fmt.Printf("Pilihan Paket: ")
		fmt.Scan(&pas[*nPAS].paketMCU.PktPas)
		fmt.Printf("Tanggal Kunjungan (DD MM YYYY) : ")
		fmt.Scan(&pas[*nPAS].waktu.D, &pas[*nPAS].waktu.M, &pas[*nPAS].waktu.Y)
		*nPAS++
		fmt.Print("\033[H\033[2J")
		fmt.Print("Lanjutkan? (Y/N) : ")
		fmt.Scan(&pilih)
	}
}

func findPasien(pas tabPAS, nPAS int) {
	var pilih, kunci string
	var d1, d2, m1, m2, y1, y2 int
	found := -1
	fmt.Print("\033[H\033[2J")
	for found != -1 {
		fmt.Println("A. Nama, B. Paket, C.Periode")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
		temp := 1
		if pilih == "A" || pilih == "B" {
			fmt.Print("Nama Pasien atau Nama Paket : ")
			fmt.Scan(&kunci)
			fmt.Println("----- Daftar Pasien -----")
			for i := 0; i < nPAS; i++ {
				if pas[i].paketMCU.PktPas == kunci || pas[i].nama == kunci {
					fmt.Printf("%-5d. %-30s %-15s %d/%d/%d\n", temp, pas[i].nama, pas[i].paketMCU.PktPas, pas[i].waktu.D, pas[i].waktu.M, pas[i].waktu.Y)
					temp++
					found = 1
				}
			}
		} else if pilih == "C" {
			fmt.Print("Mulai dari : ")
			fmt.Scan(&d1, &m1, &y1)
			fmt.Print("Sampai : ")
			fmt.Scan(&d2, &m2, &y2)
			for i := 0; i < nPAS; i++ {
				if (pas[i].waktu.Y > y1 || (pas[i].waktu.Y == y1 && pas[i].waktu.M > m1) || (pas[i].waktu.Y == y1 && pas[i].waktu.M == m1 && pas[i].waktu.D >= d1)) &&
					(pas[i].waktu.Y < y2 || (pas[i].waktu.Y == y2 && pas[i].waktu.M < m2) || (pas[i].waktu.Y == y2 && pas[i].waktu.M == m2 && pas[i].waktu.D <= d2)) {
					fmt.Printf("%-5d. %-30s %-15s %d/%d/%d\n", temp, pas[i].nama, pas[i].paketMCU.PktPas, pas[i].waktu.D, pas[i].waktu.M, pas[i].waktu.Y)
					temp++
					found = 1
				}
			}
		}
		if found == -1 {
			fmt.Print("Paket tidak Ditemukan!")
		}
	}
}

func cetakPasien(pas tabPAS, nPAS int, pak tabPKT, nPAK int) {
	var pilih string
	sortPasien(&pas, nPAS)
	for pilih != "C" {
		fmt.Print("\033[H\033[2J")
		fmt.Println("----- Daftar Pasien -----")
		fmt.Printf("%-5s %-30s %-15s %s\n", "No.", "Nama", "Paket", "Tanggal Kunjungan")
		for i := 0; i < nPAS; i++ {
			fmt.Printf("%-5d %-30s %-15s %d/%d/%d\n", i+1, pas[i].nama, pas[i].paketMCU.PktPas, pas[i].waktu.D, pas[i].waktu.M, pas[i].waktu.Y)
		}
		fmt.Println("A. Edit, B. Hapus, C. Kembali")
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilih)
		if pilih == "A" {
			editPasien(&pas, nPAS, pak, nPAK)
		} else if pilih == "B" {
			hapusPasien(&pas, &nPAS)
		}
	}
}

func editPasien(pas *tabPAS, nPAS int, pak tabPKT, nPAK int) {
	var idPas int
	var name, paket string
	found := false
	for !found {
		fmt.Print("\033[H\033[2J")
		fmt.Print("ID dan Nama Pasien yang Ingin Diedit : ")
		fmt.Scan(&idPas, &name)
		idx := cariIdxPas(*pas, nPAS, name, idPas)
		if idx != -1 {
			fmt.Print("Nama : ")
			fmt.Scan(&pas[idx].nama)

			cetakPaket(pak, nPAK)
			fmt.Print("Pilih Paket : ")
			fmt.Scan(&paket)
			idPak := cariIdxPak(pak, nPAK, paket)
			if idPak != -1 {
				pas[idx].paketMCU.PktPas = paket
			}

			fmt.Print("Tanggal Kunjungan (DD MM YYYY): ")
			fmt.Scan(&pas[idx].waktu.D, &pas[idx].waktu.M, &pas[idx].waktu.Y)

			fmt.Print("Hasil Medical Checkup : ")
			fmt.Scan(&pas[idx].rekap)
			found = true
		} else {
			fmt.Print("\033[H\033[2J")
			fmt.Scan("Data Pasien Tidak Ditemukan!")
		}
	}
}

func hapusPasien(pas *tabPAS, nPAS *int) {
	var name string
	var idPas, temp int
	found := false
	for !found {
		fmt.Print("\033[H\033[2J")
		fmt.Print("ID dan Nama Pasien yang Ingin Dihapus : ")
		fmt.Scan(&idPas, &name)
		idx := cariIdxPas(*pas, *nPAS, name, idPas)
		if idx != -1 {
			temp = idx
			for i := 0; i < *nPAS-1; i++ {
				pas[temp] = pas[temp+1]
				temp++
			}
			*nPAS--
			fmt.Println("Data berhasil dihapus")
			found = true
		} else {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Data tidak ditemukan")
		}
	}
}

func cariIdxPas(pas tabPAS, nPas int, x string, id int) int {
	found := -1
	var i int
	for i < nPas && found == -1 {
		if pas[i].nama == x && pas[i].ID == id {
			found = i
		} else {
			i++
		}
	}
	return found
}

func sortPasien(pas *tabPAS, nPAS int) {
	fmt.Print("\033[H\033[2J")
	var pilih string
	fmt.Print("Pengurutan (ASC/DSC) :")
	fmt.Scan(&pilih)
	pass := 1
	if pilih == "DSC"|| pilih == "dsc" {
		/*
			for pass < nPAS {
				i := pass
				temp1 := pas[pass].biaya
				temp2 := pas[pass].waktu
				for i > 0 && (temp1 < pas[i-1].biaya ||
					(temp1 == pas[i-1].biaya && (temp2.Y < pas[i-1].waktu.Y || (temp2.Y == pas[i-1].waktu.Y &&
						(temp2.M < pas[i-1].waktu.M || (temp2.M == pas[i-1].waktu.M && temp2.D < pas[i-1].waktu.D)))))) {
					pas[i] = pas[i-1]
					i--
				}
				pas[i].biaya = temp1
				pas[i].waktu = temp2
				pass++
			}
		*/
		// Selection Sort DESCENDING
		for pass < nPAS {
			idx := pass - 1
			i := pass
			for i < nPAS {
				if pas[idx].biaya < pas[i].biaya || (pas[idx].biaya == pas[idx].biaya &&
					(pas[idx].waktu.Y < pas[i].waktu.Y || (pas[idx].waktu.Y == pas[i].waktu.Y &&
						pas[idx].waktu.M < pas[i].waktu.M || (pas[idx].waktu.M == pas[i].waktu.M &&
						pas[idx].waktu.D < pas[i].waktu.D)))) {
					idx = i
				}
				i++
			}
			temp := pas[pass-1]
			pas[pass-1] = pas[idx]
			pas[idx] = temp
			pass++
		}
	} else if pilih == "ASC" || pilih == "asc" {

	}else {
		fmt.Print("Pilihan tidak valid silahkan pilih kembali : ")
		fmt.Scan(&pilih)
	}
	/*
		for pass < nPAS {
			i := pass
			temp1 := pas[pass].biaya
			temp2 := pas[pass].waktu
			for i > 0 && (temp1 > pas[i-1].biaya ||
				(temp1 == pas[i-1].biaya && (temp2.Y > pas[i-1].waktu.Y || (temp2.Y == pas[i-1].waktu.Y &&
					(temp2.M > pas[i-1].waktu.M || (temp2.M == pas[i-1].waktu.M && temp2.D > pas[i-1].waktu.D)))))) {
				pas[i] = pas[i-1]
				i--
			}
			pas[i].biaya = temp1
			pas[i].waktu = temp2
			pass++
		}
	*/
	// Selection Sort ASCENDING
	for pass < nPAS {
		idx := pass - 1
		i := pass
		for i < nPAS {
			if pas[idx].biaya > pas[i].biaya || (pas[idx].biaya == pas[idx].biaya &&
				(pas[idx].waktu.Y > pas[i].waktu.Y || (pas[idx].waktu.Y == pas[i].waktu.Y &&
					pas[idx].waktu.M > pas[i].waktu.M || (pas[idx].waktu.M == pas[i].waktu.M &&
					pas[idx].waktu.D > pas[i].waktu.D)))) {
				idx = i
			}
			i++
		}
		temp := pas[pass-1]
		pas[pass-1] = pas[idx]
		pas[idx] = temp
		pass++
	}

}

/*---------------------------------------------------------------------------------------------*/
