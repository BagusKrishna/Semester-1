// program untuk mengetahui apakah inputan merupakan konsonan atau bukan.
package main

import "fmt"

func main() {
	var k int
	fmt.Println("masukkan sesuatu")
	fmt.Scanf("%c", &k)
	if k == 'A' || k == 'I' || k == 'U' || k == 'E' || k == 'O' || k == 'a' || k == 'i' || k == 'u' || k == 'e' || k == 'o' {
		fmt.Print("bukan konsonan")
	} else {
		fmt.Print("konsonan")
	}
}
