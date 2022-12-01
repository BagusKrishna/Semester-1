// program ini adalah program untuk menghitung rata-rata dari suatu bilangan n
package main

import "fmt"

func main() {
	var i, n, total, tambah, bil float64
	fmt.Println("Mau diulang berapa kali? ")
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		fmt.Println("bilangan = ? ")
		fmt.Scan(&bil)
		tambah = tambah + bil
	}
	total = tambah / n
	fmt.Println("rata-ratanya adalah ", total)
}
