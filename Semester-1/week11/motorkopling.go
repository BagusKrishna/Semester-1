/*
program motorkopling

	jadi, disini kita punya motor gigi 0 - 4
	kalo mau nyala, harus lah tarik kopling (kopling = true)
	gigi 0 itu netral
	kalo gigi 0 tarik kopling, mesinnya idup. Tapi gaakan bisa jalan mau di gas mau nggak
	kalo gigi 1 - 4, kopling ditarik, mesinnya hidup, dan gaakan jalan kalau di gas ataupun tidak
	kalo gigi 1 - 4, koplingnya ga ditarik, tapi di gas, motornya jalan
	kalo gigi 1 - 4, koplingnya ga ditarik, gak di gas juga, motornya mati
*/
package main

import "fmt"

func main() {
	var gigi int
	var kopling, gas bool

	fmt.Println("Masukkan gigi")
	fmt.Scan(&gigi)

	if gigi > 5 {
		fmt.Println("Anjiir gigimu kebanyakan")
		//karena maks giginya 4

	} else {
		if gigi == 0 {
			fmt.Println("Kopling ditarik gak?")
			fmt.Scan(&kopling)

			fmt.Println("Apakah di gass?")
			fmt.Scan(&gas)

			if kopling {
				fmt.Print("Mesin menyala dan motor tidak berjalan")
				//karena gigi 0 mau gas atau enggak asal kopling ditarik pasti seperti output

			} else {
				fmt.Print("Mesin mati")
				//karena gigi 0 kalo ga di kopling dia bakal mati. (sebenernya kalo ga dikopling jg bakal idup, cuman mksdnya disini untuk awal menghidupkan)
			}
		} else if gigi == 1 || gigi == 2 || gigi == 3 || gigi == 4 {
			fmt.Println("Kopling ditarik gak?")
			fmt.Scan(&kopling)

			fmt.Println("Apakah di gass?")
			fmt.Scan(&gas)

			if kopling {
				fmt.Print("Mesin menyala dan motor tidak berjalan")
				//kalo dia di kopling aja dan gak di gas, mesinnya tetep hidup tapi gak jalan

			} else if !kopling && gas {
				fmt.Print("Motor berjalan")
				//motor itu bisa jalan kalo di gas, tapi kalo sambil tarik kopling, dia gak akan jalan, tapi tetep nyala

			} else {
				fmt.Print("Mesin mati")
				//syarat mesin mati kalo gak tarik koplingnya
			}
		}
	}
}
