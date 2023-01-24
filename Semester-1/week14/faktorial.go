//program untuk menghitung nilai dari suatu faktorial
package main

import "fmt"

func main() {
	var i, n, kali int

	kali = 1
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		kali = kali * i
	}
	fmt.Print(kali)
}
