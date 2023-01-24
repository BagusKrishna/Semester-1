/* Program untuk mengetahui berapa banyak bilangan terbesar dalam sejumlah baris angka yang diakhiri dengan bilangan 0
   Contoh: 1 2 3 4 5 6 7 8 9 10 0 */

package main

import "fmt"

func main() {
	var count, terbesar, bil int

	fmt.Scan(&terbesar)

	bil = terbesar
	count = 0

	for terbesar != 0 {
		if bil > terbesar {
			terbesar = bil
			count = 1
		} else if terbesar == bil {
			count += 1
		}
		fmt.Scan(&terbesar)
	}
	fmt.Println(count)
}
