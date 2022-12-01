// program apakah alphabet atau tidak
package main

import "fmt"

func main() {
	var char byte

	fmt.Println("masukkan apapun")
	fmt.Scanf("%c", &char)

	if 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' {
		fmt.Print("Yak, ini Alphabet masbro")
	} else {
		fmt.Print("Bukan Alphabet bos")
	}
}
