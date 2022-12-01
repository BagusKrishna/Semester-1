// program ini adalah program untuk menghitung luas dan keliling persegi sebanyak n buahe
package main

import "fmt"

func main() {
	var n, i, sisi, luas, keliling int

	fmt.Println("Mau diulang berapa kali? ")
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		fmt.Println("sisi = ? ")
		fmt.Scan(&sisi)
		luas = sisi * sisi
		keliling = 4 * sisi
		fmt.Println("Luas = ", luas, "Keliling = ", keliling)
	}
}
