// program ini adalah program untuk menghitung jumlah uang yang dikeluarkan dari dompet hingga masukan 0
package main

import "fmt"

func main() {
	var uang, keluar int

	fmt.Println("Masukkan uangmu")
	fmt.Scan(&uang)
	fmt.Println("mau ambil berapa ?")
	fmt.Scan(&keluar)
	for keluar != 0 {
		fmt.Println("mau ambil berapa ?")
		fmt.Scan(&keluar)
		uang = uang - keluar

	}
	fmt.Println("uang anda sebanyak")
	fmt.Println(uang)
}
