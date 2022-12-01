// program untuk menggandakan digit dari 2 digit yang ada

package main

import "fmt"

func main() {
	var digit, digitdepan1, digitdepan2, digitbelakang2, digitbelakang1 int

	fmt.Println("Masukkan digit")
	fmt.Scan(&digit)

	digitdepan1 = digit / 10 * 1000
	digitdepan2 = digit / 10 * 100
	digitbelakang1 = digit % 10
	digitbelakang2 = digit % 10 * 10

	digit = digitdepan1 + digitdepan2 + digitbelakang2 + digitbelakang1

	fmt.Println("Hasil = ", digit)

}
