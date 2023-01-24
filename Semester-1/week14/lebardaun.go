//program untuk mengetahui daun terlebar dari beberapa jumlah daun.
package main

import "fmt"

func main() {
	var n, bil, i, terlebar int

	fmt.Println("Masukkan banyaknya daun: ")
	fmt.Scan(&n)
	if n == 0 {
		fmt.Println("Tidak ada daun")
	} else {
		fmt.Println("Masukkan lebar daun: ")
		fmt.Scan(&bil)
		terlebar = bil
		for i = 1; i <= n-1; i++ {
			fmt.Scan(&bil)
			if bil > terlebar {
				terlebar = bil
			}

		}
		fmt.Println("Daun terlebar adalah", terlebar)
	}
}
