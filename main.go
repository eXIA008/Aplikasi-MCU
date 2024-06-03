package main
import (
	"fmt"
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
			//pasienmcu.menupasien(&dataPas, &nPas, dataPkt, nPkt)
		} else if x == 2 {
			paketmcu.MenuPaket(&dataPkt, &nPkt)
		} else if x == 3 {
			laporan.Laporan(dataPas, nPas)
		} else if x == 4 {
			fmt.Print("\033[H\033[2J")
			fmt.Println("---  Selamat Tinggal  ---")
		}
	}
}