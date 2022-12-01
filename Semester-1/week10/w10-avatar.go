// program untuk numpang sama avatar
// dewasa cuman ada 3, 1 dewasa muat 5 orang
// kecil ada 5, cuman buat 2 orang
// cari tahu perlu berapa dewasa dan kecil. kalau penuh ada yang tidak berangkat.
package main

import "fmt"

func main() {
	var penumpang, dewasa, kecil, takBerangkat int
	fmt.Println("Berapa penumpang nich kakk?? ")
	fmt.Scan(&penumpang)
	if penumpang <= 15 && penumpang%5 != 0 {
		dewasa = penumpang/5 + 1
		fmt.Print("Dewasa ", dewasa)
	} else if penumpang <= 15 && penumpang%5 == 0 {
		dewasa = penumpang / 5
		fmt.Print("Dewasa ", dewasa)
	} else if penumpang > 15 && penumpang <= 25 {
		dewasa = 3
		if (penumpang-15)%2 != 0 {
			kecil = (penumpang-15)/2 + 1
		} else {
			kecil = (penumpang - 15) / 2
		}
		fmt.Print("Dewasa ", dewasa, ", kecil ", kecil)
	} else {
		dewasa = 3
		kecil = 5
		takBerangkat = penumpang % 25
		fmt.Print("Dewasa ", dewasa, ", kecil ", kecil, ", dan ", takBerangkat, " tak berangkat")
	}
}
