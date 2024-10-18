package util

import "fmt"

const NMAX int = 10000

type User struct {
	Role string
}
	
type paket struct {
	Jenis string
	Harga int
}

type pasien struct {
	Nama, Rekap string
	Waktu       tanggal
	PaketPas    string
	Biaya       int
}

type tanggal struct {
	D, M, Y int
}

type TabPAS [NMAX]pasien
type TabPKT [NMAX]paket

func Login() User {
	var username, password string
	var currentRole User
	var ok bool
	for !ok {
		fmt.Println("--------- LOGIN ---------")
		fmt.Print("Username : ")
		fmt.Scan(&username)
		fmt.Print("Password : ")
		fmt.Scan(&password)
		if username == "petugas" && password == "petugas" {
			ok = true
			currentRole.Role = password
			fmt.Printf("Selamat datang, %s! Anda login sebagai %s.\n", username, currentRole.Role)
		} else if username == "lab" && password == "lab" {
			ok = true
			currentRole.Role = password
			fmt.Printf("Selamat datang, %s! Anda login sebagai %s.\n", username, currentRole.Role)
		} else {
			fmt.Print("\033[H\033[2J")
			fmt.Println("Username atau password salah")
			ok = false
		}
	}
	return currentRole
}

func CheckRole(currentUser User, role1, role2 string) bool {
	valid := false
	if currentUser.Role == role1 || currentUser.Role == role2 {
		valid = true
	} else {
		fmt.Println("Anda tidak memiliki izin untuk mengakses fitur ini.")
	}
	return valid
}

func CariIdxPas(pas TabPAS, nPas int, x string, h, b, t int) int {
	//Sequential Search
	found := -1
	var i int
	for i < nPas && found == -1 {
		if pas[i].Nama == x && pas[i].Waktu.D == h && pas[i].Waktu.M == b && pas[i].Waktu.Y == t {
			found = i
		} else {
			i++
		}
	}
	return found
}

	func CariIdxPak(paket TabPKT, nPAS int, x string) int {
		//Binary Search
		kr := 0
		kn := nPAS - 1
		found := -1
		for kr <= kn && found == -1{
			mid := (kr + kn)/2
			if paket[mid].Jenis < x {
				kr = mid + 1
			} else if paket[mid].Jenis > x {
				kn = mid - 1
			} else {
				found = mid
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
