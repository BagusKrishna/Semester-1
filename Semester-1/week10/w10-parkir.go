// program untuk mengetahui berapa menit atau jam seseorang melakukan parkir.
package main

import "fmt"

func main() {
	var h1, m1, h2, m2, jam, menit int

	fmt.Println("Jam berapa kak??")
	fmt.Scan(&h1, &m1, &h2, &m2)
	if h1 > h2 && m1 > m2 {
		jam = 12 - h1 + h2 - 1
		menit = 60 - m1 + m2
	} else if h1 > h2 {
		jam = 12 - h1 + h2
		menit = m2 - m1
	} else {
		jam = h2 - h1
		menit = m2 - m1
	}
	fmt.Println(jam, "jam", menit, "menit")
}
