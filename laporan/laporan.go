package laporan

import (
	"TugasBesar/util"
	"fmt"
)

func Laporan(pasien util.TabPAS, nPas int) {
	var pilih string
	var pemasukan int
	var d1, d2, m1, m2, y1, y2 int
	for pilih != "Y" {
		fmt.Print("Periode : ")
		fmt.Scan(&d1, &m1, &y1)
		fmt.Print("Sampai  : ")
		fmt.Scan(&d2, &m2, &y2)
		fmt.Println("-------- Laporan Pemasukan --------")
		fmt.Printf("----Periode %d/%d/%d - %d/%d/%d----\n", d1, m1, y1, d2, m2, y2)
		for i := 0; i < nPas; i++ {
			if (pasien[i].Waktu.Y > y1 || (pasien[i].Waktu.Y == y1 && pasien[i].Waktu.M > m1) || (pasien[i].Waktu.Y == y1 && pasien[i].Waktu.M == m1 && pasien[i].Waktu.D >= d1)) &&
				(pasien[i].Waktu.Y < y2 || (pasien[i].Waktu.Y == y2 && pasien[i].Waktu.M < m2) || (pasien[i].Waktu.Y == y2 && pasien[i].Waktu.M == m2 && pasien[i].Waktu.D <= d2)) {
				fmt.Printf("%d. %s\t%s\t%d\t%d/%d/%d\n", i+1, pasien[i].Nama, pasien[i].PaketPas, pasien[i].Biaya, pasien[i].Waktu.D, pasien[i].Waktu.M, pasien[i].Waktu.Y)
				pemasukan += pasien[i].Biaya
			}
		}
		fmt.Printf("Total Pemasukan = Rp%d \n", pemasukan)
		fmt.Println("-----------------------------------")
		fmt.Print("Kembali? (Y/N) ")
		fmt.Scan(&pilih)
	}

}
