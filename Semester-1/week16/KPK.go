// Program untuk mengetahui Kelipatan Terkecil dari 2 bilangan
package main

import "fmt"

func main() {
	var benar bool
	var terbesar, a, b int

	fmt.Scan(&a, &b)
	if a > b || a == b {
		terbesar = a
	} else {
		terbesar = b
	}
	benar = true
	for benar {
		if terbesar%a == 0 && terbesar%b == 0 {
			benar = false
		} else {
			terbesar++
		}
	}
	fmt.Print(terbesar)
}
