//program Pramuka.
// program untuk mengecek perlengkapan pramuka per anggotanya yang terdiri dari 5 item.
// tim dikatakan siap apabila semua tim memenuhi perlengkapan yang diminta.
package main

import "fmt"

func main() {
	var i, n int
	var valid bool
	var item1, item2, item3, item4, item5 bool
	fmt.Println("Berapa jumlah anggota tim??")
	fmt.Scan(&n)

	for i = 1; i <= n; {
		fmt.Println("Input true/false terhadap item (ada 5 item)")
		fmt.Scan(&item1, &item2, &item3, &item4, &item5)
		valid = (item1 == true) && (item2 == true) && (item3 == true) && (item4 == true) && (item5 == true)
		i = i + 1
		if valid != true {
			break
		}
	}
	fmt.Println(valid)
}
