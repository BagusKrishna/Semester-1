/*
	Program ini dibuat untuk menentukan diskon dan cashback yang didapatkan dengan aturan sebagai berikut:

	1. Jika semua poin adalah ganjil, maka diskon adalah 1.7 kali jumlah poin yang ganjil
	2. Jika semua poin adalah genap, maka cashback adalah 3.1 kali jumlah poin yang genap
	3. Jika terdapat poin ganjil dan genap, maka diskon adalah 1.7 kali jumlah poin yang ganjil dan cashback adalah 3.1 kali jumlah poin yang genap
	4. cashback maksimal 35%
	5. diskon maksimal 50%
	6. Jika punya membership, maka diskon dan cashback ditambah 15%


*/

package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var kata string
	var diskon, cashback float64

	fmt.Println("Apakah punya membership (Yes/No)")
	fmt.Scan(&kata)

	fmt.Println("Masukkan jumlah poin anda")
	fmt.Scan(&a, &b, &c, &d, &e)

	if a%2 == 1 && b%2 == 1 && c%2 == 1 && d%2 == 1 && e%2 == 1 {
		diskon = 1.7 * float64(c+d+e)
	} else if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
		cashback = 3.1 * float64(a+b+c)
	} else {
		diskon = 1.7 * float64(c+d+e)
		cashback = 3.1 * float64(a+b+c)
	}

	if cashback >= 35 {
		cashback = 35
	} else if diskon >= 50 {
		diskon = 50
	} else if kata == "Yes" || kata == "yes" {
		diskon = diskon + diskon*15/100
		cashback = cashback + cashback*15/100
	}

	fmt.Print("Cashback = ", cashback, "% ", "Diskon = ", diskon, "%")
}
