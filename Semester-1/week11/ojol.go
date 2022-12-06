/* 
	program ini dibuat untuk menentukan tarif yang harus dibayar oleh pengguna ojol dengan aturan sebagai berikut:
	1. pelanggan akan dilayani dari pukul 5 pagi sampai 10 malam
	2. jika pelanggan memesan antara pukul 5 sampai 9 pagi atau antara pukul 16 sampai 19 sore, maka tarifnya adalah :
		- jika jaraknya kurang dari 10 km, maka tarifnya adalah 5000 per km
		- jika jaraknya lebih dari 10 km tetapi kurang dari 20km, maka tarifnya adalah 4500 per km
	3. jika pelanggan memesan antara pukul 9 sampai 16 sore atau antara pukul 19 sampai 22 malam, maka tarifnya adalah 4000 per km
	4. selain itu, pelanggan tidak akan dilayani
	
*/

package main

import "fmt"

func main() {
	var jarak, tarif float64
	var jam, menit int

	fmt.Println("waktu pesan?(jam, menit)")
	fmt.Scan(&jam, &menit)

	fmt.Println("Jaraknya?")
	fmt.Scan(&jarak)

	if (jam > 5 || jam <= 22) && menit > 0 || jarak > 20 {
		fmt.Print("Maaf, kami tidak bisa melayani pesanan anda")

	} else if jam > 9 && jam < 16 || (jam > 19 && jam < 22) {
		tarif = jarak * 4000
		fmt.Println("Tarif sebesar = ", tarif)

	} else {
		if jarak > 0 && jarak <= 10 {
			tarif = jarak * 5000
			fmt.Println("Tarif sebesar = ", tarif)
		} else {
			tarif = jarak * 4500
			fmt.Println("Tarif sebesar = ", tarif)

		}
	}
}
