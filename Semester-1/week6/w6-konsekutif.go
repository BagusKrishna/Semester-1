// program untuk mengetahui apakah angka yang ada merupakan selisih 1 antara angka 1 dan disebelahnya
package main

import "fmt"

func main() {
	var angka, digit1, digit2, selisih1, selisih2 int
	var apakah1 bool

	fmt.Println("Masukkin angkamu kak")
	fmt.Scan(&angka)

	for angka > 0 {
		digit1 = angka % 10
		angka = angka / 10
		digit2 = angka % 10
		angka = angka / 10
		selisih1 = digit1 - digit2
		selisih2 = digit2 - digit1
		apakah1 = selisih1 == 1 || selisih2 == 1
		if selisih1 != 1 || selisih2 != 1 {
			break
		}
	}

	fmt.Println("Apakah berselisih 1?", apakah1)
}
