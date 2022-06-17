package main

import "fmt"

func main() {
	var jumlahBuah = map[string]int{"apel": 1, "melon": 2}
	// fmt.Println(jumlahBuah["apel"])
	for key, element := range jumlahBuah {
		fmt.Println(key, element)
	}
}
