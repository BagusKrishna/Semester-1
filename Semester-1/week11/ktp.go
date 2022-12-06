/*
	Program syarat membuat KTP
	Syarat = umur 17 tahun keatas
	Syarat = ada kartu keluarga
	
*/
package main

import "fmt"

func main() {
	var umur int
	var kk bool

	fmt.Println("Berapa umur kamu kack?")
	fmt.Scan(&umur)

	if umur < 17 {
		fmt.Print("Maaf kamu belum cukup umur")
	} else if umur >= 17 {

		fmt.Println("Ada kartu keluarga?")
		fmt.Scan(&kk)
		if kk {
			fmt.Print("ASikk,, bisa bikin KTP !!!")
		} else {
			fmt.Print("Loh, Kartu Keluarga mu mana kack?")
		}
	}
}
