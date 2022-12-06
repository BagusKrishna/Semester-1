/*
	Program menentukan apakah kamu dipecat atau tidak
	jika kamu maksimal kalah 5 kali berturut-turut, maka kamu dipecat
	jika tidak, maka kamu tidak dipecat
*/
package main

import "fmt"

func main() {
	var t1, t2, t3, t4, t5 string
	var k string

	fmt.Println("Pertandinganmu menang atau kalah?")
	fmt.Scan(&t1, &t2, &t3, &t4, &t5)

	k = "kalah"

	if t1 == k && t2 == k && t3 == k && t4 == k && t5 == k {
		fmt.Print("KAMU DIPECADD !")
	} else {
		fmt.Print("aman, gak dipecat koq")
	}
}
