/*
Program untuk menghitung gaji pegawai
Jabatan = staff, manager, direktur
Jika jabatan staff, maka
  - jika masa kerja < 5 tahun, maka gaji = 4000
  - jika masa kerja > 10 tahun, maka gaji = 5000 + tunjangan (tunjangan = 100 * anak)
  - jika masa kerja > 5 tahun, maka gaji = 4000 + tunjangan (tunjangan = 100 * anak)

Jika jabatan manager, maka
  - jika masa kerja > 10 tahun, maka gaji = 10000 + tunjangan(tunjangan = 300 * anak)
  - jika masa kerja < 10 tahun, maka gaji = 8500 + tunjangan(tunjangan = 300 * anak)

Jika jabatan direktur, maka gaji = 20000 + tunjangan (500 * anak)
*/
package main

import "fmt"

func main() {
	var jabat string
	var taun, anak, gaji, tunjang int

	fmt.Println("Jabatanmu apa??")
	fmt.Scan(&jabat)
	fmt.Println("Masa kerja berapa tahun? (angka saja)")
	fmt.Scan(&taun)
	fmt.Println("Anakmu berapa?")
	fmt.Scan(&anak)

	if jabat == "staff" || jabat == "Staff" {
		if taun < 5 {
			gaji = 4000
			fmt.Print(gaji, " + 0 = ", gaji)
		} else if taun > 10 {
			tunjang = 100 * anak
			gaji = 5000 + tunjang
			fmt.Print("5000 + ", tunjang, "= ", gaji)
		} else {
			tunjang = 100 * anak
			gaji = 4000 + tunjang
			fmt.Print("4000 + ", tunjang, "= ", gaji)
		}
	} else if jabat == "manager" || jabat == "Manager" {
		if taun > 10 {
			tunjang = 300 * anak
			gaji = 10000 + tunjang
			fmt.Print("10000 + ", tunjang, "= ", gaji)
		} else {
			tunjang = 300 * anak
			gaji = 8500 + tunjang
			fmt.Print("8500 + ", tunjang, "= ", gaji)
		}
	} else {
		tunjang = 500 * anak
		gaji = 20000 + tunjang
		fmt.Print("20000 + ", tunjang, "= ", gaji)
	}
}
