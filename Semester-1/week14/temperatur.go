package main

import "fmt"

func main() {
	var i, sum, tertinggi, terendah, temp, prevtemp float64

	fmt.Scan(&temp)
	prevtemp = 69
	i = 1
	sum = temp
	tertinggi = temp
	terendah = temp
	for temp != 0 || prevtemp != 0 {
		prevtemp = temp
		fmt.Scan(&temp)
		if temp != 0 || prevtemp != 0 {
			sum = sum + temp
			i++
		}
		if temp > tertinggi {
			tertinggi = temp
		} else if terendah < temp {
			terendah = temp
		}
	}

	fmt.Println("Tertinggi = ", tertinggi)
	fmt.Println("Terendah = ", terendah)
	fmt.Println("Rata-rata = ", sum/i)
}
