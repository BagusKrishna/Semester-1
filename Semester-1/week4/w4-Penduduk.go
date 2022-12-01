//program ini adalah program untuk menghitung total penduduk
//yang dimana dilihat dari kelahiran(+), kematian(-), imigrasi(+), emigrasi(-).

package main

import "fmt"

func main() {
	var penduduk, lahir, mati, imi, emi int

	fmt.Println("Masukkan jumlah penduduk")
	fmt.Scan(&penduduk)
	fmt.Println("Masukkan jumlah kelahiran")
	fmt.Scan(&lahir)
	fmt.Println("Masukkan jumlah kematian")
	fmt.Scan(&mati)
	fmt.Println("Masukkan jumlah imigrasi")
	fmt.Scan(&imi)
	fmt.Println("Masukkan jumlah emigrasi")
	fmt.Scan(&emi)

	penduduk = penduduk + lahir + imi - emi - mati

	fmt.Println("Total penduduk = ", penduduk)

}
