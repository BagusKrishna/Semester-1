// program diskon CLO
package main

import "fmt"

func main() {
	var uang, bayar int
	var clo bool

	fmt.Println("Masukkan jumlah uang :")
	fmt.Scan(&uang)

	fmt.Println("Apakah ikut CLO?")
	fmt.Scan(&clo)

	if clo {
		bayar = uang * 65 / 100

		fmt.Println("ASiiqq, kamu cuma bayar")
		fmt.Print()
		fmt.Print(bayar, " aja")
	} else {
		fmt.Println("Sori bos, NO CLO, NO DISKON")
		fmt.Println("kamu tetap bayar", uang, "yaa")
	}

}
