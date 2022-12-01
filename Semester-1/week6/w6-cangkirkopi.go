// program untuk membuat kopi.
// jadi untuk membuat kopi diperlukan x gula dan y kopi.
package main

import "fmt"

func main() {
	var perlugula, perlukopi, xgula, ykopi, bagiGula, bagiKopi int

	fmt.Println("bahan - bahan ? ")
	fmt.Scan(&perlugula, &perlukopi, &xgula, &ykopi)
	bagiGula = perlugula / xgula
	bagiKopi = perlukopi / ykopi

	if bagiGula <= bagiKopi {
		fmt.Println("Kopi gak bisa dibuat :(")
	} else {
		fmt.Println("Kopi yang berhasil dibuat = ", bagiKopi)
	}

}
