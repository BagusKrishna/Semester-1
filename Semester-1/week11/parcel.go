/*
	Program pengiriman Parcel
	1kg = Rp. 10rb

	tambahan biaya :
	1gr = Rp. 5 jika kurang dari 500gr
	1gr = Rp. 15 jika lebih dari 500gr

	namun, jika lebih dari 10kg, maka tidak ada tambahan biaya.

*/
package main

import "fmt"

func main() {
	var parcel, lebih, harga, parcelkg, parcelg, lebihg int

	fmt.Println("Parcelmu seberapa berat")
	fmt.Scan(&parcel)

	lebihg = parcel % 1000
	lebih = parcel % 1000
	parcelkg = parcel / 1000
	parcelg = parcelkg * 1000

	if parcel < 10000 {

		if lebih >= 500 {
			lebih = (lebih * 5)
			harga = lebih + (parcelg * 10)
		} else {
			lebih = (lebih * 15)
			harga = lebih + (parcelg * 10)
		}
		fmt.Println(parcelkg, "kg +", lebihg, "gr")
		fmt.Println("Rp.", parcelg, "+ Rp.", lebihg, "= Rp.", harga)

	} else {
		harga = parcelg
		fmt.Println(parcelkg, "kg +", lebihg, "gr")
		fmt.Println("Rp.", parcelg, "+ Rp. 0", "= Rp.", harga)
	}
}
